package db

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type UserStatus int16

const (
	UserStatusOnline = iota
	UserStatusStreaming
	UserStatusIdle
	UserStatusOffline
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

func (db *HarmonyDB) SessionExpireRoutine() {
	for {
		time.Sleep(15 * time.Minute)
		err := db.ExpireSessions()
		if err != nil {
			logrus.Warn(err)
			sentry.CaptureException(err)
		}
	}
}
