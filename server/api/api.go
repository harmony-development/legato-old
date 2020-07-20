package api

import (
	"net"
	"net/http"
	"strconv"

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
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

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
		))
	api.grpcWebServer = grpcweb.WrapServer(api.grpcServer)
	api.grpcWebHTTPServer = &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				api.grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if api.grpcWebServer.IsGrpcWebRequest(r) {
					api.grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	prometheusMux := http.NewServeMux()
	prometheusMux.Handle("/metrics", promhttp.Handler())
	api.prometheusServer = &http.Server{
		Handler: prometheusMux,
	}
	return api
}

// Start starts up the API on a specific port
func (api API) Start(cb chan error, port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	webLis, err := net.Listen("tcp", ":"+strconv.Itoa(port+1))
	promLis, err := net.Listen("tcp", ":"+strconv.Itoa(port+2))
	if err != nil {
		cb <- err
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
	go func() {
		cb <- api.grpcServer.Serve(lis)
	}()
	go func() {
		cb <- api.grpcWebHTTPServer.Serve(webLis)
	}()
	go func() {
		cb <- api.prometheusServer.Serve(promLis)
	}()
}
