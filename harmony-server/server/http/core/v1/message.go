package v1

import (
	"crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

type MessageData struct {
	Content string `validate:"required"`
	Embeds  []string
	Actions []string
}

// Message : Receive a message from a client.
func (h Handlers) Message(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(*MessageData)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	files := form.File["files"]
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
	var attachments = make([]string, len(files))
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
	h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
		Type: "MessageAdd",
		Data: map[string]interface{}{
			"guild":       *ctx.Location.GuildID,
			"channel":     *ctx.Location.ChannelID,
			"createdAt":   time.Now().UTC().Unix(),
			"message":     msg,
			"attachments": attachments,
			"userID":      ctx.UserID,
			"messageID":   msg.MessageID,
			"actions":     rawActions,
			"embeds":      rawEmbeds,
		},
	})
	return nil
}
