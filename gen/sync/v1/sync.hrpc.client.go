package v1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type PostboxServiceClient struct {
	client    *http.Client
	serverURL string

	Header    http.Header
	HTTPProto string
	WSProto   string
}

func NewPostboxServiceClient(url string) *PostboxServiceClient {
	return &PostboxServiceClient{
		client:    &http.Client{},
		serverURL: url,
		Header:    http.Header{},
		HTTPProto: "https",
		WSProto:   "wss",
	}
}

func (client *PostboxServiceClient) Pull() (in chan<- *Ack, out <-chan *Syn, err error) {
	u := url.URL{Scheme: client.WSProto, Host: client.serverURL, Path: "/protocol.sync.v1.PostboxService/Pull"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), client.Header)
	if err != nil {
		return nil, nil, err
	}

	inC := make(chan *Ack)
	outC := make(chan *Syn)

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
					close(inC)
					close(outC)
					return
				}

				thing := new(Syn)
				err = proto.Unmarshal(msg, thing)
				if err != nil {
					return
				}

				outC <- thing
			case send, ok := <-inC:
				if !ok {
					close(outC)
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

func (client *PostboxServiceClient) Push(r *Event) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.sync.v1.PostboxService/Push", client.HTTPProto, client.serverURL), bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for k, v := range client.Header {
		req.Header[k] = v
	}
	req.Header.Add("content-type", "application/hrpc")
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}
