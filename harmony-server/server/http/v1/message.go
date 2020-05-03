package v1

import (
	"github.com/getsentry/sentry-go"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/handling"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"io/ioutil"
	"net/http"
	"time"
)

type MessageData struct {
	Guild   string `validate:"required"`
	Channel string `validate:"required"`
	Message string `validate:"required"`
}

func (h Handlers) Message(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	files := form.File["files"]
	var data MessageData
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if len(files) > h.Deps.Config.Server.MaxAttachments {
		return echo.NewHTTPError(http.StatusBadRequest, "too many files uploaded")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests, please try again later")
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if h.Deps.State.Guilds[data.Guild] == nil || h.Deps.State.Guilds[data.Guild].Clients[ctx.UserID] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "insufficient permissions to send message")
	}
	var messageID = randstr.Hex(16)
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
			fileName := randstr.Hex(16)
			if err := h.Deps.StorageManager.AddImage(fileName, fileBytes); err != nil {
				sentry.CaptureException(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error writing file")
			}
			attachments[i] = fileName
		}
		if err := h.Deps.DB.AddAttachments(messageID, attachments); err != nil {
			sentry.CaptureException(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error committing attachment transaction")
		}
	}
	h.Deps.State.Guilds[data.Guild].Broadcast(&handling.OutPacket{
		Type: "MessageAdd",
		Data: map[string]interface{}{
			"guild":       data.Guild,
			"channel":     data.Channel,
			"createdAt":   time.Now().UTC().Unix(),
			"message":     data.Message,
			"attachments": attachments,
			"userID":      userID,
			"messageID":   messageID,
		},
	})
	if _, err := h.Deps.DB.Exec(`INSERT INTO messages(messageid, guildid, channelid, author, createdat, message)
			VALUES($1, $2, $3, $4, $5, $6)`,
			messageID, data.Guild, data.Channel, userID, data.Message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving message")
	}
	return nil
}
