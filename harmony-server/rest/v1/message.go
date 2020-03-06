package v1

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	maxFiles = 3
)

func Message(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Valid form required")
	}
	files := form.File["files"]
	message, guild, channel :=  ctx.FormValue("message"), ctx.FormValue("guildid"), ctx.FormValue("channelid")
	if message == "" || channel == "" || guild == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Valid form required")
	}
	if len(files) > maxFiles {
		return echo.NewHTTPError(http.StatusBadRequest, "too many files uploaded")
	}
	userid, err := authentication.VerifyToken(ctx.FormValue("token"))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "invalid token")
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[userid] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "not permitted to send messages in this guild/channel")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests, please try again later")
	}
	var messageID = randstr.Hex(16)
	var attachments = make([]string, len(files))
	if len(files) > 0 {
		fileTransaction, err := harmonydb.DBInst.Begin()
		if err != nil {
			golog.Warnf("error making file transaction : %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error saving files")
		}
		for i, v := range files {
			file, err := v.Open()
			if err != nil {
				golog.Warnf("Failed to parse file : %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error opening file")
			}
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				golog.Warnf("Error reading uploaded file : %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error reading file")
			}
			err = file.Close()
			if err != nil {
				golog.Warnf("Failed to close file : %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error closing files")
			}
			fname := randstr.Hex(16)
			err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), fileBytes, 0666)
			if err != nil {
				golog.Warnf("Error saving file upload : %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error writing file")
			}
			_, err = fileTransaction.Exec("INSERT INTO attachments(messageid, attachment) VALUES($1, $2)", messageID, fname)
			if err != nil {
				golog.Warnf("Error inserting into attachments : %v", err)
				go deleteFromFilestore(fname)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error linking file to message")
			} else {
				attachments[i] = fname
			}
		}
		err = fileTransaction.Commit()
		if err != nil {
			golog.Warnf("Error committing attachment transaction  : %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error committing attachment transaction")
		}
	}

	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "message",
				Data: map[string]interface{}{
					"guild":     guild,
					"channel":   channel,
					"userid":    userid,
					"createdat": time.Now().UTC().Unix(),
					"message":   message,
					"attachments": attachments,
					"messageid": messageID,
				},
			})
		}
	}
	_, err = harmonydb.DBInst.Exec("INSERT INTO messages(messageid, guildid, channelid, author, createdat, message) VALUES($1, $2, $3, $4, $5, $6)", messageID, guild, channel, userid, time.Now().UTC().Unix(), message)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving message")
	}
	return nil
}
