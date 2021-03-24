package test

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/harmony-development/legato/server/config"
	"github.com/labstack/echo/v4"
)

func DefaultConf() *config.Config {
	return &config.Config{
		Server: config.ServerConf{
			Host:           "0.0.0.0",
			Port:           2289,
			PrivateKeyPath: "harmony-key.pem",
			PublicKeyPath:  "harmony-key.pub",
			StorageBackend: "PureFlatfile",
			UseCORS:        true,
			UseTLS:         true,
			TLSCert:        "./filestore/localhost.pem",
			TLSKey:         "./filestore/localhost-key.pem",
			Policies: config.ServerPolicies{
				Avatar: config.AvatarPolicy{
					Width:   256,
					Height:  256,
					Quality: 50,
					Crop:    true,
				},
				Username: config.UsernamePolicy{
					MinLength: 2,
					MaxLength: 20,
				},
				Password: config.PasswordPolicy{
					MinLength:  5,
					MaxLength:  256,
					MinLower:   1,
					MinUpper:   1,
					MinNumbers: 1,
				},
				Attachments: config.AttachmentPolicy{
					MaximumCount: 10,
				},
				Debug: config.DebugPolicy{
					LogErrors:                  true,
					LogRequests:                true,
					RespondWithErrors:          true,
					ResponseErrorsIncludeTrace: true,
					VerboseStreamHandling:      true,
				},
				Sessions: config.SessionPolicy{
					Duration: time.Hour * 48,
				},
				MaximumCacheSizes: config.CachePolicy{
					Owner:       5096,
					Sessions:    5096,
					LinkEmbeds:  65536,
					InstantView: 65536,
				},
				APIs: config.APIPolicy{
					Messages: config.MessagesPolicy{
						MaximumGetAmount: 50,
					},
				},
				Federation: config.FederationPolicy{
					NonceLength:                       32,
					GuildLeaveNotificationQueueLength: 64,
				},
			},
		},
		Database: config.DBConf{
			Host:     "127.0.0.1",
			Username: "amadeus",
			Password: "password",
			Port:     5432,
			Name:     "harmony",
			Backend:  "postgres",
			Filename: "data.db",
		},
		Flatfile: config.FlatfileConf{
			MediaPath: "flatfile",
		},
		Sentry: config.SentryConf{
			AttachStacktraces: true,
		},
	}
}

func DummyContext(e *echo.Echo) echo.Context {
	return e.NewContext(httptest.NewRequest(http.MethodGet, "https://127.0.0.1", nil), httptest.NewRecorder())
}
