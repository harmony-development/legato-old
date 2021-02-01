package middleware

import (
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (m Middlewares) Validate(c echo.Context, meth *descriptorpb.MethodDescriptorProto, d *descriptorpb.FileDescriptorProto, h server.Handler) server.Handler {
	return func(c echo.Context, req protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error) {
		validator, ok := req.(interface{ Validate() error })
		if !ok {
			return h(c, req)
		}
		if err := validator.Validate(); err != nil {
			return nil, err
		}
		return h(c, req)
	}
}
