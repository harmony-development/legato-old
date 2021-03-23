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
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newAPI() *v1.V1 {
	return v1.New(v1.Dependencies{
		DB: MockDB{
			users:       map[uint64]*User{},
			userByEmail: map[string]*User{},
		},
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

func login(c echo.Context, a *assert.Assertions, api *v1.V1, email string, password string) (*authv1.AuthStep, error) {
	resp, err := api.BeginAuth(c, &emptypb.Empty{})
	a.Nil(err)
	a.Greater(len(resp.AuthId), 8)
	initialPage, err := api.NextStep(c, &authv1.NextStepRequest{
		AuthId: resp.AuthId,
	})
	a.Nil(err)
	a.False(initialPage.CanGoBack)
	a.IsType(&authv1.AuthStep_Choice_{}, initialPage.Step)
	a.Len(initialPage.GetChoice().Options, 3)
	a.Equal("initial-choice", initialPage.GetChoice().Title)
	loginPage, err := api.NextStep(c, &authv1.NextStepRequest{
		AuthId: resp.AuthId,
		Step: &authv1.NextStepRequest_Choice_{
			Choice: &authv1.NextStepRequest_Choice{
				Choice: "login",
			},
		},
	})
	a.Nil(err)
	a.True(loginPage.CanGoBack)
	a.IsType(&authv1.AuthStep_Form_{}, loginPage.Step)
	a.Equal("login", loginPage.GetForm().Title)
	a.Len(loginPage.GetForm().Fields, 2)
	return api.NextStep(c, &authv1.NextStepRequest{
		AuthId: resp.AuthId,
		Step: &authv1.NextStepRequest_Form_{
			Form: &authv1.NextStepRequest_Form{
				Fields: []*authv1.NextStepRequest_FormFields{
					{
						Field: &authv1.NextStepRequest_FormFields_String_{
							String_: email,
						},
					},
					{
						Field: &authv1.NextStepRequest_FormFields_Bytes{
							Bytes: []byte(password),
						},
					},
				},
			},
		},
	})
}

func TestLogin(t *testing.T) {
	var testMatrix = []struct {
		email      string
		password   string
		shouldFail bool
	}{
		{"amadeus@viktorchondria.jp", "@&GyubhjA^GYUH", false},
	}

	for _, test := range testMatrix {
		t.Run(test.email, func(t *testing.T) {
			a := assert.New(t)
			api := newAPI()
			e := echo.New()
			hashed, err := bcrypt.GenerateFromPassword([]byte(test.password), 0)
			a.Nil(err)
			api.DB.AddLocalUser(12345, test.email, "amadeus", hashed)
			loginStep, err := login(dummyContext(e), a, api, test.email, test.password)
			if test.shouldFail {
				a.Error(err)
			} else {
				a.Nil(err)
				a.NotNil(loginStep)
				a.True(loginStep.CanGoBack)
				a.IsType(&authv1.AuthStep_Session{}, loginStep.Step)
				a.Equal(uint64(12345), loginStep.GetSession().UserId)
				a.Greater(len(loginStep.GetSession().SessionToken), 8)
			}
		})
	}
}
