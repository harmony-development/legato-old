package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (m Middlewares) AuthInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)
	if !rpcConfigs[info.FullMethod].Auth {
		return handler(ctx, req)
	}
	headers, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return nil, status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	authHeader, exists := headers["authorization"]
	if !exists {
		return nil, status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	session := authHeader[0]
	userID, err := m.DB.SessionToUserID(session)
	if err != nil {
		return nil, status.Error(codes.NotFound, responses.InvalidSession)
	}
	ctx.UserID = userID
	return handler(ctx, req)
}
