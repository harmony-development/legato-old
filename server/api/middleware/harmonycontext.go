package middleware

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// HarmonyContext contains a custom context for passing data from middleware to handlers
type HarmonyContext struct {
	echo.Context
	MethodDesc *descriptor.MethodDescriptorProto
	UserID     uint64
	UserRoles  []uint64
	IsOwner    bool
	Limiter    *rate.Limiter
}

type IHarmonyWrappedServerStream interface {
	GetWrappedContext() HarmonyContext
}

type HarmonyWrappedServerStream struct {
	grpc.ServerStream
	WrappedContext HarmonyContext
}

func (ss HarmonyWrappedServerStream) GetWrappedContext() HarmonyContext {
	return ss.WrappedContext
}

func (m Middlewares) HarmonyContextInterceptor(meth *descriptorpb.MethodDescriptorProto, d *descriptorpb.FileDescriptorProto, h server.Handler) server.Handler {
	return func(c echo.Context, req protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error) {
		return h(HarmonyContext{
			Context: c,
		}, req)
	}
}
