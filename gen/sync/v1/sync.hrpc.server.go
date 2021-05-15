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
	data := []byte("\n\x12sync/v1/sync.proto\x12\x10protocol.sync.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x19google/protobuf/any.proto\"7\n\vSyncRequest\x12\x14\n\x05token\x18\x01 \x01(\tR\x05token\x12\x12\n\x04host\x18\x02 \x01(\tR\x04host\":\n\fPostBoxEvent\x12*\n\x05event\x18\x01 \x01(\v2\x14.google.protobuf.AnyR\x05event\"\x8a\x01\n\x10PostEventRequest\x12@\n\fsync_request\x18\x01 \x01(\v2\x1d.protocol.sync.v1.SyncRequestR\vsyncRequest\x124\n\x05event\x18\x02 \x01(\v2\x1e.protocol.sync.v1.PostBoxEventR\x05event2\xa6\x01\n\x0ePostboxService\x12I\n\x04Sync\x12\x1d.protocol.sync.v1.SyncRequest\x1a\x1e.protocol.sync.v1.PostBoxEvent\"\x000\x01\x12I\n\tPostEvent\x12\".protocol.sync.v1.PostEventRequest\x1a\x16.google.protobuf.Empty\"\x00B3Z1github.com/harmony-development/legato/gen/sync/v1J\xb5\x04\n\x06\x12\x04\x00\x00\x18\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n\t\n\x02\x03\x01\x12\x03\x03\x00#\n\b\n\x01\x02\x12\x03\x05\x00\x19\n\b\n\x01\b\x12\x03\a\x00H\n\t\n\x02\b\v\x12\x03\a\x00H\n\n\n\x02\x04\x00\x12\x04\t\x00\f\x01\n\n\n\x03\x04\x00\x01\x12\x03\t\b\x13\n\v\n\x04\x04\x00\x02\x00\x12\x03\n\x02\x13\n\f\n\x05\x04\x00\x02\x00\x05\x12\x03\n\x02\b\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\n\t\x0e\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\n\x11\x12\n\v\n\x04\x04\x00\x02\x01\x12\x03\v\x02\x12\n\f\n\x05\x04\x00\x02\x01\x05\x12\x03\v\x02\b\n\f\n\x05\x04\x00\x02\x01\x01\x12\x03\v\t\r\n\f\n\x05\x04\x00\x02\x01\x03\x12\x03\v\x10\x11\n\t\n\x02\x04\x01\x12\x03\x0e\x007\n\n\n\x03\x04\x01\x01\x12\x03\x0e\b\x14\n\v\n\x04\x04\x01\x02\x00\x12\x03\x0e\x175\n\f\n\x05\x04\x01\x02\x00\x06\x12\x03\x0e\x17*\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\x0e+0\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\x0e34\n\n\n\x02\x04\x02\x12\x04\x10\x00\x13\x01\n\n\n\x03\x04\x02\x01\x12\x03\x10\b\x18\n\v\n\x04\x04\x02\x02\x00\x12\x03\x11\x02\x1f\n\f\n\x05\x04\x02\x02\x00\x06\x12\x03\x11\x02\r\n\f\n\x05\x04\x02\x02\x00\x01\x12\x03\x11\x0e\x1a\n\f\n\x05\x04\x02\x02\x00\x03\x12\x03\x11\x1d\x1e\n\v\n\x04\x04\x02\x02\x01\x12\x03\x12\x02\x19\n\f\n\x05\x04\x02\x02\x01\x06\x12\x03\x12\x02\x0e\n\f\n\x05\x04\x02\x02\x01\x01\x12\x03\x12\x0f\x14\n\f\n\x05\x04\x02\x02\x01\x03\x12\x03\x12\x17\x18\n\n\n\x02\x06\x00\x12\x04\x15\x00\x18\x01\n\n\n\x03\x06\x00\x01\x12\x03\x15\b\x16\n\v\n\x04\x06\x00\x02\x00\x12\x03\x16\x028\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03\x16\x06\n\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03\x16\v\x16\n\f\n\x05\x06\x00\x02\x00\x06\x12\x03\x16!'\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03\x16(4\n\v\n\x04\x06\x00\x02\x01\x12\x03\x17\x02D\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03\x17\x06\x0f\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03\x17\x10 \n\f\n\x05\x06\x00\x02\x01\x03\x12\x03\x17+@b\x06proto3")

	err := proto.Unmarshal(data, Syncᐳv1ᐳsync)
	if err != nil {
		panic(err)
	}
}

var PostboxServiceData *descriptorpb.ServiceDescriptorProto = new(descriptorpb.ServiceDescriptorProto)

func init() {
	data := []byte("\n\x0ePostboxService\x12I\n\x04Sync\x12\x1d.protocol.sync.v1.SyncRequest\x1a\x1e.protocol.sync.v1.PostBoxEvent\"\x000\x01\x12I\n\tPostEvent\x12\".protocol.sync.v1.PostEventRequest\x1a\x16.google.protobuf.Empty\"\x00")

	err := proto.Unmarshal(data, PostboxServiceData)
	if err != nil {
		panic(err)
	}
}

type PostboxServiceServer interface {
	Sync(ctx echo.Context, r *SyncRequest, out chan *PostBoxEvent)

	PostEvent(ctx echo.Context, r *PostEventRequest) (resp *empty.Empty, err error)
}

var PostboxServiceServerSyncData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x04Sync\x12\x1d.protocol.sync.v1.SyncRequest\x1a\x1e.protocol.sync.v1.PostBoxEvent\"\x000\x01")

	err := proto.Unmarshal(data, PostboxServiceServerSyncData)
	if err != nil {
		panic(err)
	}
}

var PostboxServiceServerPostEventData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\tPostEvent\x12\".protocol.sync.v1.PostEventRequest\x1a\x16.google.protobuf.Empty\"\x00")

	err := proto.Unmarshal(data, PostboxServiceServerPostEventData)
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

		"/protocol.sync.v1.PostboxService/Sync": h.SyncHandler,

		"/protocol.sync.v1.PostboxService/PostEvent": h.PostEventHandler,
	}
}

func (h *PostboxServiceHandler) SyncHandler(c echo.Context) error {

	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), map[string][]string{
		"Sec-WebSocket-Protocol": {"harmony"},
	})
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	defer ws.Close()

	in := new(SyncRequest)
	_, message, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	if err := proto.Unmarshal(message, in); err != nil {
		c.Logger().Error(err)
		return nil
	}
	switch c.Request().Header.Get("Content-Type") {
	case "application/hrpc-json":
		if err = protojson.Unmarshal(message, in); err != nil {
			return err
		}
	default:
		if err = proto.Unmarshal(message, in); err != nil {
			return err
		}
	}

	out := make(chan *PostBoxEvent, 100)

	h.Server.Sync(c, in, out)

	defer ws.Close()

	for msg := range out {

		w, err := ws.NextWriter(websocket.BinaryMessage)
		if err != nil {

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

			close(out)
			c.Logger().Error(err)
			return nil
		}

		if _, err := w.Write(response); err != nil {

			close(out)
			c.Logger().Error(err)
			return nil
		}
		if err := w.Close(); err != nil {

			close(out)
			c.Logger().Error(err)
			return nil
		}
	}

	return nil

}

func (h *PostboxServiceHandler) PostEventHandler(c echo.Context) error {

	requestProto := new(PostEventRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.PostEvent(c, req.(*PostEventRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(PostboxServiceServerPostEventData, PostboxServiceData, Syncᐳv1ᐳsync, invoker)
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
