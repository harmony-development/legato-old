package v1

import (
	"testing"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/api/chat/v1/pubsub_backends/inprocess"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/responses"
	"github.com/harmony-development/legato/server/test"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/require"
)

func setupChatAPI() *V1 {
	conf := test.DefaultConf()
	logger := test.MockLogger{}
	db := test.NewMockDB()
	perms := permissions.NewManager(db, logger)
	md := middleware.New(middleware.Dependencies{
		Logger: logger,
		DB:     db,
		Perms:  perms,
	})
	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{})
	attachments := test.NewMockAttachmentsBackend()
	pubsub := &inprocess.StreamManager{}
	pubsub.Init(logger, db)
	return &V1{
		Dependencies: Dependencies{
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
	testTable := []struct {
		name          string
		picture       string
		expectedError string
	}{
		{
			name: "Harmony Development",
		},
		{
			name:          "this is a very long guild name, far too long to be acceptable to be used anywhere.",
			expectedError: responses.NameTooLong,
		},
	}
	a := require.New(t)
	chatAPI := setupChatAPI()
	c := test.DummyContext(echo.New())
	ctx := middleware.HarmonyContext{
		Context: c,
		UserID:  12345,
	}
	for _, testCase := range testTable {
		createdGuild, err := chatAPI.CreateGuild(ctx, &chatv1.CreateGuildRequest{
			GuildName: testCase.name,
		})
		if testCase.expectedError != "" {
			a.Error(err)
			a.Equal(testCase.expectedError, err.Error())
		} else {
			a.NoError(err)
			a.NotZero(createdGuild.GuildId)
			guild, err := chatAPI.DB.GetGuildByID(createdGuild.GuildId)
			a.NoError(err)
			a.Equal(testCase.name, guild.GuildName)
			a.Empty(guild.PictureUrl)
			channels, err := chatAPI.DB.ChannelsForGuild(createdGuild.GuildId)
			a.NoError(err)
			a.Len(channels, 1)
			inGuild, err := chatAPI.DB.UserInGuild(12345, createdGuild.GuildId)
			a.NoError(err)
			a.True(inGuild)
		}
	}
}

func TestJoinLeave(t *testing.T) {
	a := require.New(t)
	chatAPI := setupChatAPI()
	c := test.DummyContext(echo.New())
	ctx := middleware.HarmonyContext{
		Context: c,
		UserID:  22321123,
	}
	_, err := chatAPI.DB.CreateGuild(12345, 727, 420, "Harmony", "")
	a.NoError(err)
	inv, err := chatAPI.DB.CreateInvite(727, -1, "harmony")
	a.NoError(err)
	joinResp, err := chatAPI.JoinGuild(ctx, &chatv1.JoinGuildRequest{
		InviteId: inv.InviteID,
	})
	a.NoError(err)
	a.Equal(727, joinResp.GuildId)
	inGuild, err := chatAPI.DB.UserInGuild(ctx.UserID, 727)
	a.NoError(err)
	a.True(inGuild)
	_, err = chatAPI.JoinGuild(ctx, &chatv1.JoinGuildRequest{
		InviteId: inv.InviteID,
	})
	a.Error(err)
}
