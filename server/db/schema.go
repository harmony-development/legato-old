package db

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

// Migrate applies the DB layout to the connected DB
func (db *HarmonyDB) Migrate() error {
	data, err := ioutil.ReadFile("sql/schemas/models.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(strings.ReplaceAll(string(data), "--migration-only", ""))
	return err
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
