package cassandra

import "embed"

//go:embed migrations/*
var migrations embed.FS // nolint: gochecknoglobals

func (db *DB) Migrate() error {
	return nil
}
