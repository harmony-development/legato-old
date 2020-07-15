package v1

import (
	"crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/util"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

type MessageData struct {
	Content string   `json:"content"`
	Embeds  []string `json:"embeds"`
	Actions []string `json:"actions"`
}

// Message : Receive a message from a client.
func (h Handlers) Message(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(MessageData)
	if data.Content == "" && len(data.Embeds) == 0 && len(data.Actions) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	form, err := ctx.MultipartForm()
	var attachments []string
	var files []*multipart.FileHeader
	if err != nil {
		goto pastTheForm
	}
	files = form.File["files"]
	attachments = make([]string, len(files))
	if len(files) > h.Deps.Config.Server.MaxAttachments {
		return echo.NewHTTPError(http.StatusBadRequest, "too many files uploaded")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests, please try again later")
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if h.Deps.State.Guilds[*ctx.Location.GuildID] == nil || h.Deps.State.Guilds[*ctx.Location.GuildID].Clients[ctx.UserID] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "insufficient permissions to send message")
	}
	if len(files) > 0 {
		for i, v := range files {
			file, err := v.Open()
			if err != nil {
				sentry.CaptureException(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error opening file")
			}
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				sentry.CaptureException(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error reading file")
			}
			err = file.Close()
			if err != nil {
				sentry.CaptureException(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error closing files")
			}
			sum := sha512.New().Sum(fileBytes)
			fileID, err := h.Deps.DB.GetFileIDFromHash(sum)
			if err != nil {
				fileName := randstr.Hex(16)
				if err := h.Deps.StorageManager.AddImage(fileName, fileBytes); err != nil {
					sentry.CaptureException(err)
					return echo.NewHTTPError(http.StatusInternalServerError, "Error writing file")
				}
			} else {
				attachments[i] = fileID
			}
		}
	}
pastTheForm:
	var actions [][]byte
	var embeds [][]byte
	if len(data.Actions) > 0 {
		for _, action := range data.Actions {
			parsed, err := CleanAction([]byte(action))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			actions = append(actions, parsed)
		}
	}
	if len(data.Embeds) > 0 {
		for _, embed := range data.Embeds {
			parsed, err := CleanEmbed([]byte(embed))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			embeds = append(embeds, parsed)
		}
	}
	msg, err := h.Deps.DB.AddMessage(*ctx.Location.ChannelID, *ctx.Location.GuildID, ctx.UserID, data.Content, attachments, embeds, actions)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error storing message in DB")
	}
	var rawEmbeds, rawActions []json.RawMessage
	for _, embed := range embeds {
		rawEmbeds = append(rawEmbeds, json.RawMessage(embed))
	}
	for _, action := range actions {
		rawActions = append(rawActions, json.RawMessage(action))
	}
	if h.Deps.State.Guilds[*ctx.Location.GuildID] == nil {
		return c.JSON(http.StatusOK, MessageCreateResponse{util.U64TS(msg.MessageID)})
	}
	h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
		Type: MessageCreateEventType,
		Data: MessageCreateEvent{
			GuildID:     util.U64TS(*ctx.Location.GuildID),
			ChannelID:   util.U64TS(*ctx.Location.ChannelID),
			CreatedAt:   msg.CreatedAt.Unix(),
			Message:     msg.Content,
			Attachments: attachments,
			AuthorID:    util.U64TS(ctx.UserID),
			MessageID:   util.U64TS(msg.MessageID),
			Actions:     rawActions,
			Embeds:      rawEmbeds,
		},
	})
	return c.JSON(http.StatusOK, MessageCreateResponse{util.U64TS(msg.MessageID)})
}
