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

func setupChatAPI(t testing.TB) *V1 {
	cfg := test.DefaultConf()
	logger := test.MockLogger{T: t, Config: cfg}
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
			Config:         cfg,
			Middleware:     md,
			StorageBackend: attachments,
			Streams:        pubsub,
		},
	}
}

func TestCreateGuild(t *testing.T) {
	testTable := []struct {
		testName      string
		name          string
		picture       string
		expectedError string
	}{
		{
			testName: "A guild should successfully be created with just a name",
			name:     "Harmony Development",
		},
		{
			testName:      "A guild with a long name should not be able to be created",
			name:          "this is a very long guild name, far too long to be acceptable to be used anywhere.",
			expectedError: responses.NameTooLong,
		},
	}
	a := require.New(t)
	chatAPI := setupChatAPI(t)
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
	chatAPI := setupChatAPI(t)
	c := test.DummyContext(echo.New())
	ownerCTX := middleware.HarmonyContext{
		Context: c,
		UserID:  22321123,
	}
	memberCTX := middleware.HarmonyContext{
		Context: c,
		UserID:  12345,
	}
	bannedMemberCTX := middleware.HarmonyContext{
		Context: c,
		UserID:  54321,
	}
	_, _ = chatAPI.DB.CreateGuild(ownerCTX.UserID, 727, 420, "Harmony", "")
	inv, _ := chatAPI.DB.CreateInvite(727, -1, "harmony")
	_ = chatAPI.DB.BanUser(727, bannedMemberCTX.UserID)

	joinResp, err := chatAPI.JoinGuild(memberCTX, &chatv1.JoinGuildRequest{
		InviteId: inv.InviteID,
	})
	a.NoError(err)
	a.Equal(uint64(727), joinResp.GuildId)
	inGuild, err := chatAPI.DB.UserInGuild(memberCTX.UserID, 727)
	a.NoError(err)
	a.True(inGuild)
	_, err = chatAPI.JoinGuild(memberCTX, &chatv1.JoinGuildRequest{
		InviteId: inv.InviteID,
	})
	a.Error(err, "It should not allow a member to join a second time")
	_, err = chatAPI.JoinGuild(ownerCTX, &chatv1.JoinGuildRequest{
		InviteId: inv.InviteID,
	})
	a.Error(err, "It should not allow the owner to join their own guild")
	_, err = chatAPI.JoinGuild(bannedMemberCTX, &chatv1.JoinGuildRequest{
		InviteId: inv.InviteID,
	})
	a.Error(err, "It should not let banned people join a guild")
	_, err = chatAPI.LeaveGuild(memberCTX, &chatv1.LeaveGuildRequest{
		GuildId: 727,
	})
	a.NoError(err, "It should allow a member to leave the guild")
	_, err = chatAPI.LeaveGuild(memberCTX, &chatv1.LeaveGuildRequest{
		GuildId: 727,
	})
	a.EqualError(err, responses.NotJoined, "It should not allow a member to leave a guild they aren't in")
	_, err = chatAPI.LeaveGuild(ownerCTX, &chatv1.LeaveGuildRequest{
		GuildId: 727,
	})
	a.EqualError(err, responses.IsOwner, "It should not allow the owner to leave their guild")
}
