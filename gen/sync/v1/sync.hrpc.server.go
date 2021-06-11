package v1

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/websocket"
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func BindPB(obj interface{}, c echo.Context) error {
	buf, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	ct := c.Request().Header.Get("Content-Type")
	switch ct {
	case "application/hrpc", "application/octet-stream":
		if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
			return err
		}
	case "application/hrpc-json":
		if err = protojson.Unmarshal(buf, obj.(proto.Message)); err != nil {
			return err
		}
	}

	return nil
}

var Syncᐳv1ᐳsync *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x12sync/v1/sync.proto\x12\x10protocol.sync.v1\x1a\x1bgoogle/protobuf/empty.proto\"\xe5\x02\n\x05Event\x12e\n\x17user_removed_from_guild\x18\x01 \x01(\v2,.protocol.sync.v1.Event.UserRemovedFromGuildH\x00R\x14userRemovedFromGuild\x12Y\n\x13user_added_to_guild\x18\x02 \x01(\v2(.protocol.sync.v1.Event.UserAddedToGuildH\x00R\x10userAddedToGuild\x1aJ\n\x14UserRemovedFromGuild\x12\x17\n\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x19\n\bguild_id\x18\x02 \x01(\x04R\aguildId\x1aF\n\x10UserAddedToGuild\x12\x17\n\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x19\n\bguild_id\x18\x02 \x01(\x04R\aguildIdB\x06\n\x04kind\"A\n\x10PostEventRequest\x12-\n\x05event\x18\x01 \x01(\v2\x17.protocol.sync.v1.EventR\x05event\" \n\x03Ack\x12\x19\n\bevent_id\x18\x01 \x01(\x04R\aeventId\"O\n\x03Syn\x12\x19\n\bevent_id\x18\x01 \x01(\x04R\aeventId\x12-\n\x05event\x18\x02 \x01(\v2\x17.protocol.sync.v1.EventR\x05event2\x87\x01\n\x0ePostboxService\x12:\n\x04Pull\x12\x15.protocol.sync.v1.Ack\x1a\x15.protocol.sync.v1.Syn\"\x00(\x010\x01\x129\n\x04Push\x12\x17.protocol.sync.v1.Event\x1a\x16.google.protobuf.Empty\"\x00B3Z1github.com/harmony-development/legato/gen/sync/v1J\xe3\x1b\n\x06\x12\x04\x00\x00j\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n\b\n\x01\x02\x12\x03\x04\x00\x19\n\b\n\x01\b\x12\x03\x06\x00H\n\t\n\x02\b\v\x12\x03\x06\x00H\n\n\n\x02\x04\x00\x12\x04\b\x00\x16\x01\n\n\n\x03\x04\x00\x01\x12\x03\b\b\r\n\f\n\x04\x04\x00\x03\x00\x12\x04\t\x02\f\x03\n\f\n\x05\x04\x00\x03\x00\x01\x12\x03\t\n\x1e\n\r\n\x06\x04\x00\x03\x00\x02\x00\x12\x03\n\x04\x17\n\x0e\n\a\x04\x00\x03\x00\x02\x00\x05\x12\x03\n\x04\n\n\x0e\n\a\x04\x00\x03\x00\x02\x00\x01\x12\x03\n\v\x12\n\x0e\n\a\x04\x00\x03\x00\x02\x00\x03\x12\x03\n\x15\x16\n\r\n\x06\x04\x00\x03\x00\x02\x01\x12\x03\v\x04\x18\n\x0e\n\a\x04\x00\x03\x00\x02\x01\x05\x12\x03\v\x04\n\n\x0e\n\a\x04\x00\x03\x00\x02\x01\x01\x12\x03\v\v\x13\n\x0e\n\a\x04\x00\x03\x00\x02\x01\x03\x12\x03\v\x16\x17\n\f\n\x04\x04\x00\x03\x01\x12\x04\r\x02\x10\x03\n\f\n\x05\x04\x00\x03\x01\x01\x12\x03\r\n\x1a\n\r\n\x06\x04\x00\x03\x01\x02\x00\x12\x03\x0e\b\x1b\n\x0e\n\a\x04\x00\x03\x01\x02\x00\x05\x12\x03\x0e\b\x0e\n\x0e\n\a\x04\x00\x03\x01\x02\x00\x01\x12\x03\x0e\x0f\x16\n\x0e\n\a\x04\x00\x03\x01\x02\x00\x03\x12\x03\x0e\x19\x1a\n\r\n\x06\x04\x00\x03\x01\x02\x01\x12\x03\x0f\b\x1c\n\x0e\n\a\x04\x00\x03\x01\x02\x01\x05\x12\x03\x0f\b\x0e\n\x0e\n\a\x04\x00\x03\x01\x02\x01\x01\x12\x03\x0f\x0f\x17\n\x0e\n\a\x04\x00\x03\x01\x02\x01\x03\x12\x03\x0f\x1a\x1b\n\f\n\x04\x04\x00\b\x00\x12\x04\x12\x02\x15\x03\n\f\n\x05\x04\x00\b\x00\x01\x12\x03\x12\b\f\n\v\n\x04\x04\x00\x02\x00\x12\x03\x13\x045\n\f\n\x05\x04\x00\x02\x00\x06\x12\x03\x13\x04\x18\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\x13\x190\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\x1334\n\v\n\x04\x04\x00\x02\x01\x12\x03\x14\x04-\n\f\n\x05\x04\x00\x02\x01\x06\x12\x03\x14\x04\x14\n\f\n\x05\x04\x00\x02\x01\x01\x12\x03\x14\x15(\n\f\n\x05\x04\x00\x02\x01\x03\x12\x03\x14+,\n\n\n\x02\x04\x01\x12\x04\x18\x00\x1a\x01\n\n\n\x03\x04\x01\x01\x12\x03\x18\b\x18\n\v\n\x04\x04\x01\x02\x00\x12\x03\x19\x02\x12\n\f\n\x05\x04\x01\x02\x00\x06\x12\x03\x19\x02\a\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\x19\b\r\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\x19\x10\x11\n<\n\x02\x04\x02\x12\x04\x1d\x00\x1f\x01\x1a0 Acknowledgement of an event pulled using Pull.\n\n\n\n\x03\x04\x02\x01\x12\x03\x1d\b\v\n\v\n\x04\x04\x02\x02\x00\x12\x03\x1e\x02\x16\n\f\n\x05\x04\x02\x02\x00\x05\x12\x03\x1e\x02\b\n\f\n\x05\x04\x02\x02\x00\x01\x12\x03\x1e\t\x11\n\f\n\x05\x04\x02\x02\x00\x03\x12\x03\x1e\x14\x15\n:\n\x02\x04\x03\x12\x04\"\x00%\x01\x1a. A synchronisation message pulled using Pull.\n\n\n\n\x03\x04\x03\x01\x12\x03\"\b\v\n\v\n\x04\x04\x03\x02\x00\x12\x03#\x02\x16\n\f\n\x05\x04\x03\x02\x00\x05\x12\x03#\x02\b\n\f\n\x05\x04\x03\x02\x00\x01\x12\x03#\t\x11\n\f\n\x05\x04\x03\x02\x00\x03\x12\x03#\x14\x15\n\v\n\x04\x04\x03\x02\x01\x12\x03$\x02\x12\n\f\n\x05\x04\x03\x02\x01\x06\x12\x03$\x02\a\n\f\n\x05\x04\x03\x02\x01\x01\x12\x03$\b\r\n\f\n\x05\x04\x03\x02\x01\x03\x12\x03$\x10\x11\n\xb2\x13\n\x02\x06\x00\x12\x04g\x00j\x01\x1a\xa5\x13 # Postbox\n\n The postbox service forms the core of Harmony's server <-> server communications.\n\n It concerns the transfer of Events between servers, as well as ensuring reliable\n delivery of them.\n\n The semantics of events are documented in the event types. The postbox service\n is solely reliable for reliable pushing and pulling.\n\n ## Authorisation\n\n Requests are authorised using a JWT token in the Authorization HTTP header.\n\n The JWT token is signed using SHA-RSA-256 with the homeserver's private key,\n\n It contains the following fields, described using Go JSON semantics:\n ```\n Self string\n Time uint53\n ```\n\n Self is the server name of the server initiating the transaction. For Pull,\n this tells the server being connected to which homeservers' events it should send.\n For Push, this tells the server being connected to which homeservers' events it is\n receiving.\n\n Time is the UTC UNIX time in seconds of when the request is started. Servers should reject\n JWTs with a time too far from the current time, at their discretion. A recommended\n variance is 1 minute.\n\n ## Events\n\n In this section, we will use sender and recipient to refer to the servers\n sending the events and the server receiving the events respectively.\n\n When an event that a recipient would be interested in receiving occurs, the\n sender should check whether or not there is an active Pull by the receiver.\n If there is one, the server should dispatch the event to its queue as described\n later in this document.\n If there is not an active Pull by the receiver, the sender will attempt to Push\n to the receiver. If the Push RPC fails, the event will be dispatched to the\n sender's queue for the receiver.\n\n ### The Event Queue\n\n The event queue is an abstract data structure. It is filled by a sender when\n a Push fails.\n\n It is emptied by Pull requests. When the receiver initiates a Pull, the sender\n sends up to 100 Syns in sequential order before waiting on Acks. Events sent\n as a Syn but without an Ack are considered in-flight.\n\n An event will be taken out of flight if it is Acked by the receiver.\n\n If the Pull is cancelled or errors out before the sender receives an Ack for\n an event in-flight, it will be returned to the queue to be sent when the receiver\n performs another Pull.\n\n When an event is Acked and removed from flight, an older event from the queue should be\n sent.\n\n In essence, the queue is a LIFO stack. Newer events should be sent and acked before older events.\n\n\n\n\n\x03\x06\x00\x01\x12\x03g\b\x16\n\v\n\x04\x06\x00\x02\x00\x12\x03h\x02.\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03h\x06\n\n\f\n\x05\x06\x00\x02\x00\x05\x12\x03h\v\x11\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03h\x12\x15\n\f\n\x05\x06\x00\x02\x00\x06\x12\x03h &\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03h'*\n\v\n\x04\x06\x00\x02\x01\x12\x03i\x024\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03i\x06\n\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03i\v\x10\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03i\x1b0b\x06proto3")

	err := proto.Unmarshal(data, Syncᐳv1ᐳsync)
	if err != nil {
		panic(err)
	}
}

