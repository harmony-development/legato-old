package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	v1 "github.com/harmony-development/legato/server/api/authsvc/v1"
	authstate "github.com/harmony-development/legato/server/api/authsvc/v1/pubsub_backends/integrated"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/responses"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newAPI() *v1.V1 {
	return v1.New(v1.Dependencies{
		DB: MockDB{
			users:         map[uint64]*User{},
			userByEmail:   map[string]*User{},
			userBySession: map[string]uint64{},
		},
		Logger:      MockLogger{},
		Sonyflake:   sonyflake.NewSonyflake(sonyflake.Settings{}),
		AuthManager: MockAuthManager{},
		AuthState:   authstate.New(MockLogger{}),
		Config:      defaultConf(),
	})
}

func defaultConf() *config.Config {
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

func dummyContext(e *echo.Echo) echo.Context {
	return e.NewContext(httptest.NewRequest(http.MethodGet, "https://127.0.0.1", nil), httptest.NewRecorder())
}

func beginAuth(c echo.Context, api *v1.V1) (string, error) {
	resp, err := api.BeginAuth(c, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.AuthId, nil
}

func initialChoice(c echo.Context, api *v1.V1, authID string) (*authv1.AuthStep, error) {
	return api.NextStep(c, &authv1.NextStepRequest{
		AuthId: authID,
	})
}

func TestInitialChoice(t *testing.T) {
	a := require.New(t)
	api := newAPI()
	ctx := dummyContext(echo.New())
	authID, err := beginAuth(ctx, api)
	a.NoError(err)
	a.NotEmpty(authID)
	step, err := api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
	})
	a.NoError(err)
	a.False(step.CanGoBack)
	a.IsType(&authv1.AuthStep_Choice_{}, step.Step)
	a.Equal("initial-choice", step.GetChoice().Title)
	a.ElementsMatch(step.GetChoice().Options, []string{"login", "register", "other-options"})
}

func TestStepBack(t *testing.T) {
	a := require.New(t)
	api := newAPI()
	ctx := dummyContext(echo.New())
	authID, _ := beginAuth(ctx, api)
	api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
	})
	_, err := api.StepBack(ctx, &authv1.StepBackRequest{
		AuthId: authID,
	})
	a.NotNil(err)
	_, err = api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
		Step: &authv1.NextStepRequest_Choice_{
			Choice: &authv1.NextStepRequest_Choice{
				Choice: "login",
			},
		},
	})
	a.NoError(err)
	step, err := api.StepBack(ctx, &authv1.StepBackRequest{
		AuthId: authID,
	})
	a.NoError(err)
	a.IsType(&authv1.AuthStep_Choice_{}, step.Step)
	a.Equal("initial-choice", step.GetChoice().Title)
	a.Equal(step.CanGoBack, false)
}

