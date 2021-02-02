package middleware

import (
	"sync"
	"time"

	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
)

type Location uint64

const (
	NoLocation    = 0
	GuildLocation = 1 << iota
	ChannelLocation
	MessageLocation
	JoinedLocation
	AuthorLocation
)

func (flag *Location) Set(b Location)     { *flag = b | *flag }
func (flag *Location) Clear(b Location)   { *flag = b &^ *flag }
func (flag *Location) Toggle(b Location)  { *flag = b ^ *flag }
func (flag Location) Has(b Location) bool { return b&flag != 0 }

type RateLimit struct {
	Duration time.Duration
	Burst    int
}

type RPCConfig struct {
	RateLimit  RateLimit
	WantsRoles bool
	Location   Location
}

var rpcConfigs = map[string]RPCConfig{}

func RegisterRPCConfig(config RPCConfig, name ...string) {
	for _, name := range name {
		rpcConfigs[name] = config
	}
}

func GetRPCConfig(name string) RPCConfig {
	return rpcConfigs[name]
}

type Dependencies struct {
	Logger logger.ILogger
	DB     db.IHarmonyDB
	Perms  *permissions.Manager
}

type Middlewares struct {
	Dependencies
	RateLock *sync.RWMutex
	// RateLimits is a map of IP rate limits for each RPC route
	RateLimits map[string]map[string]visitor
}

func New(deps Dependencies) *Middlewares {
	return &Middlewares{
		Dependencies: deps,
		RateLock:     new(sync.RWMutex),
		RateLimits:   make(map[string]map[string]visitor),
	}
}
