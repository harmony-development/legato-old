package socket

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	. "github.com/logrusorgru/aurora"
)

func startSocketServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(Red(err).Bold())
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println(Green("Socket connecting, ID : " + s.ID()).Bold())
		return nil
	})
}
