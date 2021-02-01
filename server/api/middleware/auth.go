package middleware

import (
	"errors"

	"github.com/harmony-development/legato/server/http/responses"
	"github.com/labstack/echo/v4"
)

func (m *Middlewares) AuthHandler(c echo.Context) (uint64, error) {
	session := c.Request().Header.Get("Authorization")

	userID, err := m.DB.SessionToUserID(session)
	if err != nil {
		println("bad session")
		return 0, errors.New(responses.InvalidSession)
	}
	return userID, nil
}
