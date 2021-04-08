package types

import (
	"time"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
)

type IHarmonyDB interface {
	Migrate() error
	SessionExpireRoutine()
	CreateGuild(owner, id, channelID uint64, guildName, picture string) (*entgen.Guild, error)
	DeleteGuild(guildID uint64) error
	GetOwner(guildID uint64) (uint64, error)
	IsOwner(guildID, userID uint64) (bool, error)
	CreateInvite(guildID uint64, possibleUses int32, name string) (*entgen.Invite, error)
	UpdateChannelInformation(guildID, channelID uint64, name *string, metadata []byte) error
	AddMemberToGuild(userID, guildID uint64) error
	AddChannelToGuild(guildID, channelID uint64, channelName string, previous, next *uint64, kind ChannelKind, md []byte) (c entgen.Channel, err error)
	DeleteChannelFromGuild(guildID, channelID uint64) error

	AddTextMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, content string) (time.Time, error)
	AddFilesMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, files []*harmonytypesv1.Attachment) (time.Time, error)
	AddEmbedMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, embeds []*harmonytypesv1.Embed) (time.Time, error)

	UpdateTextMessage(messageID uint64, content string) (time.Time, error)

	DeleteMessage(messageID uint64) error

	ResolveGuildID(inviteID string) (uint64, error)
	IncrementInvite(inviteID string) error
	DeleteInvite(inviteID string) error
	SessionToUserID(session string) (uint64, error)
	UserInGuild(userID, guildID uint64) (bool, error)

	GetMessages(channelID uint64) ([]*entgen.Message, error)
	GetMessagesBefore(channelID uint64, date time.Time) ([]*entgen.Message, error)

	UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) error
	GetGuildPicture(guildID uint64) (string, error)
	GetInvites(guildID uint64) ([]*entgen.Invite, error)
	DeleteMember(guildID, userID uint64) error
	BanUser(guildID, userID uint64) error
	IsBanned(guildID, userID uint64) (bool, error)
	UnbanUser(guildID, userID uint64) error
	GetLocalGuilds(userID uint64) ([]uint64, error)
	ChannelsForGuild(guildID uint64) ([]*entgen.Channel, error)
	MembersInGuild(guildID uint64) ([]uint64, error)
	CountMembersInGuild(guildID uint64) (int64, error)
	GetMessage(messageID uint64) (*entgen.Message, error)
	GetUserByEmail(email string) (UserData, error)
	GetUserByID(userID uint64) (UserData, error)
	AddSession(userID uint64, session string) error
	ExtendSession(session string) error
	GetLocalUserForForeignUser(userID uint64, host string) (uint64, error)
	AddLocalUser(userID uint64, email, username string, passwordHash []byte) error
	AddForeignUser(host string, userID, localUserID uint64, username, avatar string) error
	EmailExists(email string) (bool, error)
	ExpireSessions() error
	UpdateUsername(userID uint64, username string) error
	GetAvatar(userID uint64) (*string, error)
	HasGuildWithID(guildID uint64) (bool, error)
	HasChannelWithID(guildID, channelID uint64) (bool, error)
	HasMessageWithID(messageID uint64) (bool, error)
	GetGuildByID(guildID uint64) (*entgen.Guild, error)
	SetStatus(userID uint64, status harmonytypesv1.UserStatus) error
	SetUsername(userID uint64, username string) error
	SetAvatar(userID uint64, avatar string) error
	SetIsBot(userID uint64, isBot bool) error
	GetUserMetadata(userID uint64, appID string) (string, error)
	GetGuildList(userID uint64) ([]*entgen.GuildListEntry, error)
	GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error)
	AddGuildToList(userID, guildID uint64, homeServer string) error
	MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error
	GetChannelListPosition(channelID uint64) (string, error)
	MoveChannel(channelID uint64, previousID, nextID *uint64) error
	RemoveGuildFromList(userID, guildID uint64, homeServer string) error
	UserIsLocal(userID uint64) error
	CreateEmotePack(userID, packID uint64, packName string) error
	IsPackOwner(userID, packID uint64) (bool, error)
	AddEmoteToPack(packID uint64, imageID string, name string) error
	DeleteEmoteFromPack(packID uint64, emoteID string) error
	DeleteEmotePack(packID uint64) error
	GetEmotePacks(userID uint64) ([]*entgen.EmotePack, error)
	GetEmotePackEmotes(packID uint64) ([]*entgen.Emote, error)
	DequipEmotePack(userID, packID uint64) error
	AddRoleToGuild(guildID, roleID uint64, name string, color int, hoist, pingable bool) error
	RemoveRoleFromGuild(guildID, roleID uint64) error
	MoveRole(guildID, roleID, previousRole, nextRole uint64) error
	GetGuildRoles(guildID uint64) ([]*entgen.Role, error)
	SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []PermissionsNode) error
	GetPermissions(roleID uint64) (permissions []PermissionsNode, err error)
	GetPermissionsData(guildID uint64) (PermissionsData, error)
	RolesForUser(guildID, userID uint64) ([]uint64, error)
	ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) error
	ModifyRole(roleID uint64, name *string, color *int, hoist, pingable *bool) error
	DeleteFileMeta(fileID string) error
	GetFileIDByHash(hash []byte) (string, error)
	AddFileHash(fileID string, hash []byte) error
	SetFileMetadata(fileID string, contentType, name string, size int) error
	GetFileMetadata(fileID string) (*entgen.File, error)
	GetFirstChannel(guildID uint64) (uint64, error)
}
