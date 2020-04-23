package db

import "database/sql"

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
