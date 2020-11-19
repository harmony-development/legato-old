package middleware

import (
	"context"
	"fmt"

	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (m Middlewares) AuthInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)

	if !GetRPCConfig(info.FullMethod).Auth {
		return handler(ctx, req)
	}

	userid, err := AuthHandler(m.DB, ctx)

	if err != nil {
		return nil, err
	}

	ctx.UserID = userid

	return handler(ctx, req)
}

func AuthHandler(database db.IHarmonyDB, c context.Context) (uint64, error) {
	headers, exists := metadata.FromIncomingContext(c)
	if !exists {
		println("no header from incoming context")
		return 0, status.Error(codes.Unauthenticated, responses.InvalidSession)
	}
	authHeader, exists := headers["auth"]
	if !exists {
		fmt.Println(headers)
		println("no auth header")
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
