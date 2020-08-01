package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) ErrorInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)

	resp, err := handler(ctx, req)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}
		return nil, m.Logger.ErrorResponse(codes.Unknown, err, responses.UnknownError)
	}
	return resp, err
}
