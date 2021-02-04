package types

import (
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
)

type IBackend interface {
	New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (IHarmonyDB, error)
	Name() string
}
