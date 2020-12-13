package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	corev1 "github.com/harmony-development/legato/gen/core"
	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
	"github.com/ztrue/tracerr"

	lru "github.com/hashicorp/golang-lru"
	_ "github.com/lib/pq"
)

var ErrNotLocal = errors.New("User is not local")

// HarmonyDB is a wrapper for the SQL HarmonyDB
type HarmonyDB struct {
	*sql.DB
	queries      *queries.Queries
	Logger       logger.ILogger
	Config       *config.Config
	OwnerCache   *lru.Cache
	SessionCache *lru.Cache
	Sonyflake    *sonyflake.Sonyflake
}

type PermissionsNode struct {
	Node  string
	Allow bool
}

func (p *PermissionsNode) Deserialize(s string) (ok bool) {
	trimmed := strings.Split(strings.TrimSuffix(strings.TrimPrefix(s, "("), ")"), ",")

	if len(trimmed) != 3 {
		return false
	}

	if trimmed[2] == "t" {
		p.Allow = true
	} else {
		p.Allow = false
	}

	p.Node = trimmed[1]

	return true
}

func (p PermissionsNode) Serialize() string {
	node := "f"
	if p.Allow {
		node = "t"
	}
	return fmt.Sprintf("(%s,%s)", p.Node, node)
}

type PermissionsData struct {
	Roles      map[uint64][]PermissionsNode
	Categories map[uint64][]uint64
	Channels   map[uint64]map[uint64][]PermissionsNode
}

type IHarmonyDB interface {
	Migrate() error
	SessionExpireRoutine()
	CreateGuild(owner, id, channelID uint64, guildName, picture string) (*queries.Guild, error)
	DeleteGuild(guildID uint64) error
	GetOwner(guildID uint64) (uint64, error)
	IsOwner(guildID, userID uint64) (bool, error)
	CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error)
	SetChannelName(guildID, channelID uint64, name string) error
	AddMemberToGuild(userID, guildID uint64) error
	AddChannelToGuild(guildID uint64, channelName string, previous, next uint64, category bool) (queries.Channel, error)
	DeleteChannelFromGuild(guildID, channelID uint64) error
	AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64) (*queries.Message, error)
	DeleteMessage(messageID, channelID, guildID uint64) error
	GetMessageOwner(messageID uint64) (uint64, error)
	ResolveGuildID(inviteID string) (uint64, error)
	IncrementInvite(inviteID string) error
	DeleteInvite(inviteID string) error
	SessionToUserID(session string) (uint64, error)
	UserInGuild(userID, guildID uint64) (bool, error)
	GetMessageDate(messageID uint64) (time.Time, error)
	GetMessages(guildID, channelID uint64) ([]queries.Message, error)
	GetMessagesBefore(guildID, channelID uint64, date time.Time) ([]queries.Message, error)
	UpdateGuildName(guildID uint64, newName string) error
	GetGuildPicture(guildID uint64) (string, error)
	SetGuildPicture(guildID uint64, pictureURL string) error
	GetInvites(guildID uint64) ([]queries.Invite, error)
	DeleteMember(guildID, userID uint64) error
	GetLocalGuilds(userID uint64) ([]uint64, error)
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
	GetGuildByID(guildID uint64) (queries.Guild, error)
	UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string) (time.Time, error)
	SetStatus(userID uint64, status profilev1.UserStatus) error
	GetUserMetadata(userID uint64, appID string) (string, error)
	GetNonceInfo(nonce string) (queries.GetNonceInfoRow, error)
	AddNonce(nonce string, userID uint64, homeServer string) error
	GetGuildList(userID uint64) ([]queries.GetGuildListRow, error)
	GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error)
	AddGuildToList(userID, guildID uint64, homeServer string) error
	MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error
	GetChannelListPosition(guildID, channelID uint64) (string, error)
	MoveChannel(guildID, channelID, previousID, nextID uint64) error
	RemoveGuildFromList(userID, guildID uint64, homeServer string) error
	UserIsLocal(userID uint64) error
	CreateEmotePack(userID, packID uint64, packName string) error
	IsPackOwner(userID, packID uint64) (bool, error)
	AddEmoteToPack(packID uint64, imageID string, name string) error
	DeleteEmoteFromPack(packID uint64, imageID string) error
	DeleteEmotePack(packID uint64) error
	GetEmotePacks(userID uint64) ([]queries.GetEmotePacksRow, error)
	GetEmotePackEmotes(packID uint64) ([]queries.GetEmotePackEmotesRow, error)
	DequipEmotePack(userID, packID uint64) error
	AddRoleToGuild(guildID uint64, role *corev1.Role) error
	RemoveRoleFromGuild(guildID, roleID uint64) error
	GetRolePositions(guildID, before, previous uint64) (pos string, retErr error)
	MoveRole(guildID, roleID, beforeRole, previousRole uint64) error
	GetGuildRoles(guildID uint64) ([]*corev1.Role, error)
	SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []PermissionsNode) error
	GetPermissions(guildID uint64, channelID uint64, roleID uint64) (permissions []PermissionsNode, err error)
	GetPermissionsData(guildID uint64) (PermissionsData, error)
	RolesForUser(guildID, userID uint64) ([]uint64, error)
	ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) error
	ModifyRole(guildID, roleID uint64, name string, color int32, hoist, pingable, updateName, updateColor, updateHoist, updatePingable bool) error
	DeleteFileMeta(fileID string) error
	GetFileIDByHash(hash []byte) (string, error)
	AddFileHash(fileID string, hash []byte) error
	SetFileMetadata(fileID string, contentType, name string, size int32) error
	GetFileMetadata(fileID string) (queries.GetFileMetadataRow, error)
}

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (*HarmonyDB, error) {
	db := &HarmonyDB{}
	db.Config = cfg
	db.Logger = logger
	db.Sonyflake = idgen
	var err error
	if db.DB, err = sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Host,
		cfg.Database.Port,
		map[bool]string{true: "enable", false: "disable"}[cfg.Database.SSL],
	)); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if err = db.Ping(); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if err = db.Migrate(); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if db.queries, err = queries.Prepare(context.Background(), db); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if db.OwnerCache, err = lru.New(cfg.Server.Policies.MaximumCacheSizes.Owner); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if db.SessionCache, err = lru.New(cfg.Server.Policies.MaximumCacheSizes.Sessions); err != nil {
		return nil, tracerr.Wrap(err)
	}
	go db.SessionExpireRoutine()
	return db, nil
}
