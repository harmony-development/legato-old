package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/harmony-development/hrpc/server"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// UnaryRecoveryFunc recovers unary requests
func (m Middlewares) UnaryRecoveryFunc(meth *descriptorpb.MethodDescriptorProto, d *descriptorpb.FileDescriptorProto, h server.Handler) server.Handler {
	return func(c echo.Context, req protoreflect.ProtoMessage) (msg protoreflect.ProtoMessage, err error) {
		defer func() {
			if r := recover(); r != nil {
				m.Logger.Exception(fmt.Errorf("%+v", r))
				m.Logger.Exception(errors.New(string(debug.Stack())))
				err = echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerError)
			}
		}()
		return h(c, req)
	}
}
