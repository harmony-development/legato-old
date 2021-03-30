package ent_shared

import (
	"context"
	"runtime"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"
	"github.com/ztrue/tracerr"

	// backend
	_ "github.com/lib/pq"
	"github.com/sony/sonyflake"
)

type database struct {
	*entgen.Client
	types.DummyDB
}

var ctx = context.Background()

// New creates a new DB connection
func New(c *entgen.Client, cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (types.IHarmonyDB, error) {
	db := &database{}
	db.Client = c
	if err := db.Schema.Create(context.Background()); err != nil {
		return nil, tracerr.Wrap(err)
	}

	//go db.SessionExpireRoutine()

	return db, nil
}

func (db *database) TxX() *entgen.Tx {
	tx, err := db.Tx(ctx)
	if err != nil {
		panic(err)
	}
	return tx
}

func doRecovery(err *error) {
	r := recover()
	if r == nil {
		return
	}
	ierr, ok := r.(error)
	if !ok {
		panic(r)
	}

	frames := make([]tracerr.Frame, 0, 40)
	skip := 0
	for {
		pc, path, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		frame := tracerr.Frame{
			Func: fn.Name(),
			Line: line,
			Path: path,
		}
		frames = append(frames, frame)
		skip++
	}

	*err = tracerr.CustomError(ierr, frames)
}
