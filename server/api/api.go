package api

import (
	"net"

	corev1 "github.com/harmony-development/legato/gen/core"
	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/core"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/api/profile"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

// API contains the component of the server responsible for APIs
type API struct {
	grpcServer *grpc.Server
	CoreKit    *core.Service
}

// New creates a new API instance
func New() *API {
	m := middleware.Middlewares{}
	return &API{
		grpcServer: grpc.NewServer(grpc_middleware.WithUnaryServerChain(
			m.HarmonyContextInterceptor,
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(m.RecoveryFunc)),
			m.RateLimitInterceptor,
		)),
	}
}

// Start starts up the API on a specific port
func (api API) Start(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	corev1.RegisterCoreServiceServer(api.grpcServer, core.New(&core.Dependencies{}).V1)
	profilev1.RegisterProfileServiceServer(api.grpcServer, profile.New(&profile.Dependencies{}).V1)
	return api.grpcServer.Serve(lis)
}
