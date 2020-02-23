package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type deleteGuildData struct {
	Token       string `mapstructure:"token"`
	Guild       string `mapstructure:"guild"`
}

func OnDeleteGuild(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data deleteGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil || globals.Guilds[data.Guild].Owner != ws.Userid {
		sendErr(ws, "You can't delete this guild since you don't own it!")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're deleting guilds too fast, try again in a bit")
		return
	}
	transaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		sendErr(ws, "We couldn't delete the guild. Try again in a bit please")
		return
	}
	_, err = transaction.Exec("DELETE FROM messages WHERE guildid=$1", data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to delete the guild messages, cancelling deletion")
		golog.Warnf("Error deleting guild messages : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM channels WHERE guildid=$1", data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to delete the guild channels, cancelling deletion")
		golog.Warnf("Error deleting guild channels : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM  invites WHERE guildid=$1", data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to delete the guild invites, cancelling deletion")
		golog.Warnf("Error deleting guild invites : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM guildmembers WHERE guildid=$1", data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to kick all the members, cancelling deletion")
		golog.Warnf("Error deleting guild : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM guilds WHERE guildid=$1", data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to delete the guild, cancelling deletion")
		golog.Warnf("Error deleting guild : %v", err)
		return
	}
	err = transaction.Commit()
	if err != nil {
		sendErr(ws, "We weren't able to complete guild deletion, cancelling")
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deleteguild",
				Data: map[string]interface{}{
					"guild": data.Guild,
				},
			})
		}
	}
	delete(globals.Guilds, data.Guild)
	return
}