package v1

import (
	"time"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/db"
)

type Dependencies struct {
	DB db.IHarmonyDB
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

func (v1 *V1) Connect(s voicev1.VoiceService_ConnectServer) error {
	return nil
}
