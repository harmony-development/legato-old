package protocol

import (
	v1 "harmony-server/server/http/core/v1"
	"harmony-server/util"

	"net/http"
	"unicode"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
)

type RegisterData struct {
	Email    string `validate:"required,email"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (h API) Register(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(RegisterData)
	if len(data.Username) < h.Deps.Config.Server.UsernamePolicy.MinLength ||
		len(data.Username) > h.Deps.Config.Server.UsernamePolicy.MaxLength {
		resp := responses.UsernameLength(
			h.Deps.Config.Server.UsernamePolicy.MinLength,
			h.Deps.Config.Server.UsernamePolicy.MaxLength,
		)
		return c.JSON(
			http.StatusNotAcceptable,
			&resp,
		)
	}
	if len(data.Password) < h.Deps.Config.Server.PasswordPolicy.MinLength ||
		len(data.Password) > h.Deps.Config.Server.PasswordPolicy.MaxLength {
		return ctx.JSON(
			http.StatusNotAcceptable,
			responses.PasswordLength(
				h.Deps.Config.Server.PasswordPolicy.MinLength,
				h.Deps.Config.Server.PasswordPolicy.MaxLength,
			),
		)
	}
	stats := getPasswordStats(data.Password)
	if stats.upper < h.Deps.Config.Server.PasswordPolicy.MinUpper ||
		stats.lower < h.Deps.Config.Server.PasswordPolicy.MinLower ||
		stats.numbers < h.Deps.Config.Server.PasswordPolicy.MinNumbers ||
		stats.symbols < h.Deps.Config.Server.PasswordPolicy.MinSymbols {
		return ctx.JSON(
			http.StatusNotAcceptable,
			responses.PasswordPolicy(
				h.Deps.Config.Server.PasswordPolicy.MinUpper,
				h.Deps.Config.Server.PasswordPolicy.MinLower,
				h.Deps.Config.Server.PasswordPolicy.MinNumbers,
				h.Deps.Config.Server.PasswordPolicy.MinSymbols,
			),
		)
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Deps.Logger.Exception(err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	exists, err := h.Deps.DB.EmailExists(data.Email)
	if err != nil {
		h.Deps.Logger.Exception(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	} else if exists {
		return echo.NewHTTPError(http.StatusConflict, responses.AlreadyRegistered)
	}
	userID, err := h.Deps.Sonyflake.NextID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	if err := h.Deps.DB.AddLocalUser(userID, data.Email, data.Username, hash); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	session := randstr.Hex(16)
	if err := h.Deps.DB.AddSession(userID, session); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.JSON(http.StatusOK, v1.RegisterResponse{Session: session, UserID: util.U64TS(userID)})
}

type passwordStats struct {
	upper   int
	lower   int
	numbers int
	symbols int
}

func getPasswordStats(password string) passwordStats {
	var stats passwordStats
	for _, c := range password {
		if unicode.IsUpper(c) {
			stats.upper++
		} else if unicode.IsLower(c) {
			stats.lower++
		} else if unicode.IsNumber(c) {
			stats.numbers++
		} else if unicode.IsSymbol(c) {
			stats.symbols++
		}
	}
	return stats
}
