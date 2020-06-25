package hm

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Middlewares) WithGuild(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		guildIDString := c.Param("guild_id")
		parsed, err := strconv.ParseUint(guildIDString, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "guild ID in invalid form")
		}
		valid, err := m.DB.HasGuildWithID(parsed)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if !valid {
			return echo.NewHTTPError(http.StatusBadRequest, "guild ID doesn't exist")
		}
		ctx.Location.GuildID = &parsed
		return handler(ctx)
	}
}

func (m *Middlewares) WithChannel(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		channelIDString := c.Param("channel_id")
		parsed, err := strconv.ParseUint(channelIDString, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "channel ID in invalid form")
		}
		valid, err := m.DB.HasChannelWithID(*ctx.Location.GuildID, parsed)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if !valid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		ctx.Location.ChannelID = &parsed
		return handler(ctx)
	}
}

func (m *Middlewares) WithMessage(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		messageIDString := c.Param("message_id")
		parsed, err := strconv.ParseUint(messageIDString, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "message ID in invalid form")
		}
		msg, err := m.DB.GetMessage(uint64(parsed))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		ctx.Location.Message = &msg
		return handler(ctx)
	}
}

func (m *Middlewares) WithUser(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		userIDString := c.Param("user_id")
		parsed, err := strconv.ParseUint(userIDString, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "user ID in invalid form")
		}
		user, err := m.DB.GetUserByID(parsed)
		if err != nil {
			m.Logger.Exception(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		ctx.Location.User = &user
		return handler(ctx)
	}
}
