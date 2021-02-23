package v1

import (
	"io/ioutil"
	"net/http"

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

var Voiceᐳv1ᐳvoice *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x14voice/v1/voice.proto\x12\x11protocol.voice.v1\x1a\x1bgoogle/protobuf/empty.proto\"\x85\x01\n\x06Signal\x12%\n\rice_candidate\x18\x01 \x01(\tH\x00R\ficeCandidate\x12K\n\x14renegotiation_needed\x18\x02 \x01(\v2\x16.google.protobuf.EmptyH\x00R\x13renegotiationNeededB\a\n\x05event\"E\n\x0eConnectRequest\x12\x1d\n\nchannel_id\x18\x01 \x01(\x04R\tchannelId\x12\x14\n\x05offer\x18\x02 \x01(\tR\x05offer\")\n\x0fConnectResponse\x12\x16\n\x06answer\x18\x01 \x01(\tR\x06answer\"3\n\x12StreamStateRequest\x12\x1d\n\nchannel_id\x18\x01 \x01(\x04R\tchannelId2\xb7\x01\n\fVoiceService\x12R\n\aConnect\x12!.protocol.voice.v1.ConnectRequest\x1a\".protocol.voice.v1.ConnectResponse\"\x00\x12S\n\vStreamState\x12%.protocol.voice.v1.StreamStateRequest\x1a\x19.protocol.voice.v1.Signal\"\x000\x01B4Z2github.com/harmony-development/legato/gen/voice/v1J\x94\x05\n\x06\x12\x04\x00\x00\x1a\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n\b\n\x01\x02\x12\x03\x04\x00\x1a\n\b\n\x01\b\x12\x03\x06\x00I\n\t\n\x02\b\v\x12\x03\x06\x00I\n\n\n\x02\x04\x00\x12\x04\b\x00\r\x01\n\n\n\x03\x04\x00\x01\x12\x03\b\b\x0e\n\f\n\x04\x04\x00\b\x00\x12\x04\t\x02\f\x03\n\f\n\x05\x04\x00\b\x00\x01\x12\x03\t\b\r\n\v\n\x04\x04\x00\x02\x00\x12\x03\n\x04\x1d\n\f\n\x05\x04\x00\x02\x00\x05\x12\x03\n\x04\n\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\n\v\x18\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\n\x1b\x1c\n\v\n\x04\x04\x00\x02\x01\x12\x03\v\x043\n\f\n\x05\x04\x00\x02\x01\x06\x12\x03\v\x04\x19\n\f\n\x05\x04\x00\x02\x01\x01\x12\x03\v\x1a.\n\f\n\x05\x04\x00\x02\x01\x03\x12\x03\v12\n\n\n\x02\x04\x01\x12\x04\x0f\x00\x12\x01\n\n\n\x03\x04\x01\x01\x12\x03\x0f\b\x16\n\v\n\x04\x04\x01\x02\x00\x12\x03\x10\x02\x18\n\f\n\x05\x04\x01\x02\x00\x05\x12\x03\x10\x02\b\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\x10\t\x13\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\x10\x16\x17\n\v\n\x04\x04\x01\x02\x01\x12\x03\x11\x02\x13\n\f\n\x05\x04\x01\x02\x01\x05\x12\x03\x11\x02\b\n\f\n\x05\x04\x01\x02\x01\x01\x12\x03\x11\t\x0e\n\f\n\x05\x04\x01\x02\x01\x03\x12\x03\x11\x11\x12\n\t\n\x02\x04\x02\x12\x03\x13\x00.\n\n\n\x03\x04\x02\x01\x12\x03\x13\b\x17\n\v\n\x04\x04\x02\x02\x00\x12\x03\x13\x1a,\n\f\n\x05\x04\x02\x02\x00\x05\x12\x03\x13\x1a \n\f\n\x05\x04\x02\x02\x00\x01\x12\x03\x13!'\n\f\n\x05\x04\x02\x02\x00\x03\x12\x03\x13*+\n\t\n\x02\x04\x03\x12\x03\x15\x005\n\n\n\x03\x04\x03\x01\x12\x03\x15\b\x1a\n\v\n\x04\x04\x03\x02\x00\x12\x03\x15\x1d3\n\f\n\x05\x04\x03\x02\x00\x05\x12\x03\x15\x1d#\n\f\n\x05\x04\x03\x02\x00\x01\x12\x03\x15$.\n\f\n\x05\x04\x03\x02\x00\x03\x12\x03\x1512\n\n\n\x02\x06\x00\x12\x04\x17\x00\x1a\x01\n\n\n\x03\x06\x00\x01\x12\x03\x17\b\x14\n\v\n\x04\x06\x00\x02\x00\x12\x03\x18\x02:\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03\x18\x06\r\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03\x18\x0e\x1c\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03\x18'6\n\v\n\x04\x06\x00\x02\x01\x12\x03\x19\x02@\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03\x19\x06\x11\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03\x19\x12$\n\f\n\x05\x06\x00\x02\x01\x06\x12\x03\x19/5\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03\x196<b\x06proto3")

	err := proto.Unmarshal(data, Voiceᐳv1ᐳvoice)
	if err != nil {
		panic(err)
	}
}

