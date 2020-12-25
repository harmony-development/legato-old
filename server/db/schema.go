package db

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/harmony-development/legato/server/data"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

// Migrate applies the DB layout to the connected DB
func (db *HarmonyDB) Migrate() error {
	data := data.AssertByteArray(ioutil.ReadAll(data.AssertFile(data.FS(false).Open("/sql/"))))

	_, err := db.Exec(strings.ReplaceAll(string(data), "--migration-only", ""))
	return tracerr.Wrap(err)
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
