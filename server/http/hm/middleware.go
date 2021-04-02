package hm

import (
	"net/http"
	"sync"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// LocationContext stores information about the location of items
type LocationContext struct {
	GuildID   *uint64
	ChannelID *uint64
	Message   *harmonytypesv1.Message
	User      *types.UserData
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
	DB         types.IHarmonyDB
	Logger     logger.ILogger
	RateLimits map[string]map[string]*visitor
	RateLock   sync.RWMutex
	Config     *config.Config
}

func (hc *HarmonyContext) VerifyOwner(db types.IHarmonyDB, guildID, userID uint64) error {
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
func New(db types.IHarmonyDB, logger logger.ILogger, cfg *config.Config) *Middlewares {
	m := &Middlewares{
		DB:         db,
		Logger:     logger,
		Config:     cfg,
		RateLimits: make(map[string]map[string]*visitor),
	}
	go m.RateCleanup()
	return m
}
