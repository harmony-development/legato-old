package types

import (
	"time"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
)

type IHarmonyDB interface {
	Migrate() error
	SessionExpireRoutine()
	CreateGuild(owner, id, channelID uint64, guildName, picture string) (*GuildData, error)
	DeleteGuild(guildID uint64) error
	GetOwner(guildID uint64) (uint64, error)
	IsOwner(guildID, userID uint64) (bool, error)
	CreateInvite(guildID uint64, possibleUses int32, name string) (*InviteData, error)
	UpdateChannelInformation(guildID, channelID uint64, name *string, metadata *harmonytypesv1.Metadata) error
	AddMemberToGuild(userID, guildID uint64) error
	AddChannelToGuild(guildID, channelID uint64, channelName string, previous, next *uint64, kind ChannelKind, md *harmonytypesv1.Metadata) (c ChannelData, err error)
	DeleteChannelFromGuild(guildID, channelID uint64) error

	UpdateTextMessage(messageID uint64, content string) (time.Time, error)

	DeleteMessage(messageID uint64) error

	ResolveGuildID(inviteID string) (uint64, error)
	IncrementInvite(inviteID string) error
	DeleteInvite(inviteID string) error
	SessionToUserID(session string) (uint64, error)
	LocalUserIDToForeignUserID(id uint64) (uint64, string, error)
	UserInGuild(userID, guildID uint64) (bool, error)

	GetHostQueue(host string) ([]byte, error)
	SetHostQueue(host string, data []byte) error

	AddMessage(guildID, channelID, messageID uint64, authorID uint64, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, content *harmonytypesv1.Content) (time.Time, error)
	GetMessage(messageID uint64) (*MessageData, error)
	GetMessages(channelID uint64) ([]*MessageData, error)
	GetMessagesBefore(channelID uint64, date time.Time) ([]*MessageData, error)

	UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) error
	GetGuildPicture(guildID uint64) (string, error)
	GetInvites(guildID uint64) ([]*InviteData, error)
	DeleteMember(guildID, userID uint64) error
	BanUser(guildID, userID uint64) error
	IsBanned(guildID, userID uint64) (bool, error)
	UnbanUser(guildID, userID uint64) error
	GetLocalGuilds(userID uint64) ([]uint64, error)
	ChannelsForGuild(guildID uint64) ([]*ChannelData, error)
	MembersInGuild(guildID uint64) ([]uint64, error)
	CountMembersInGuild(guildID uint64) (int64, error)
	GetUserByEmail(email string) (UserData, error)
	GetUserByID(userID uint64) (UserData, error)
	AddSession(userID uint64, session string) error
	ExtendSession(session string) error
	GetLocalUserForForeignUser(userID uint64, host string) (uint64, error)
	AddLocalUser(userID uint64, email, username string, passwordHash []byte) error
	AddForeignUser(host string, userID, localUserID uint64, username, avatar string) error
	EmailExists(email string) (bool, error)
	ExpireSessions() error
	GetAvatar(userID uint64) (*string, error)
	HasGuildWithID(guildID uint64) (bool, error)
	HasChannelWithID(guildID, channelID uint64) (bool, error)
	HasMessageWithID(messageID uint64) (bool, error)
	GetGuildByID(guildID uint64) (*GuildData, error)
	SetStatus(userID uint64, status harmonytypesv1.UserStatus) error
	SetUsername(userID uint64, username string) error
	SetAvatar(userID uint64, avatar string) error
	SetIsBot(userID uint64, isBot bool) error
	GetUserMetadata(userID uint64, appID string) (string, error)
	GetGuildList(userID uint64) ([]*GuildListEntryData, error)
	GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error)
	AddGuildToList(userID, guildID uint64, homeServer string) error
	MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error
	GetChannelListPosition(channelID uint64) (string, error)
	MoveChannel(channelID uint64, previousID, nextID *uint64) error
	RemoveGuildFromList(userID, guildID uint64, homeServer string) error
	UserIsLocal(userID uint64) (isLocal bool, err error)
	CreateEmotePack(userID, packID uint64, packName string) error
	IsPackOwner(userID, packID uint64) (bool, error)
	GetMessageOwner(messageID uint64) (uint64, error)
	AddEmoteToPack(packID uint64, imageID string, name string) error
	DeleteEmoteFromPack(packID uint64, emoteID string) error
	DeleteEmotePack(packID uint64) error
	GetEmotePacks(userID uint64) ([]*EmotePackData, error)
	GetEmotePackEmotes(packID uint64) ([]*EmoteData, error)
	DequipEmotePack(userID, packID uint64) error
	AddRoleToGuild(guildID, roleID uint64, name string, color int, hoist, pingable bool) error
	RemoveRoleFromGuild(guildID, roleID uint64) error
	MoveRole(guildID, roleID, previousRole, nextRole uint64) error
	GetGuildRoles(guildID uint64) ([]*RoleData, error)
	SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []PermissionsNode) error
	GetPermissions(roleID uint64) (permissions []PermissionsNode, err error)
	GetPermissionsData(guildID uint64) (PermissionsData, error)
	RolesForUser(guildID, userID uint64) ([]uint64, error)
	ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) error

	ModifyRole(roleID uint64, name string, color int, hoist, pingable, updateName, updateColor, updateHoist, updatePingable bool) error

	DeleteFileMeta(fileID string) error
	GetFileIDByHash(hash []byte) (string, error)
	AddFileHash(fileID string, hash []byte) error
	SetFileMetadata(fileID string, contentType, name string, size int) error
	GetFileMetadata(fileID string) (*FileData, error)
	GetFirstChannel(guildID uint64) (uint64, error)
}
