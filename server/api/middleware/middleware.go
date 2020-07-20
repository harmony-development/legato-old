package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"golang.org/x/time/rate"
)

type Permission uint64

const (
	NoPermission  = 0
	ModifyInvites = 1 << iota
	ModifyChannels
	ModifyGuild
	Owner
)

func (flag *Permission) Set(b Permission)     { *flag = b | *flag }
func (flag *Permission) Clear(b Permission)   { *flag = b &^ *flag }
func (flag *Permission) Toggle(b Permission)  { *flag = b ^ *flag }
func (flag Permission) Has(b Permission) bool { return b&flag != 0 }
func (flag Permission) HasAll(b ...Permission) bool {
	for _, perm := range b {
		if perm&flag == 0 {
			return false
		}
	}
	return true
}

func (flag Permission) HasAny(b ...Permission) bool {
	for _, perm := range b {
		if perm&flag != 0 {
			return true
		}
	}
	return false
}

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
	Auth       bool
	Permission Permission
	Location   Location
}

var rpcConfigs = map[string]RPCConfig{}

func RegisterRPCConfig(config RPCConfig, name ...string) {
	for _, name := range name {
		rpcConfigs[name] = config
	}
}

func GetRPCConfig(name string) RPCConfig {
	val, _ := rpcConfigs[name]
	return val
}

// HarmonyContext contains a custom context for passing data from middleware to handlers
type HarmonyContext struct {
	context.Context
	UserID  uint64
	Limiter *rate.Limiter
}

type Dependencies struct {
	Logger logger.ILogger
	DB     db.IHarmonyDB
}

type Middlewares struct {
	Dependencies
	RateLock *sync.RWMutex
	// RateLimits is a map of IP rate limits for each RPC route
	RateLimits map[string]map[string]visitor
}

func New(deps Dependencies) Middlewares {
	return Middlewares{
		Dependencies: deps,
		RateLock:     new(sync.RWMutex),
		RateLimits:   make(map[string]map[string]visitor),
	}
}
