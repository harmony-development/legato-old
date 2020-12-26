package v1

import (
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	lru "github.com/hashicorp/golang-lru"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB     db.IHarmonyDB
	Logger logger.ILogger
	Config *config.Config
}

// V1 contains the gRPC handler for v1
type V1 struct {
	Dependencies
	instantViewData
	linkData
}

// Initialize initializes the V1 service
func (v1 *V1) Initialize() {
	v1.linkData.linkLRU, _ = lru.NewARC(v1.Config.Server.Policies.MaximumCacheSizes.LinkEmbeds)
	v1.instantViewData.instantLRU, _ = lru.NewARC(v1.Config.Server.Policies.MaximumCacheSizes.InstantView)
}
