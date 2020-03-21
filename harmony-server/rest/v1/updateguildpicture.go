package v1

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"
	"harmony-server/globals"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"io/ioutil"
	"net/http"
	"path"
)

func UpdateGuildPicture(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Form required")
	}
	files := form.File["files"]
	guild := ctx.Param("guildid")
	if len(files) == 0 || guild == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Valid form required")
	}
	file, err := files[0].Open()
	if err != nil {
		golog.Debugf("Error opening file : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Form required")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests, please try again later")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			golog.Warnf("Error closing file : %v", err)
		}
	}()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
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
	var oldPictureID string
	err = db.DBInst.QueryRow("SELECT picture FROM guilds WHERE guildid=$1", guild).Scan(&oldPictureID)
	_, err = db.DBInst.Exec("UPDATE guilds SET picture=$1 WHERE guildid=$2 AND owner=$3", fname, guild, ctx.User.ID)
	if err != nil {
		golog.Warnf("Error updating picture. %v", err)
		go deleteFromFilestore(fname)
		return echo.NewHTTPError(http.StatusInternalServerError, "error linking picture to guild")
	}
	go deleteFromFilestore(path.Base(oldPictureID))
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "updateguildpicture",
				Data: map[string]interface{}{
					"guild":   guild,
					"picture": fname,
					"success": true,
				},
			})
		}
	}
	return nil
}