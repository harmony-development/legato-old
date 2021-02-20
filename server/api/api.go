package api

import (
	"unsafe"

	"github.com/alecthomas/repr"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/harmony-development/hrpc/server"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	mediaproxyv1 "github.com/harmony-development/legato/gen/mediaproxy/v1"
	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/authsvc"
	"github.com/harmony-development/legato/server/api/chat"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/api/mediaproxy"
	hm "github.com/harmony-development/legato/server/api/middleware"
	voicev1impl "github.com/harmony-development/legato/server/api/voice/v1"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/http"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/responses"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sony/sonyflake"
	"google.golang.org/protobuf/proto"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
)

type Dependencies struct {
	Logger         logger.ILogger
	DB             types.IHarmonyDB
	Sonyflake      *sonyflake.Sonyflake
	AuthManager    *auth.Manager
	Config         *config.Config
	Permissions    *permissions.Manager
	StorageBackend backend.AttachmentBackend
}

// API contains the component of the server responsible for APIs
type API struct {
	*echo.Echo
	Dependencies
}

// New creates a new API instance
func New(deps Dependencies) *API {
	api := &API{
		Echo:         echo.New(),
		Dependencies: deps,
	}

	m := hm.New(hm.Dependencies{
		Logger: deps.Logger,
		DB:     deps.DB,
		Perms:  api.Permissions,
	})

	api.Echo.HTTPErrorHandler = func(e error, c echo.Context) {
		if deps.Config.Sentry.Enabled {
			sentry.CaptureException(e)
		}
		if deps.Config.Server.Policies.Debug.LogErrors && e != nil {
			c.Logger().Error(repr.String(e))
		}
		switch v := e.(type) {
		case *responses.Error:
			data, err := proto.Marshal((*harmonytypesv1.Error)(unsafe.Pointer(v)))
			if err != nil {
				c.Logger().Error(e)
				return
			}
			if err := c.Blob(400, "application/octet-stream", data); err != nil {
				c.Logger().Error(err)
			}
		default:
			err := &harmonytypesv1.Error{
				Identifier: responses.InternalServerError,
			}
			if api.Config.Server.Policies.Debug.RespondWithErrors {
				err.HumanMessage = v.Error()
			}
			data, i := proto.Marshal(err)
			if i != nil {
				c.Logger().Error(i)
				return
			}
			if err := c.Blob(500, "application/octet-stream", data); err != nil {
				c.Logger().Error(err)
			}
		}
	}
	api.Echo.Use(middleware.Logger())
	api.Echo.Use(middleware.AddTrailingSlash())
	api.Echo.Use(middleware.Recover())

	if deps.Config.Sentry.Enabled {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: deps.Config.Sentry.DSN,
		}); err != nil {
			panic(err)
		}
		api.Echo.Use(sentryecho.New(sentryecho.Options{
			Repanic: true,
		}))
	}
	if deps.Config.Server.UseCORS {
		api.Echo.Use(middleware.CORS())
	}

	http.New(api.Echo, http.Dependencies{
		DB:             deps.DB,
		Logger:         deps.Logger,
		Config:         deps.Config,
		StorageBackend: api.StorageBackend,
	})

	authService := authv1.NewAuthServiceHandler(authsvc.New(&authsvc.Dependencies{
		DB:          deps.DB,
		Logger:      deps.Logger,
		Sonyflake:   deps.Sonyflake,
		Config:      deps.Config,
		AuthManager: deps.AuthManager,
	}).V1)
	chatService := chatv1.NewChatServiceHandler(chat.New(&chat.Dependencies{
		DB:             deps.DB,
		Logger:         deps.Logger,
		Sonyflake:      deps.Sonyflake,
		Perms:          deps.Permissions,
		Config:         deps.Config,
		StorageBackend: deps.StorageBackend,
		Middlewares:    m,
	}).V1)
	mediaProxyService := mediaproxyv1.NewMediaProxyServiceHandler(mediaproxy.New(&mediaproxy.Dependencies{
		DB:     deps.DB,
		Logger: deps.Logger,
		Config: deps.Config,
	}).V1)
	voiceService := voicev1.NewVoiceServiceHandler(&voicev1impl.V1{
		Dependencies: voicev1impl.Dependencies{
			DB: deps.DB,
		},
	})

	hrpcServer := server.NewHRPCServer(api.Echo, authService, chatService, mediaProxyService, voiceService)

	hrpcServer.SetUnaryPre(server.ChainHandlerTransformers(
		m.UnaryRecoveryFunc,
		m.HarmonyContextInterceptor,
		m.RateLimitInterceptor,
		m.Validate,
		m.MethodMetadataInterceptor,
	))

	return api
}