var PostboxServiceData *descriptorpb.ServiceDescriptorProto = new(descriptorpb.ServiceDescriptorProto)

func init() {
	data := []byte("\n\x0ePostboxService\x12:\n\x04Pull\x12\x15.protocol.sync.v1.Ack\x1a\x15.protocol.sync.v1.Syn\"\x00(\x010\x01\x129\n\x04Push\x12\x17.protocol.sync.v1.Event\x1a\x16.google.protobuf.Empty\"\x00")

	err := proto.Unmarshal(data, PostboxServiceData)
	if err != nil {
		panic(err)
	}
}

type PostboxServiceServer interface {
	Pull(ctx echo.Context, in chan *Ack, out chan *Syn)

	Push(ctx echo.Context, r *Event) (resp *empty.Empty, err error)
}

var PostboxServiceServerPullData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x04Pull\x12\x15.protocol.sync.v1.Ack\x1a\x15.protocol.sync.v1.Syn\"\x00(\x010\x01")

	err := proto.Unmarshal(data, PostboxServiceServerPullData)
	if err != nil {
		panic(err)
	}
}

var PostboxServiceServerPushData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x04Push\x12\x17.protocol.sync.v1.Event\x1a\x16.google.protobuf.Empty\"\x00")

	err := proto.Unmarshal(data, PostboxServiceServerPushData)
	if err != nil {
		panic(err)
	}
}

