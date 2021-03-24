package test

import (
	"testing"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	v1 "github.com/harmony-development/legato/server/api/chat/v1"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/api/chat/v1/pubsub_backends/inprocess"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/require"
)

func setupChatAPI() *v1.V1 {
	conf := defaultConf()
	logger := MockLogger{}
	db := MockDB{
		userBySession: map[string]uint64{},
		userByEmail:   map[string]*User{},
		users:         map[uint64]*User{},
	}
	perms := permissions.NewManager(db, logger)
	md := middleware.New(middleware.Dependencies{
		Logger: logger,
		DB:     db,
		Perms:  perms,
	})
	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{})
	attachments := &MockAttachments{
		files: map[string]*Attachment{},
	}
	pubsub := &inprocess.StreamManager{}
	pubsub.Init(logger, db)
	return &v1.V1{
		Dependencies: v1.Dependencies{
			DB:             db,
			Logger:         logger,
			Sonyflake:      sonyflake,
			Perms:          perms,
			Config:         conf,
			Middleware:     md,
			StorageBackend: attachments,
			Streams:        pubsub,
		},
	}
}

func TestCreateGuild(t *testing.T) {
	a := require.New(t)
	chatAPI := setupChatAPI()
	ctx := dummyContext(echo.New())

	name := "Harmony Development"
	resp, err := chatAPI.CreateGuild(ctx, &chatv1.CreateGuildRequest{
		GuildName: name,
	})
	a.NoError(err)
	a.NotZero(resp.GuildId)
	guild, err := chatAPI.DB.GetGuildByID(resp.GuildId)
	a.NoError(err)
	a.Equal(name, guild.GuildName)
	a.Empty(guild.PictureUrl)
}
