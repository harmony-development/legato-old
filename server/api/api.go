package api

import (
	corev1 "harmony-server/gen/core"
	profilev1 "harmony-server/gen/profile"
	"harmony-server/server/api/core"
	"harmony-server/server/api/profile"
	"net"

	"google.golang.org/grpc"
)

// API contains the component of the server responsible for APIs
type API struct {
	grpcServer *grpc.Server
	CoreKit    *core.Service
}

// New creates a new API instance
func New() *API {
	return &API{
		grpcServer: grpc.NewServer(),
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
