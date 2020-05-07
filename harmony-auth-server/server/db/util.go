package db

import (
	"database/sql"
	"fmt"
)

// Transact does a query in a transaction so you can have transaction capabilities later
func (db DB) Transact(query string, args ...interface{}) (sql.Result, *sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, err
	}
	res, err := tx.Exec(query, args)
	if err != nil {
		return nil, nil, err
	}
	return res, tx, nil
}

// ContainsRow checks if a row exists
func (db DB) ContainsRow(query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}