type PostboxServiceHandler struct {
	Server       PostboxServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     server.HandlerTransformer
	upgrader     websocket.Upgrader
}

func NewPostboxServiceHandler(s PostboxServiceServer) *PostboxServiceHandler {
	return &PostboxServiceHandler{
		Server: s,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
	}
}

func (h *PostboxServiceHandler) SetUnaryPre(s server.HandlerTransformer) {
	h.UnaryPre = s
}

func (h *PostboxServiceHandler) Routes() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{

		"/protocol.sync.v1.PostboxService/Pull": h.PullHandler,

		"/protocol.sync.v1.PostboxService/Push": h.PushHandler,
	}
}

func (h *PostboxServiceHandler) PullHandler(c echo.Context) error {

	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	defer ws.Close()

	in := make(chan *Ack, 100)

	out := make(chan *Syn, 100)

	h.Server.Pull(c, in, out)

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

	defer ws.Close()

	for {
		select {
		case data, ok := <-msgs:
			if !ok {
				close(in)
				close(out)
				return nil
			}

			item := new(Ack)
			switch c.Request().Header.Get("Content-Type") {
			case "application/hrpc-json":
				if err = protojson.Unmarshal(data, item); err != nil {
					close(in)
					close(out)
					c.Logger().Error(err)
					return nil
				}
			default:
				if err = proto.Unmarshal(data, item); err != nil {
					close(in)
					close(out)
					c.Logger().Error(err)
					return nil
				}
			}

			in <- item
		case msg, ok := <-out:
			if !ok {
				close(in)
				close(out)
				return nil
			}

			w, err := ws.NextWriter(websocket.BinaryMessage)
			if err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}

			var response []byte

			switch c.Request().Header.Get("Content-Type") {
			case "application/hrpc-json":
				response, err = protojson.Marshal(msg)
			default:
				response, err = proto.Marshal(msg)
			}

			if err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}

			if _, err := w.Write(response); err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}
			if err := w.Close(); err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}
		}

	}

}

func (h *PostboxServiceHandler) PushHandler(c echo.Context) error {

	requestProto := new(Event)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.Push(c, req.(*Event))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(PostboxServiceServerPushData, PostboxServiceData, Syncᐳv1ᐳsync, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}
