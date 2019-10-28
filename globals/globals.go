package globals

import (
	"github.com/bluskript/harmony-server/socket"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	SocketServer  *socket.WebSocket
	Collections   map[string]*mongo.Collection
	MongoInstance *mongo.Client
	Router        *mux.Router
}

var HarmonyServer = Server{
	SocketServer:  nil,
	Collections:   make(map[string]*mongo.Collection),
	MongoInstance: nil,
	Router:        nil,
}