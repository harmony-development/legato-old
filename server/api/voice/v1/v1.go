package v1

import (
	"time"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/labstack/echo/v4"
)

type Dependencies struct {
	DB types.IHarmonyDB
}

type V1 struct {
	Dependencies
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.VoiceService/Connect")
}

func (v1 *V1) Connect(ctx echo.Context, in chan *voicev1.ClientSignal, out chan *voicev1.Signal) {
}
