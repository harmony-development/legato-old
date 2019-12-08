package main

import (
	"harmony-server/socket"
	"harmony-server/socket/event"
	"net/http"
)

func handleSocket(w http.ResponseWriter, r *http.Request) {
	ws := socket.NewSocket(w, r)
	ws.Bind("login", event.OnLogin)
	ws.Bind("register", event.OnRegister)
	ws.Bind("getguilds", event.OnGetGuilds)
	ws.Bind("message", event.OnMessage)
	ws.Bind("getmessages", event.OnGetMessages)
	ws.Bind("getchannels", event.OnGetChannels)
	ws.Bind("joinguild", event.OnJoinGuild)
	ws.Bind("createguild", event.OnCreateGuild)
	ws.Bind("leaveguild", event.OnLeaveGuild)
	ws.Bind("updateguildpicture", event.OnUpdateGuildPicture)
	ws.Bind("updateguildname", event.OnUpdateGuildName)
	ws.Bind("getinvites", event.OnGetInvites)
	ws.Bind("addchannel", event.OnAddChannel)
	ws.Bind("deletechannel", event.OnDeleteChannel)
	ws.Bind("deleteinvite", event.OnDeleteInvite)
	ws.Bind("createinvite", event.OnCreateInvite)
}