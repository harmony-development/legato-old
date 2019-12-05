package event

import (
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type createGuildData struct {
	Token string
	Guildname string
}

func OnCreateGuild(ws *socket.Client, rawMap map[string]interface{}) {
	var data createGuildData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Guildname, ok = rawMap["guildname"].(string); !ok {
		return
	}
	userid := verifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	guildid := randstr.Hex(16)
	_, err := harmonydb.DBInst.Exec(`INSERT INTO guilds(guildid, guildname, picture) VALUES(?, ?, ?); 
										   INSERT INTO guildmembers(userid, guildid) VALUES(?, ?);
										   INSERT INTO channels(channelid, guildid, channelname) VALUES(?, ?, ?)`, guildid, data.Guildname, "", userid, guildid, randstr.Hex(16), guildid, "general")
	if err != nil {
		golog.Warnf("Error creating guild : %v", err)
		ws.Send(&socket.Packet{
			Type: "createguild",
			Data: map[string]interface{}{
				"message": "error creating guild",
			},
		})
		return
	}
	ws.Send(&socket.Packet{
		Type: "createguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
}