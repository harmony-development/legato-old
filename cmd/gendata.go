package cmd

import (
	"time"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
	"golang.org/x/crypto/bcrypt"
)

const (
	Email    = `ilopona@toki.pona`
	Password = `10kekeAke`
	Username = `yooter`
)

func GenData() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Unable to load config", err)
	}

	database, err := db.New(cfg, logger.New(cfg))
	if err != nil {
		logrus.Fatal("Unable to connect to database", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal("Unable to hash password", err)
	}

	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Unix(cfg.Server.SnowflakeStart, 0),
	})

	userID, err := sonyflake.NextID()
	if err != nil {
		logrus.Fatal("Unable to get snowflake", err)
	}

	if err := database.AddLocalUser(userID, Email, Username, hash); err != nil {
		logrus.Fatal("Unable to add user", err)
	}
}
