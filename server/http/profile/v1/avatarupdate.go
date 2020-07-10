package v1

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/h2non/bimg"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/util"
)

func (h Handlers) AvatarUpdate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	avatar, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InvalidRequest)
	}
	file, err := avatar.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	resized, err := bimg.Resize(bytes, bimg.Options{
		Width:   h.Config.Server.Avatar.Width,
		Height:  h.Config.Server.Avatar.Height,
		Quality: h.Config.Server.Avatar.Quality,
		Crop:    h.Config.Server.Avatar.Crop,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	sum := sha512.New().Sum(resized)
	fileID, err := h.DB.GetFileIDFromHash(sum)
	if err != nil {
		fileID = randstr.Hex(16)
		if err := ioutil.WriteFile(fmt.Sprintf("./avatars/%v", fileID), resized, 0666); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
	} else {
		if err := h.DB.UpdateAvatar(ctx.UserID, fileID); err != nil {
			sentry.CaptureException(err)
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
	}
	oldAvatar, err := h.DB.GetAvatar(ctx.UserID)
	if err == nil && oldAvatar.Valid {
		h.StorageManager.DeleteAvatar(oldAvatar.String)
	}
	for c := range h.State.UserUpdateListeners {
		c.Send(&client.OutPacket{
			Type: AvatarUpdateEventType,
			Data: AvatarUpdateEvent{
				UserID:    util.U64TS(ctx.UserID),
				NewAvatar: fileID,
			},
		})
	}
	return ctx.NoContent(http.StatusOK)
}
