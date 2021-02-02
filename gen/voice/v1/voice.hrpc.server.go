package v1

import "github.com/labstack/echo/v4"
import "io/ioutil"
import "net/http"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"
import "google.golang.org/protobuf/types/descriptorpb"
import "github.com/harmony-development/hrpc/server"

func BindPB(obj interface{}, c echo.Context) error {
	buf, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return err
	}

	return nil
}

var Voiceᐳv1ᐳvoice *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x14voice/v1/voice.proto\x12\x11protocol.voice.v1\x1a\x1bgoogle/protobuf/empty.proto\"\xf1\x01\n\fClientSignal\x12@\n\x06answer\x18\x01 \x01(\v2&.protocol.voice.v1.ClientSignal.AnswerH\x00R\x06answer\x12I\n\tcandidate\x18\x02 \x01(\v2).protocol.voice.v1.ClientSignal.CandidateH\x00R\tcandidate\x1a \n\x06Answer\x12\x16\n\x06answer\x18\x01 \x01(\tR\x06answer\x1a)\n\tCandidate\x12\x1c\n\tcandidate\x18\x01 \x01(\tR\tcandidateB\a\n\x05event\"\xdf\x01\n\x06Signal\x12F\n\tcandidate\x18\x01 \x01(\v2&.protocol.voice.v1.Signal.ICECandidateH\x00R\tcandidate\x127\n\x05offer\x18\x02 \x01(\v2\x1f.protocol.voice.v1.Signal.OfferH\x00R\x05offer\x1a,\n\fICECandidate\x12\x1c\n\tcandidate\x18\x01 \x01(\tR\tcandidate\x1a\x1d\n\x05Offer\x12\x14\n\x05offer\x18\x01 \x01(\tR\x05offerB\a\n\x05event2[\n\fVoiceService\x12K\n\aConnect\x12\x1f.protocol.voice.v1.ClientSignal\x1a\x19.protocol.voice.v1.Signal\"\x00(\x010\x01B4Z2github.com/harmony-development/legato/gen/voice/v1J\xd3\x06\n\x06\x12\x04\x00\x00\x1e\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n\b\n\x01\x02\x12\x03\x04\x00\x1a\n\b\n\x01\b\x12\x03\x06\x00I\n\t\n\x02\b\v\x12\x03\x06\x00I\n\n\n\x02\x04\x00\x12\x04\b\x00\x10\x01\n\n\n\x03\x04\x00\x01\x12\x03\b\b\x14\n\v\n\x04\x04\x00\x03\x00\x12\x03\t\x02'\n\f\n\x05\x04\x00\x03\x00\x01\x12\x03\t\n\x10\n\r\n\x06\x04\x00\x03\x00\x02\x00\x12\x03\t\x13%\n\x0e\n\a\x04\x00\x03\x00\x02\x00\x05\x12\x03\t\x13\x19\n\x0e\n\a\x04\x00\x03\x00\x02\x00\x01\x12\x03\t\x1a \n\x0e\n\a\x04\x00\x03\x00\x02\x00\x03\x12\x03\t#$\n\v\n\x04\x04\x00\x03\x01\x12\x03\n\x02-\n\f\n\x05\x04\x00\x03\x01\x01\x12\x03\n\n\x13\n\r\n\x06\x04\x00\x03\x01\x02\x00\x12\x03\n\x16+\n\x0e\n\a\x04\x00\x03\x01\x02\x00\x05\x12\x03\n\x16\x1c\n\x0e\n\a\x04\x00\x03\x01\x02\x00\x01\x12\x03\n\x1d&\n\x0e\n\a\x04\x00\x03\x01\x02\x00\x03\x12\x03\n)*\n\f\n\x04\x04\x00\b\x00\x12\x04\f\x02\x0f\x03\n\f\n\x05\x04\x00\b\x00\x01\x12\x03\f\b\r\n\v\n\x04\x04\x00\x02\x00\x12\x03\r\x04\x16\n\f\n\x05\x04\x00\x02\x00\x06\x12\x03\r\x04\n\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\r\v\x11\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\r\x14\x15\n\v\n\x04\x04\x00\x02\x01\x12\x03\x0e\x04\x1c\n\f\n\x05\x04\x00\x02\x01\x06\x12\x03\x0e\x04\r\n\f\n\x05\x04\x00\x02\x01\x01\x12\x03\x0e\x0e\x17\n\f\n\x05\x04\x00\x02\x01\x03\x12\x03\x0e\x1a\x1b\n\n\n\x02\x04\x01\x12\x04\x12\x00\x1a\x01\n\n\n\x03\x04\x01\x01\x12\x03\x12\b\x0e\n\v\n\x04\x04\x01\x03\x00\x12\x03\x13\x020\n\f\n\x05\x04\x01\x03\x00\x01\x12\x03\x13\n\x16\n\r\n\x06\x04\x01\x03\x00\x02\x00\x12\x03\x13\x19.\n\x0e\n\a\x04\x01\x03\x00\x02\x00\x05\x12\x03\x13\x19\x1f\n\x0e\n\a\x04\x01\x03\x00\x02\x00\x01\x12\x03\x13 )\n\x0e\n\a\x04\x01\x03\x00\x02\x00\x03\x12\x03\x13,-\n\v\n\x04\x04\x01\x03\x01\x12\x03\x14\x02%\n\f\n\x05\x04\x01\x03\x01\x01\x12\x03\x14\n\x0f\n\r\n\x06\x04\x01\x03\x01\x02\x00\x12\x03\x14\x12#\n\x0e\n\a\x04\x01\x03\x01\x02\x00\x05\x12\x03\x14\x12\x18\n\x0e\n\a\x04\x01\x03\x01\x02\x00\x01\x12\x03\x14\x19\x1e\n\x0e\n\a\x04\x01\x03\x01\x02\x00\x03\x12\x03\x14!\"\n\f\n\x04\x04\x01\b\x00\x12\x04\x16\x02\x19\x03\n\f\n\x05\x04\x01\b\x00\x01\x12\x03\x16\b\r\n\v\n\x04\x04\x01\x02\x00\x12\x03\x17\x04\x1f\n\f\n\x05\x04\x01\x02\x00\x06\x12\x03\x17\x04\x10\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\x17\x11\x1a\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\x17\x1d\x1e\n\v\n\x04\x04\x01\x02\x01\x12\x03\x18\x04\x14\n\f\n\x05\x04\x01\x02\x01\x06\x12\x03\x18\x04\t\n\f\n\x05\x04\x01\x02\x01\x01\x12\x03\x18\n\x0f\n\f\n\x05\x04\x01\x02\x01\x03\x12\x03\x18\x12\x13\n\n\n\x02\x06\x00\x12\x04\x1c\x00\x1e\x01\n\n\n\x03\x06\x00\x01\x12\x03\x1c\b\x14\n\v\n\x04\x06\x00\x02\x00\x12\x03\x1d\x02=\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03\x1d\x06\r\n\f\n\x05\x06\x00\x02\x00\x05\x12\x03\x1d\x0e\x14\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03\x1d\x15!\n\f\n\x05\x06\x00\x02\x00\x06\x12\x03\x1d,2\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03\x1d39b\x06proto3")

	err := proto.Unmarshal(data, Voiceᐳv1ᐳvoice)
	if err != nil {
		panic(err)
	}
}

type VoiceServiceServer interface {
	Connect(ctx echo.Context, in chan *ClientSignal, out chan *Signal)
}

var VoiceServiceServerConnectData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\aConnect\x12\x1f.protocol.voice.v1.ClientSignal\x1a\x19.protocol.voice.v1.Signal\"\x00(\x010\x01")

	err := proto.Unmarshal(data, VoiceServiceServerConnectData)
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
	}
}

func (h *VoiceServiceHandler) ConnectHandler(c echo.Context) error {

	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	in := make(chan *ClientSignal)
	err = nil

	out := make(chan *Signal)

	h.Server.Connect(c, in, out)

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
				continue
			}

			item := new(ClientSignal)
			if err := proto.Unmarshal(data, item); err != nil {
				close(in)
				close(out)
				return err
			}

			in <- item

		case msg, ok := <-out:
			if !ok {
				continue
			}

			w, err := ws.NextWriter(websocket.BinaryMessage)
			if err != nil {

				close(in)

				close(out)
				return err
			}

			response, err := proto.Marshal(msg)
			if err != nil {

				close(in)

				close(out)
				return err
			}

			w.Write(response)
			if err := w.Close(); err != nil {

				close(in)

				close(out)
				return err
			}
		}
	}

}
