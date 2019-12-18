package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type createGuildData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnCreateGuild(ws *globals.Client, rawMap map[string]interface{}) {
	var data createGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	guildid := randstr.Hex(16)
	createGuildTransaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		sendErr(ws, "That guild didn't work. Please try again")
		golog.Warnf("Error beginning createGuildTransaction : %v", err)
		return
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO guilds(guildid, guildname, picture, owner) VALUES($1, $2, $3, $4);`, guildid, data.Guild, "", userid)
	if err != nil {
		createGuildError(ws)
		golog.Warnf("Error inserting guild : %v", err)
		return
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO guildmembers(userid, guildid) VALUES($1, $2);`, userid, guildid)
	if err != nil {
		createGuildError(ws)
		golog.Warnf("Error inserting guild member on guild create : %v", err)
		return
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO channels(channelid, guildid, channelname) VALUES($1, $2, $3)`, randstr.Hex(16), guildid, "general")
	if err != nil {
		createGuildError(ws)
		golog.Warnf("Error inserting channel on guild create : %v", err)
		return
	}
	err = createGuildTransaction.Commit()
	if err != nil {
		createGuildError(ws)
		golog.Warnf("Error commiting createGuildTransaction : %v", err)
		return
	}
	ws.Send(&globals.Packet{
		Type: "createguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
}

func createGuildError(ws *globals.Client) {
	ws.Send(&globals.Packet{
		Type: "createguild",
		Data: map[string]interface{}{
			"message": "error creating guild",
		},
	})
}