package middleware

import (
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Auth checks if the user is authenticated using the auth interface.
func (i *Middlewares) Auth(
	meth *descriptorpb.MethodDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto,
	d *descriptorpb.FileDescriptorProto,
	h server.Handler,
) server.Handler {
	return func(c echo.Context, req protoreflect.ProtoMessage) (err error) {
		// The Sec-Websocket-Protocol header is used to store the session token.
	}
}
