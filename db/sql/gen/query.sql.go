// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package gen

import (
	"context"
)

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM AuthSessions WHERE SessionID = $1
`

func (q *Queries) DeleteSession(ctx context.Context, sessionid string) error {
	_, err := q.db.Exec(ctx, deleteSession, sessionid)
	return err
}

const getSession = `-- name: GetSession :one
SELECT UserID FROM AuthSessions WHERE SessionID = $1
`

func (q *Queries) GetSession(ctx context.Context, sessionid string) (int64, error) {
	row := q.db.QueryRow(ctx, getSession, sessionid)
	var userid int64
	err := row.Scan(&userid)
	return userid, err
}

const setSession = `-- name: SetSession :exec
INSERT INTO AuthSessions(UserID, SessionID) VALUES($1, $2)
`

type SetSessionParams struct {
	Userid    int64
	Sessionid string
}

func (q *Queries) SetSession(ctx context.Context, arg SetSessionParams) error {
	_, err := q.db.Exec(ctx, setSession, arg.Userid, arg.Sessionid)
	return err
}
