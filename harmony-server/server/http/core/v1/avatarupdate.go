package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"harmony-server/server/http/socket/client"
)

type AvatarUpdateEvent struct {
	UserID    uint64
	NewAvatar string
}

func (h Handlers) AvatarUpdate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
	}
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
		Width:   h.Deps.Config.Server.Avatar.Width,
		Height:  h.Deps.Config.Server.Avatar.Height,
		Quality: h.Deps.Config.Server.Avatar.Quality,
		Crop:    h.Deps.Config.Server.Avatar.Crop,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	fname := randstr.Hex(16)
	if err := ioutil.WriteFile(fmt.Sprintf("./avatars/%v", fname), resized, 0666); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	oldAvatar, err := h.Deps.DB.GetAvatar(ctx.UserID)
	if err == nil && !oldAvatar.Valid {
		h.Deps.StorageManager.DeleteAvatar(oldAvatar.String)
	}
	for c := range h.Deps.State.UserUpdateListeners {
		c.Send(&client.OutPacket{
			Type: "AvatarUpdate",
			Data: AvatarUpdateEvent{
				UserID:    ctx.UserID,
				NewAvatar: fname,
			},
		})
	}
	return ctx.NoContent(http.StatusOK)
}
