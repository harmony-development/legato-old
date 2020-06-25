package hm

import (
	"net/http"
	"sync"

	"harmony-server/server/db"
	"harmony-server/server/db/queries"
	"harmony-server/server/logger"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// LocationContext stores information about the location of items
type LocationContext struct {
	GuildID   *uint64
	ChannelID *uint64
	Message   *queries.Message
	User      *queries.GetUserRow
}

// A HarmonyContext adds rate limiting and a user ID to an echo.Context
type HarmonyContext struct {
	echo.Context
	Limiter  *rate.Limiter
	UserID   uint64
	Data     interface{}
	Location LocationContext
}

// Middlewares contains middlewares for Harmony
type Middlewares struct {
	DB         db.IHarmonyDB
	Logger     *logger.Logger
	RateLimits map[string]map[string]*visitor
	RateLock   sync.RWMutex
}

func (hc *HarmonyContext) VerifyOwner(db db.IHarmonyDB, guildID, userID uint64) error {
	owner, err := db.GetOwner(guildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to verify ownership, please try again later")
	}
	if owner != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "insufficient permissions")
	}
	return nil
}

// New instantiates the middlewares for Harmony
func New(db db.IHarmonyDB, logger *logger.Logger) *Middlewares {
	m := &Middlewares{
		DB:         db,
		Logger:     logger,
		RateLimits: make(map[string]map[string]*visitor),
	}
	go m.RateCleanup()
	return m
}
