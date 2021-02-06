package v1

import (
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type VoiceServiceClient struct {
	client    *http.Client
	serverURL string

	Header    http.Header
	HTTPProto string
	WSProto   string
}

func NewVoiceServiceClient(url string) *VoiceServiceClient {
	return &VoiceServiceClient{
		client:    &http.Client{},
		serverURL: url,
		Header:    http.Header{},
		HTTPProto: "https",
		WSProto:   "wss",
	}
}

func (client *VoiceServiceClient) Connect() (in chan<- *ClientSignal, out <-chan *Signal, err error) {
	u := url.URL{Scheme: client.WSProto, Host: client.serverURL, Path: "/protocol.voice.v1.VoiceService/Connect"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), client.Header)
	if err != nil {
		return nil, nil, err
	}

	inC := make(chan *ClientSignal)
	outC := make(chan *Signal)

	go func() {
		defer c.Close()

		msgs := make(chan []byte)

		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					close(msgs)
					break
				}
				msgs <- message
			}
		}()

		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return
				}

				thing := new(Signal)
				err = proto.Unmarshal(msg, thing)
				if err != nil {
					return
				}

				outC <- thing
			case send, ok := <-inC:
				if !ok {
					return
				}

				data, err := proto.Marshal(send)
				if err != nil {
					return
				}

				err = c.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					return
				}
			}
		}
	}()

	return inC, outC, nil
}
