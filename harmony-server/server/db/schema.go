package db

import (
	"io/ioutil"
	"strings"
)

// Migrate applies the DB layout to the connected DB
func (db *HarmonyDB) Migrate() error {
	data, err := ioutil.ReadFile("sql/schemas/models.sql")
	if err != nil {
		return err
	}
	models := strings.Split(string(data), ";")
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, q := range models {
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return nil
}
