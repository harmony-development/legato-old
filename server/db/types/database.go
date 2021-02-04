package types

import (
	"database/sql"
	"time"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
)

type IHarmonyDB interface {
	Migrate() error
	SessionExpireRoutine()
	CreateGuild(owner, id, channelID uint64, guildName, picture string) (*queries.Guild, error)
	DeleteGuild(guildID uint64) error
	GetOwner(guildID uint64) (uint64, error)
	IsOwner(guildID, userID uint64) (bool, error)
	CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error)
	UpdateChannelInformation(guildID, channelID uint64, name string, updateName bool, metadata *harmonytypesv1.Metadata, updateMetadata bool) error
	AddMemberToGuild(userID, guildID uint64) error
	AddChannelToGuild(guildID uint64, channelName string, previous, next uint64, category bool, md *harmonytypesv1.Metadata) (queries.Channel, error)
	DeleteChannelFromGuild(guildID, channelID uint64) error
	AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64, md *harmonytypesv1.Metadata) (*queries.Message, error)
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
	UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) error
	GetGuildPicture(guildID uint64) (string, error)
	GetInvites(guildID uint64) ([]queries.Invite, error)
	DeleteMember(guildID, userID uint64) error
	GetLocalGuilds(userID uint64) ([]uint64, error)
	ChannelsForGuild(guildID uint64) ([]queries.Channel, error)
	MembersInGuild(guildID uint64) ([]uint64, error)
	CountMembersInGuild(guildID uint64) (int64, error)
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
	UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string, metadata *harmonytypesv1.Metadata, updateMetadata bool) (time.Time, error)
	SetStatus(userID uint64, status harmonytypesv1.UserStatus) error
	SetUsername(userID uint64, username string) error
	SetAvatar(userID uint64, avatar string) error
	SetIsBot(userID uint64, isBot bool) error
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
	AddRoleToGuild(guildID uint64, role *chatv1.Role) error
	RemoveRoleFromGuild(guildID, roleID uint64) error
	GetRolePositions(guildID, before, previous uint64) (pos string, retErr error)
	MoveRole(guildID, roleID, beforeRole, previousRole uint64) error
	GetGuildRoles(guildID uint64) ([]*chatv1.Role, error)
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
	GetFirstChannel(guildID uint64) (uint64, error)
}
