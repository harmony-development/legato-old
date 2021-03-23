package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	v1 "github.com/harmony-development/legato/server/api/authsvc/v1"
	authstate "github.com/harmony-development/legato/server/api/authsvc/v1/pubsub_backends/integrated"
	"github.com/harmony-development/legato/server/config"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newAPI() *v1.V1 {
	return v1.New(v1.Dependencies{
		DB:          MockDB{},
		Logger:      MockLogger{},
		Sonyflake:   sonyflake.NewSonyflake(sonyflake.Settings{}),
		AuthManager: MockAuthManager{},
		AuthState:   authstate.New(MockLogger{}),
		Config: &config.Config{
			Server: config.ServerConf{
				Policies: config.ServerPolicies{
					EnablePasswordResetForm: false,
				},
			},
		},
	})
}

func dummyContext(e *echo.Echo) echo.Context {
	return e.NewContext(httptest.NewRequest(http.MethodGet, "https://127.0.0.1", nil), httptest.NewRecorder())
}

func loginTester(t *testing.T) {

}

func TestLogin(t *testing.T) {
	var testMatrix = []struct {
		email      string
		password   string
		shouldFail bool
	}{
		{"amadeus@viktorchondria.jp", "@&GyubhjA^GYUH", true},
	}

	for _, test := range testMatrix {
		t.Run(test.email, func(t *testing.T) {
			a := assert.New(t)
			api := newAPI()
			e := echo.New()
			resp, err := api.BeginAuth(dummyContext(e), &emptypb.Empty{})
			assert.NoError(t, err)
			assert.Greater(t, len(resp.AuthId), 8)
			initialPage, err := api.NextStep(dummyContext(e), &authv1.NextStepRequest{
				AuthId: resp.AuthId,
			})
			a.NoError(err)
			a.False(initialPage.CanGoBack)
			a.IsType(&authv1.AuthStep_Choice_{}, initialPage.Step)
			a.Len(initialPage.GetChoice().Options, 3)
			a.Equal("initial-choice", initialPage.GetChoice().Title)
			loginPage, err := api.NextStep(dummyContext(e), &authv1.NextStepRequest{
				AuthId: resp.AuthId,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "login",
					},
				},
			})
			a.NoError(err)
			a.True(loginPage.CanGoBack)
			a.IsType(&authv1.AuthStep_Form_{}, loginPage.Step)
			a.Equal("login", loginPage.GetForm().Title)
			a.Len(loginPage.GetForm().Fields, 2)
			step, err := api.NextStep(dummyContext(e), &authv1.NextStepRequest{
				AuthId: resp.AuthId,
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
			if test.shouldFail {
				a.Error(err)
			} else {
				a.NoError(err)
				a.IsType(&authv1.AuthStep_Session{}, step.Step)
				a.NotZero(step.GetSession().UserId)
				a.Greater(len(step.GetSession().SessionToken), 8)
			}
		})
	}
}