var VoiceServiceData *descriptorpb.ServiceDescriptorProto = new(descriptorpb.ServiceDescriptorProto)

func init() {
	data := []byte("\n\fVoiceService\x12R\n\aConnect\x12!.protocol.voice.v1.ConnectRequest\x1a\".protocol.voice.v1.ConnectResponse\"\x00\x12S\n\vStreamState\x12%.protocol.voice.v1.StreamStateRequest\x1a\x19.protocol.voice.v1.Signal\"\x000\x01")

	err := proto.Unmarshal(data, VoiceServiceData)
	if err != nil {
		panic(err)
	}
}

type VoiceServiceServer interface {
	Connect(ctx echo.Context, r *ConnectRequest) (resp *ConnectResponse, err error)

	StreamState(ctx echo.Context, r *StreamStateRequest, out chan *Signal)
}

var VoiceServiceServerConnectData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\aConnect\x12!.protocol.voice.v1.ConnectRequest\x1a\".protocol.voice.v1.ConnectResponse\"\x00")

	err := proto.Unmarshal(data, VoiceServiceServerConnectData)
	if err != nil {
		panic(err)
	}
}

var VoiceServiceServerStreamStateData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vStreamState\x12%.protocol.voice.v1.StreamStateRequest\x1a\x19.protocol.voice.v1.Signal\"\x000\x01")

	err := proto.Unmarshal(data, VoiceServiceServerStreamStateData)
	if err != nil {
		panic(err)
	}
}

type VoiceServiceHandler struct {
	Server       VoiceServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     server.HandlerTransformer
	upgrader     websocket.Upgrader
}

func NewVoiceServiceHandler(s VoiceServiceServer) *VoiceServiceHandler {
	return &VoiceServiceHandler{
		Server: s,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *VoiceServiceHandler) SetUnaryPre(s server.HandlerTransformer) {
	h.UnaryPre = s
}

func (h *VoiceServiceHandler) Routes() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{

		"/protocol.voice.v1.VoiceService/Connect": h.ConnectHandler,

		"/protocol.voice.v1.VoiceService/StreamState": h.StreamStateHandler,
	}
}

func (h *VoiceServiceHandler) ConnectHandler(c echo.Context) error {

	requestProto := new(ConnectRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.Connect(c, req.(*ConnectRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(VoiceServiceServerConnectData, VoiceServiceData, Voiceᐳv1ᐳvoice, invoker)
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

func (h *VoiceServiceHandler) StreamStateHandler(c echo.Context) error {

	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	defer ws.Close()

	in := new(StreamStateRequest)
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

	out := make(chan *Signal, 100)

	h.Server.StreamState(c, in, out)

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
