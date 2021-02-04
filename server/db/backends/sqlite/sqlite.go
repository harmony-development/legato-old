package sqlite

import (
	"context"
	"fmt"
	"runtime"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/backends/sqlite/ent"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"
	"github.com/ztrue/tracerr"

	// backend
	_ "github.com/mattn/go-sqlite3"
	"github.com/sony/sonyflake"
)

type sqliteBackend struct {
}

func (p sqliteBackend) New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (types.IHarmonyDB, error) {
	return New(cfg, logger, idgen)
}
func (p sqliteBackend) Name() string {
	return "sqlite"
}

func init() {
	db.RegisterBackend(sqliteBackend{})
}

type database struct {
	*ent.Client
	types.DummyDB
}

var ctx = context.Background()

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

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (types.IHarmonyDB, error) {
	db := &database{}
	err := error(nil)

	db.Client, err = ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", cfg.Database.Filename))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	err = db.Schema.Create(context.Background())
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	//go db.SessionExpireRoutine()

	return db, nil
}
