package globals

import (
	"database/sql"
	"github.com/bluskript/harmony-server/socket"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	JwtSecret        string
	SocketServer     *socket.WebSocket
	DatabaseInstance *sql.DB
	Router           *mux.Router
}

var HarmonyServer = Server{
	SocketServer:     nil,
	DatabaseInstance: nil,
	Router:           nil,
}
