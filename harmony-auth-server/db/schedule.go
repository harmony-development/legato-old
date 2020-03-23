package db

import (
	"github.com/kataras/golog"
	"time"
)

// ExpireSessions deletes sessions that have expired from the table.
func ExpireSessions() {
	for {
		if _, err := DB.Exec("DELETE FROM sessions WHERE expiration<$1", time.Now().Unix()); err != nil {
			golog.Warn("Error expiring old sessions ", err)
		}
		time.Sleep(20 * time.Minute)
	}
}