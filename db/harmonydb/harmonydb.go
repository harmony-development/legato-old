package harmonydb

import (
	"context"
	"fmt"
	"time"

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
	l.Info("Waiting for postgres to appear...")
	waitPostgres, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var pool *pgxpool.Pool
	for {
		if conn, err := pgxpool.Connect(waitPostgres, connString); err == nil {
			pool = conn
			break
		} else {
			l.WithError(err).Warn("Failed to connect to postgres")
		}
		select {
		case <-time.After(1 * time.Second):
			continue
		case <-waitPostgres.Done():
			return nil, waitPostgres.Err()
		}
	}
	q := gen.New(pool)

	l.Info("Waiting for redis to appear...")
	waitRedis, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Redis.Hosts,
		Password: cfg.Redis.Password,
	})
	for {
		if _, err := rdb.Ping(waitRedis).Result(); err == nil {
			break
		} else {
			l.WithError(err).Warn("Failed to connect to redis")
		}
		select {
		case <-time.After(1 * time.Second):
			continue
		case <-waitRedis.Done():
			return nil, waitRedis.Err()
		}
	}

	return &HarmonyDB{
		Context: ctx,
		queries: q,
		rdb:     rdb,
	}, nil
}
