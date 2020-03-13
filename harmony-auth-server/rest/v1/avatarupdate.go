package v1

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"harmony-auth-server/rest/util"
	"io/ioutil"
	"net/http"
	"path"
)

func AvatarUpdate(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Valid form required")
	}
	session := ctx.FormValue("session")
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many avatar updates, please try again later")
	}
	user, err := db.GetUserBySession(session)
	if err != nil {
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
		Height: 128,
		Width: 128,
		Quality: 60,
		Crop: true,
	})
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), scaled, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving file upload")
	}
	var oldAvatarID string
	err = db.DB.QueryRow("SELECT avatar FROM users WHERE id=$1", user.ID).Scan(&oldAvatarID)
	_, err = db.DB.Exec("UPDATE users SET avatar=$1 WHERE id=$2", fname, user.ID)
	if err != nil {
		go db.DeleteFromAvatars(fname)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	go db.DeleteFromAvatars(path.Base(oldAvatarID))
	servers, err := db.ListServersTransaction(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to broadcast username update, please try again later")
	}
	for _, server := range servers {
		go server.SendAvatarUpdate(user.ID, fname)
	}
	return nil
}