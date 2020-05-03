package socket

import (
	"harmony-server/server/db"
	"harmony-server/server/http/socket/events"
	"harmony-server/server/http/socket/handling"
	"harmony-server/server/state"
	"time"
)

type Handler struct {
	DB *db.DB
	State *state.State
}

func NewHandler(deps struct{
	DB *db.DB
	State *state.State
}, bus handling.Bus) *Handler {
	var h Handler
	h.State = deps.State
	h.DB = deps.DB
	e := events.NewEvents(h.DB, h.State)
	bus.Bind(handling.NewEvent(e.Ping, "ping", 3 * time.Second, 10))
	bus.Bind(handling.NewEvent(e.Subscribe, "subscribe", 3 * time.Second, 10))
	return &h
}