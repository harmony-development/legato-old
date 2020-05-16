package db

import (
	"database/sql"
	"fmt"
)

// ContainsRow returns whether the database has a given row
func (db *HarmonyDB) ContainsRow(query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
