package v1

import (
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/enriquebris/goconcurrentqueue"
	"github.com/golang/protobuf/proto"
	syncv1 "github.com/harmony-development/legato/gen/sync/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB          types.IHarmonyDB
	Logger      logger.ILogger
	Sonyflake   *sonyflake.Sonyflake
	Config      *config.Config
	AuthManager *auth.Manager
	Middleware  *middleware.Middlewares

	EventDispatcher func(string, *syncv1.Event)
}

// V1 contains the gRPC handler for v1
type V1 struct {
	Dependencies

	Queues map[string]*goconcurrentqueue.FIFO
}

type Token struct {
	jwt.StandardClaims
	Self string
	Time int64
}

func (s *V1) PersistQueue(host string) error {
	q, ok := s.Queues[host]
	if !ok {
		return nil
	}

	ln := q.GetLen()

	ev := []*syncv1.Event{}

	for i := 0; i < ln; i++ {
		v, _ := q.Get(i)
		ev = append(ev, v.(*syncv1.Event))
	}

	datas := [][]byte{}

	for _, it := range ev {
		data, err := proto.Marshal(it)
		if err != nil {
			return err
		}
		datas = append(datas, data)
	}

	data, err := json.Marshal(datas)
	if err != nil {
		return err
	}

	return s.DB.SetHostQueue(host, data)
}

func (s *V1) GetQueue(host string) error {
	if _, ok := s.Queues[host]; !ok {
		s.Queues[host] = goconcurrentqueue.NewFIFO()
	}

	q := s.Queues[host]

	queue, err := s.DB.GetHostQueue(host)
	if err != nil {
		return err
	}

	datas := [][]byte{}

	if err := json.Unmarshal(queue, &datas); err != nil {
		return err
	}

	for _, item := range datas {
		msg := new(syncv1.Event)

		err := proto.Unmarshal(item, msg)
		if err != nil {
			return err
		}

		q.Enqueue(msg)
	}

	return nil
}

func (s *V1) validateAuthorization(auth string) (string, error) {
	t, err := jwt.ParseWithClaims(auth, &Token{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Invalid signing method: %v", t.Header["alg"])
		}
		pem, err := s.AuthManager.GetPublicKey(t.Claims.(*Token).Self)
		if err != nil {
			return nil, err
		}
		pubKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, err
		}
		return pubKey, nil
	})
	if err != nil {
		return "", err
	}
	return t.Claims.(*Token).Self, nil
}

func (s *V1) Pull(c echo.Context, in chan *syncv1.Ack, out chan *syncv1.Syn) {
	from, err := s.validateAuthorization(c.Request().Header.Get("Authorization"))
	if err != nil {
		s.Logger.Warn("Failed Pull request: %s", err)
	}

	if _, ok := s.Queues[from]; !ok {
		s.Queues[from] = goconcurrentqueue.NewFIFO()
		s.GetQueue(from)
	}

	id := uint64(0)
	for {
		id++

		item, _ := s.Queues[from].DequeueOrWaitForNextElement()
		s.PersistQueue(from)

		out <- &syncv1.Syn{EventId: id, Event: item.(*syncv1.Event)}
		ack := <-in

		if id != ack.EventId {
			close(in)
			close(out)
		}
	}
}

func (s *V1) Push(c echo.Context, r *syncv1.Event) (*emptypb.Empty, error) {
	from, err := s.validateAuthorization(c.Request().Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	s.EventDispatcher(from, r)
	return &emptypb.Empty{}, nil
}
