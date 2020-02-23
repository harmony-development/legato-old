package globals

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
	"time"
)

type (
	Event func(ws *Client, data map[string]interface{})

	EventBus map[string]Event

	Packet struct {
		Type string `json:"type"`
		Data map[string]interface{} `json:"data"`
	}

	Guild struct {
		Clients map[string][]*Client
		Owner string
	}

	Client struct {
		Connection *websocket.Conn
		EventBus map[string]Event
		Userid string
		Authed bool
		LastPong time.Time
		Out chan []byte
	}

	RESTClient struct {
		limiter *rate.Limiter
		lastReq time.Time
	}
)

func (bus EventBus) Bind(s string, event Event) {
	bus[s] = event
}

func (ws *Client) Send(p *Packet) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	ws.Out <- data
}

func GetRESTClient(userid string) *rate.Limiter {
	if RESTApiLimit[userid] == nil {
		RESTApiLimit[userid] = &RESTClient{
			limiter: rate.NewLimiter(rate.Every(3 * time.Second), 3),
			lastReq: time.Now(),
		}
	} else {
		RESTApiLimit[userid].lastReq = time.Now()
	}
	return RESTApiLimit[userid].limiter
}

func RateLimitCleanup() {
	for {
		time.Sleep(time.Minute)
		for userid, client := range RESTApiLimit {
			if time.Now().Sub(client.lastReq) > 3*time.Minute {
				delete(RESTApiLimit, userid)
			}
		}
	}
}

var Guilds = make(map[string]*Guild)
var RESTApiLimit = make(map[string]*RESTClient)
var Bus EventBus