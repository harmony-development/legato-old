package v1

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"harmony-auth-server/harmony/auth"
	"harmony-auth-server/harmony/http/hm"
	"harmony-auth-server/harmony/http/util"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"
	"io/ioutil"
	"net/http"
)

type avatarUpdateData struct {
	Session    string `validate:"required"`
	APIVersion string `validate:"required"`
}

// AvatarUpdate handles request to update a user's avatar
func (h Handlers) AvatarUpdate(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := new(avatarUpdateData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Valid form required")
	}
	var session *auth.Session
	var exists bool
	if session, exists = h.AuthManager.Sessions.GetSession(data.Session); !exists {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests, please try again later")
	}
	fileBytes, err := util.GetFile(form, "files")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	scaled, err := bimg.NewImage(fileBytes).Process(bimg.Options{
		Height:  h.Config.Server.AvatarHeight,
		Width:   h.Config.Server.AvatarWidth,
		Quality: h.Config.Server.AvatarQuality,
		Crop:    h.Config.Server.AvatarCrop,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), scaled, 0666)
	if err != nil {
		logrus.Warnf("Error saving file upload : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving file upload")
	}

	var oldAvatarID string

	err = h.DB.QueryRow("SELECT avatar FROM users WHERE id=$1", session.UserID).Scan(&oldAvatarID)
	if err != nil {
		h.StorageManager.DeleteAvatar(fname)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	_, err = h.DB.Exec("UPDATE users SET avatar=$1 WHERE id=$2", fname, session.UserID)
	if err != nil {
		h.StorageManager.DeleteAvatar(fname)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}

	h.StorageManager.DeleteAvatar(fname)

	servers, err := h.DB.GetInstanceList(session.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to broadcast username update, please try again later")
	}

	for _, server := range servers {
		go server.SendAvatarUpdate(session.UserID, fname, data.APIVersion)
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated profile picture",
	})
}