func TestLogin(t *testing.T) {
	var testMatrix = []struct {
		email       string
		password    string
		expectError string
	}{
		{"amadeus@home.cern", "@&GyubhjA^GYUH1", ""},
		{"amadeus@home.cern", "", ""},
	}

	for _, test := range testMatrix {
		t.Run(test.email, func(t *testing.T) {
			a := require.New(t)
			api := newAPI()
			ctx := dummyContext(echo.New())
			hashed, err := bcrypt.GenerateFromPassword([]byte(test.password), 0)
			a.NoError(err)
			api.DB.AddLocalUser(12345, test.email, "amadeus", hashed)
			authID, _ := beginAuth(ctx, api)
			initialChoice(ctx, api, authID)
			api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "login",
					},
				},
			})
			sessionStep, err := api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Form_{
					Form: &authv1.NextStepRequest_Form{
						Fields: []*authv1.NextStepRequest_FormFields{
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: test.email,
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_Bytes{
									Bytes: []byte(test.password),
								},
							},
						},
					},
				},
			})
			if test.expectError != "" {
				a.Error(err)
				a.Equal(test.expectError, err.Error())
			} else {
				a.NoError(err)
				a.NotNil(sessionStep)
				a.True(sessionStep.CanGoBack)
				a.IsType(&authv1.AuthStep_Session{}, sessionStep.Step)
				a.Equal(uint64(12345), sessionStep.GetSession().UserId)
				a.Greater(len(sessionStep.GetSession().SessionToken), 8)
				id, err := api.DB.SessionToUserID(sessionStep.GetSession().SessionToken)
				a.NoError(err)
				a.Equal(uint64(12345), id)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	var testMatrix = []struct {
		email       string
		username    string
		password    string
		expectError string
		name        string
	}{
		{
			name:        "Normal Registration",
			email:       "amadeus@home.cern",
			username:    "Amadeus",
			password:    "@&GyubhjA^GYUH1",
			expectError: "",
		},
		{
			name:        "Registration with a bad email",
			email:       "this is not an email",
			username:    "Amadeus",
			password:    "@&GyubhjA^GYUH1",
			expectError: responses.BadEmail,
		},
		{
			name:        "Registration with a bad password",
			email:       "amadeus@home.cern",
			username:    "Amadeus",
			password:    "this is not a password that will work since its just lowercase",
			expectError: responses.BadPassword,
		},
		{
			name:        "Registration with a short username",
			email:       "amadeus@home.cern",
			username:    "a",
			password:    "@&GyubhjA^GYUH1",
			expectError: responses.BadUsername,
		},
		{
			name:        "Registration with a long username",
			email:       "amadeus@home.cern",
			username:    "Hello my name is very long. Nobody should have a name this long. Long names are annoying. Just a few more words to make sure...",
			password:    "@&GyubhjA^GYUH1",
			expectError: responses.BadUsername,
		},
	}

	for _, test := range testMatrix {
		t.Run(test.name, func(t *testing.T) {
			a := require.New(t)
			api := newAPI()
			ctx := dummyContext(echo.New())
			authID, _ := beginAuth(ctx, api)
			initialChoice(ctx, api, authID)
			api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "register",
					},
				},
			})
			sessionStep, err := api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Form_{
					Form: &authv1.NextStepRequest_Form{
						Fields: []*authv1.NextStepRequest_FormFields{
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: test.email,
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: test.username,
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_Bytes{
									Bytes: []byte(test.password),
								},
							},
						},
					},
				},
			})
			if test.expectError != "" {
				a.EqualError(err, test.expectError)
			} else {
				a.NoError(err)
				a.NotNil(sessionStep)
				a.True(sessionStep.CanGoBack)
				a.IsType(&authv1.AuthStep_Session{}, sessionStep.Step)
				user, err := api.DB.GetUserByID(sessionStep.GetSession().UserId)
				a.NoError(err)
				a.Equal(test.username, user.Username)
				userID, err := api.DB.SessionToUserID(sessionStep.GetSession().SessionToken)
				a.Equal(sessionStep.GetSession().UserId, userID)
				a.Greater(len(sessionStep.GetSession().SessionToken), 8)
			}
		})
	}
}

func BenchmarkBeginAuth(b *testing.B) {
	ctx := dummyContext(echo.New())
	api := newAPI()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			beginAuth(ctx, api)
		}
	})
}

func BenchmarkLogin(b *testing.B) {
	ctx := dummyContext(echo.New())
	api := newAPI()
	hashed, err := bcrypt.GenerateFromPassword([]byte("@&GyubhjA^GYUH1"), 0)
	if err != nil {
		panic(err)
	}
	api.DB.AddLocalUser(12345, "amadeus@home.cern", "amadeus", hashed)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			authID, err := beginAuth(ctx, api)
			if err != nil {
				panic(err)
			}
			_, err = api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "login",
					},
				},
			})
			if err != nil {
				panic(err)
			}
			_, err = api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Form_{
					Form: &authv1.NextStepRequest_Form{
						Fields: []*authv1.NextStepRequest_FormFields{
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: "amadeus@home.cern",
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_Bytes{
									Bytes: []byte("@&GyubhjA^GYUH1"),
								},
							},
						},
					},
				},
			})
			if err != nil {
				panic(err)
			}
		}
	})
}
