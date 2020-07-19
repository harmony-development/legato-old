package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/logger"

	lru "github.com/hashicorp/golang-lru"
	_ "github.com/lib/pq"
)

// HarmonyDB is a wrapper for the SQL HarmonyDB
type HarmonyDB struct {
	*sql.DB
	queries      *queries.Queries
	Logger       logger.ILogger
	Config       *config.Config
	OwnerCache   *lru.Cache
	SessionCache *lru.Cache
}

type IHarmonyDB interface {
	Migrate() error
	SessionExpireRoutine()
	CreateGuild(owner, id uint64, guildName string, picture string) (*queries.Guild, error)
	DeleteGuild(guildID uint64) error
	GetOwner(guildID uint64) (uint64, error)
	IsOwner(guildID uint64, userID uint64) (bool, error)
	CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error)
	SetChannelName(guildID, channelID uint64, name string) error
	AddMemberToGuild(userID uint64, guildID uint64) error
	AddChannelToGuild(guildID uint64, channelName string) (queries.Channel, error)
	DeleteChannelFromGuild(guildID, channelID uint64) error
	AddMessage(channelID, guildID, userID uint64, message string, attachments []string, embeds, actions [][]byte) (*queries.Message, error)
	DeleteMessage(messageID uint64, channelID uint64, guildID uint64) error
	GetMessageOwner(messageID uint64) (uint64, error)
	ResolveGuildID(inviteID string) (uint64, error)
	IncrementInvite(inviteID string) error
	DeleteInvite(inviteID string) error
	SessionToUserID(session string) (uint64, error)
	UserInGuild(userID uint64, guildID uint64) (bool, error)
	GetAttachments(messageID uint64) ([]string, error)
	GetMessageDate(messageID uint64) (time.Time, error)
	GetMessages(guildID uint64, channelID uint64) ([]queries.Message, error)
	GetMessagesBefore(guildID uint64, channelID uint64, date time.Time) ([]queries.Message, error)
	UpdateGuildName(guildID uint64, newName string) error
	GetGuildPicture(guildID uint64) (string, error)
	SetGuildPicture(guildID uint64, pictureURL string) error
	AddAttachments(messageID uint64, attachments []string) error
	GetInvites(guildID uint64) ([]queries.Invite, error)
	DeleteMember(guildID, userID uint64) error
	GuildsForUser(userID uint64) ([]uint64, error)
	GuildsForUserWithData(userID uint64) ([]queries.Guild, error)
	ChannelsForGuild(guildID uint64) ([]queries.Channel, error)
	MembersInGuild(guildID uint64) ([]uint64, error)
	GetMessage(messageID uint64) (queries.Message, error)
	GetUserByEmail(email string) (queries.GetUserByEmailRow, error)
	GetUserByID(userID uint64) (queries.GetUserRow, error)
	AddSession(userID uint64, session string) error
	GetLocalUserForForeignUser(userID uint64, homeserver string) (uint64, error)
	AddLocalUser(userID uint64, email, username string, passwordHash []byte) error
	AddForeignUser(homeServer string, userID, localUserID uint64, username, avatar string) (uint64, error)
	EmailExists(email string) (bool, error)
	ExpireSessions() error
	UpdateUsername(userID uint64, username string) error
	GetAvatar(userID uint64) (sql.NullString, error)
	UpdateAvatar(userID uint64, avatar string) error
	HasGuildWithID(guildID uint64) (bool, error)
	HasChannelWithID(guildID, channelID uint64) (bool, error)
	HasMessageWithID(guildID, channelID, messageID uint64) (bool, error)
	AddFileHash(fileID string, hash []byte) error
	GetFileIDFromHash(hash []byte) (string, error)
	GetGuildByID(guildID uint64) (queries.Guild, error)
	UpdateMessage(messageID uint64, content *string, embeds, actions *[][]byte) (time.Time, error)
	SetStatus(userID uint64, status profilev1.UserStatus) error
	GetUserMetadata(userID uint64, appID string) (string, error)
	GetNonceInfo(nonce string) (queries.GetNonceInfoRow, error)
	AddNonce(nonce string, userID uint64, homeServer string) error
	GetGuildList(userID uint64) ([]queries.GetGuildListRow, error)
	GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error)
	AddGuildToList(userID, guildID uint64, homeServer string) error
	MoveGuild(userID uint64, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error
	RemoveGuildFromList(userID, guildID uint64, homeServer string) error
}

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger) (*HarmonyDB, error) {
	db := &HarmonyDB{}
	db.Config = cfg
	db.Logger = logger
	var err error
	if db.DB, err = sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cfg.DB.User,
		cfg.DB.Password,
		"harmony",
		cfg.DB.Host,
		cfg.DB.Port,
		map[bool]string{true: "enable", false: "disable"}[cfg.DB.SSL],
	)); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = db.Migrate(); err != nil {
		return nil, err
	}
	if db.queries, err = queries.Prepare(context.Background(), db); err != nil {
		return nil, err
	}
	if db.OwnerCache, err = lru.New(cfg.Server.OwnerCacheMax); err != nil {
		return nil, err
	}
	if db.SessionCache, err = lru.New(cfg.Server.SessionCacheMax); err != nil {
		return nil, err
	}
	go db.SessionExpireRoutine()
	return db, nil
}
