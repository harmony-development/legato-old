package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"io/ioutil"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"
)

// UpdateGuildPictureData is the data for a guild picture update request
type UpdateGuildPictureData struct {
	Guild int64 `validate:"required"`
}

// UpdateGuildPicture is the request to update a guild's picture
func (h Handlers) UpdateGuildPicture(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data UpdateGuildPictureData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Form required")
	}
	files := form.File["files"]
	if len(files) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	file, err := files[0].Open()
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error opening file")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			sentry.CaptureException(err)
		}
	}()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	scaled, err := bimg.NewImage(fileBytes).Process(bimg.Options{
		Height:  128,
		Width:   128,
		Quality: 60,
		Crop:    true,
	})
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	fileName := randstr.Hex(16)
	if err := h.Deps.StorageManager.AddGuildPicture(fileName, scaled); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving file upload")
	}
	oldPicture, err := h.Deps.DB.GetGuildPicture(data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error removing old picture")
	}
	if err := h.Deps.DB.SetGuildPicture(data.Guild, fileName); err != nil {
		h.Deps.StorageManager.DeleteGuildPicture(fileName)
		return echo.NewHTTPError(http.StatusInternalServerError, "error updating picture")
	}
	h.Deps.StorageManager.DeleteGuildPicture(oldPicture)
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
		Type: "GuildPictureUpdate",
		Data: map[string]interface{}{
			"guild":   data.Guild,
			"picture": fileName,
		},
	})
	return ctx.NoContent(http.StatusOK)
}
