package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func CheckAuth(database db.IHarmonyDB, c context.Context) (uint64, error) {
	headers, exists := metadata.FromIncomingContext(c)
	if !exists {
		println("no header from incoming context")
		return 0, status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	authHeader, exists := headers["authorization"]
	if !exists {
		println("no authorization header")
		return 0, status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	session := authHeader[0]
	userID, err := database.SessionToUserID(session)
	if err != nil {
		println("bad session")
		return 0, status.Error(codes.NotFound, responses.InvalidSession)
	}
	return userID, nil
}

func (m Middlewares) AuthInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !GetRPCConfig(info.FullMethod).Auth {
		return handler(c, req)
	}

	ctx := c.(HarmonyContext)
	userID, err := CheckAuth(m.DB, c)
	if err != nil {
		return nil, err
	}
	ctx.UserID = userID
	return handler(ctx, req)
}
