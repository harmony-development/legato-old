package api

import (
	"net"
	"net/http"

	corev1 "github.com/harmony-development/legato/gen/core"
	foundationv1 "github.com/harmony-development/legato/gen/foundation"
	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/core"
	"github.com/harmony-development/legato/server/api/foundation"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/api/profile"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/sonyflake"
	"golang.org/x/sync/errgroup"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Dependencies struct {
	Logger      logger.ILogger
	DB          db.IHarmonyDB
	Sonyflake   *sonyflake.Sonyflake
	AuthManager *auth.Manager
	Config      *config.Config
}

// API contains the component of the server responsible for APIs
type API struct {
	Dependencies
	grpcServer        *grpc.Server
	grpcWebServer     *grpcweb.WrappedGrpcServer
	grpcWebHTTPServer *http.Server
	prometheusServer  *http.Server
	CoreKit           *core.Service
}

// New creates a new API instance
func New(deps Dependencies) *API {
	api := &API{
		Dependencies: deps,
	}
	m := middleware.New(middleware.Dependencies{
		Logger: deps.Logger,
		DB:     deps.DB,
	})
	api.grpcServer = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			m.HarmonyContextInterceptor,
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(m.RecoveryFunc)),
			grpc_prometheus.UnaryServerInterceptor,
			m.ErrorInterceptor,
			m.RateLimitInterceptor,
			m.ValidatorInterceptor,
			m.AuthInterceptor,
			m.LocationInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(m.RecoveryFunc)),
			grpc_prometheus.StreamServerInterceptor,
			m.HarmonyContextInterceptorStream,
			m.ErrorInterceptorStream,
			m.RateLimitStreamInterceptorStream,
		))
	api.grpcWebServer = grpcweb.WrapServer(api.grpcServer, grpcweb.WithOriginFunc(func(_ string) bool {
		return true
	}), grpcweb.WithWebsockets(true))
	api.grpcWebHTTPServer = &http.Server{
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			api.grpcWebServer.ServeHTTP(resp, req)
		}),
	}
	prometheusMux := http.NewServeMux()
	prometheusMux.Handle("/metrics", promhttp.Handler())
	api.prometheusServer = &http.Server{
		Handler: prometheusMux,
	}

	corev1.RegisterCoreServiceServer(api.grpcServer, core.New(&core.Dependencies{
		DB:        api.DB,
		Logger:    api.Logger,
		Sonyflake: api.Sonyflake,
	}).V1)
	profilev1.RegisterProfileServiceServer(api.grpcServer, &profile.New(profile.Dependencies{
		DB: api.DB,
	}).V1)
	foundationv1.RegisterFoundationServiceServer(api.grpcServer, foundation.New(&foundation.Dependencies{
		DB:          api.DB,
		Logger:      api.Logger,
		Sonyflake:   api.Sonyflake,
		AuthManager: api.AuthManager,
		Config:      api.Config,
	}))
	reflection.Register(api.grpcServer)
	grpc_prometheus.Register(api.grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	return api
}

// Start starts up the API on a specific port
func (api API) Start(grpcListener, grpcWebListener, prometheusListener net.Listener) error {
	errGrp := errgroup.Group{}

	errGrp.Go(func() error {
		err := api.grpcServer.Serve(grpcListener)
		api.Logger.CheckException(err)
		return err
	})
	errGrp.Go(func() error {
		err := api.grpcWebHTTPServer.Serve(grpcWebListener)
		api.Logger.CheckException(err)
		return err
	})
	errGrp.Go(func() error {
		err := api.prometheusServer.Serve(prometheusListener)
		api.Logger.CheckException(err)
		return err
	})

	return errGrp.Wait()
}
