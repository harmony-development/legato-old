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
		Config:      defaultConf(),
	})
}

func defaultConf() *config.Config {
	return &config.Config{
		Server: config.ServerConf{
			Policies: config.ServerPolicies{
				EnablePasswordResetForm: false,
			},
		},
	}
}

func dummyContext(e *echo.Echo) echo.Context {
	return e.NewContext(httptest.NewRequest(http.MethodGet, "https://127.0.0.1", nil), httptest.NewRecorder())
}

func beginAuth(c echo.Context, a *assert.Assertions, api *v1.V1) (string, error) {
	resp, err := api.BeginAuth(c, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.AuthId, nil
}

func initialChoice(c echo.Context, a *assert.Assertions, api *v1.V1, authID string) (*authv1.AuthStep, error) {
	return api.NextStep(c, &authv1.NextStepRequest{
		AuthId: authID,
	})
}

func login(c echo.Context, a *assert.Assertions, api *v1.V1, email string, password string) (*authv1.AuthStep, error) {
	authID, err := beginAuth(c, a, api)
	a.Nil(err)
	a.NotEmpty(authID)
	initialStep, err := initialChoice(c, a, api, authID)
	a.Nil(err)
	a.False(initialStep.CanGoBack)
	a.IsType(&authv1.AuthStep_Choice_{}, initialStep.Step)
	a.Len(initialStep.GetChoice().Options, 3)
	a.Equal("initial-choice", initialStep.GetChoice().Title)
	loginPage, err := api.NextStep(c, &authv1.NextStepRequest{
		AuthId: authID,
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
		AuthId: authID,
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

func TestInitialChoice(t *testing.T) {
	a := assert.New(t)
	api := newAPI()
	ctx := dummyContext(echo.New())
	authID, err := beginAuth(ctx, a, api)
	a.Nil(err)
	a.NotEmpty(authID)
	step, err := api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
	})
	a.Nil(err)
	a.False(step.CanGoBack)
	a.IsType(&authv1.AuthStep_Choice_{}, step.Step)
	a.Equal("initial-choice", step.GetChoice().Title)
	a.ElementsMatch(step.GetChoice().Options, []string{"login", "register", "other-options"})
}

func TestStepBack(t *testing.T) {
	a := assert.New(t)
	api := newAPI()
	ctx := dummyContext(echo.New())
	authID, _ := beginAuth(ctx, a, api)
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
	a.Nil(err)
	step, err := api.StepBack(ctx, &authv1.StepBackRequest{
		AuthId: authID,
	})
	a.Nil(err)
	a.IsType(&authv1.AuthStep_Choice_{}, step.Step)
	a.Equal("initial-choice", step.GetChoice().Title)
	a.Equal(step.CanGoBack, false)
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
				a.NotNil(err)
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
