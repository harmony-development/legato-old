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

// AddSampleData adds sample values to the DB for testing
func (db *HarmonyDB) AddSampleData() error {
	if err := db.Queries.AddGuild(
		"harmony-devs",
		"82ee9c8dc9e165205548b7c3833e7372",
		"Harmony Development",
		"",
	); err != nil {
		return err
	}

	if err := db.Queries.AddInvite("join-harmony-dev", "harmony-devs"); err != nil {
		return err
	}

	if err := db.Queries.AddUserToGuild("82ee9c8dc9e165205548b7c3833e7372", "harmony-devs"); err != nil {
		return err
	}

	if err := db.AddMemberToGuild("dadcd6bf8c0338cbfc9aa9c369ea93cc", "harmony-devs"); err != nil {
		return err
	}

	if err := db.AddChannelToGuild("yf8e3bhbp0Bk8UKG", "harmony-devs", "media"); err != nil {
		return err
	}

	return nil
}
