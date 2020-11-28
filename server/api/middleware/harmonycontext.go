package middleware

import (
	"context"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

// HarmonyContext contains a custom context for passing data from middleware to handlers
type HarmonyContext struct {
	context.Context
	UserID    uint64
	UserRoles []uint64
	Limiter   *rate.Limiter
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

func (m Middlewares) HarmonyContextInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := HarmonyContext{
		Context: c,
	}
	return handler(ctx, req)
}

func (m Middlewares) HarmonyContextInterceptorStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, WrapServerStream(ss))
}

func WrapServerStream(stream grpc.ServerStream) HarmonyWrappedServerStream {
	if existing, ok := stream.(HarmonyWrappedServerStream); ok {
		return existing
	}
	return HarmonyWrappedServerStream{ServerStream: stream, WrappedContext: HarmonyContext{
		Context: stream.Context(),
	}}
}
