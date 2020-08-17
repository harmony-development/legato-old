package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (m Middlewares) AuthInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)
	if err := m.authHandler(info.FullMethod, ctx); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (m Middlewares) AuthInterceptorStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	wrappedStream := ss.(HarmonyWrappedServerStream)
	if err := m.authHandler(info.FullMethod, wrappedStream.WrappedContext); err != nil {
		return err
	}
	return handler(srv, wrappedStream)
}

func (m Middlewares) authHandler(fullMethod string, ctx HarmonyContext) error {
	if !GetRPCConfig(fullMethod).Auth {
		return nil
	}

	headers, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		println("no header from incoming context")
		return status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	authHeader, exists := headers["auth"]
	if !exists {
		println("no auth header")
		return status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	session := authHeader[0]
	userID, err := m.DB.SessionToUserID(session)
	if err != nil {
		println("bad session")
		return status.Error(codes.NotFound, responses.InvalidSession)
	}
	ctx.UserID = userID
	return nil
}
