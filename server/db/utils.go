package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
	"encoding/json"
)

func toSqlString(input string) sql.NullString {
	return sql.NullString{String: input, Valid: true}
}

func toSqlInt64(input uint64) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(input), Valid: true}
}

type executor struct {
	err error
}

func (e *executor) Execute(f func() error) {
	if e.err != nil {
		return
	}
	e.err = f()
}

var ctx = context.Background()

func mustDeserialize(d json.RawMessage, v interface{}) {
	err := json.Unmarshal(d, v)
	if err != nil {
		panic(err)
	}
}

func mustSerialize(v interface{}) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}
