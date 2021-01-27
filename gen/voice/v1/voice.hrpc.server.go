package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"

type VoiceServiceServer interface {
	Connect(ctx context.Context, in chan *ClientSignal, out chan *Signal, headers http.Header)
}

type VoiceServiceHandler struct {
	Server       VoiceServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	upgrader     websocket.Upgrader
}

func NewVoiceServiceHandler(s VoiceServiceServer, errHandler func(err error, w http.ResponseWriter)) *VoiceServiceHandler {
	return &VoiceServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *VoiceServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.voice.v1.VoiceService/Connect":
		{
			var err error

			in := make(chan *ClientSignal)
			err = nil

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *Signal)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				msgs := make(chan []byte)

				go func() {
					for {
						_, message, err := ws.ReadMessage()
						if err != nil {
							close(msgs)
							break
						}
						msgs <- message
					}
				}()

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case data, ok := <-msgs:
						if !ok {
							return
						}

						item := new(ClientSignal)
						err = proto.Unmarshal(data, item)
						if err != nil {
							close(in)
							close(out)
							return
						}

						in <- item

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(in)

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(in)

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(in)

							close(out)
							return
						}
					}
				}
			}()

			h.Server.Connect(req.Context(), in, out, req.Header)
		}

	}
}
