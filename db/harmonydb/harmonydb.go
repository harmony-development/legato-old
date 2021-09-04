package harmonydb

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/go-redis/redis/v8"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/sql/gen"
	"github.com/jackc/pgx/v4/pgxpool"
)

// The default Harmony DB implementation. This uses sqlc.
type HarmonyDB struct {
	// embed a context for easy use in queries
	context.Context
	queries *gen.Queries
	rdb     *redis.ClusterClient
}

func New(l log.Interface, cfg *config.Config) (*HarmonyDB, error) {
	ctx := context.TODO()

	username, password, host, port, db :=
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DB

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		username,
		password,
		host,
		port,
		db,
	)

	conn, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	q := gen.New(conn)

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Redis.Hosts,
		Password: cfg.Redis.Password,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &HarmonyDB{
		Context: ctx,
		queries: q,
		rdb:     rdb,
	}, nil
}
