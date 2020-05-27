package hm

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Permission int

const (
	NoPermission = iota
	ModifyInvites
	ModifyChannels
	ModifyGuild
	Owner
)

func (m *Middlewares) ForGuildPermission(perm Permission) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			owner, err := m.DB.GetOwner(*ctx.Location.GuildID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if owner == ctx.UserID {
				return handler(ctx)
			}
			switch perm {
			case NoPermission:
				return handler(ctx)
			}
			return echo.NewHTTPError(http.StatusForbidden)
		}
	}
}

func (m *Middlewares) ForChannelPermission(perm Permission) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			owner, err := m.DB.GetOwner(*ctx.Location.GuildID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if owner == ctx.UserID {
				return handler(ctx)
			}
			switch perm {
			case NoPermission:
				return handler(ctx)
			}
			return echo.NewHTTPError(http.StatusForbidden)
		}
	}
}
