package types

import (
	"database/sql"
	"time"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
)

type DummyDB struct{}

func (d DummyDB) Migrate() error {
	panic("unimplemented")
}

func (d DummyDB) SessionExpireRoutine() {
	panic("unimplemented")
}

func (d DummyDB) CreateGuild(owner, id, channelID uint64, guildName, picture string) (*queries.Guild, error) {
	panic("unimplemented")
}

func (d DummyDB) DeleteGuild(guildID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) GetOwner(guildID uint64) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) IsOwner(guildID, userID uint64) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error) {
	panic("unimplemented")
}

func (d DummyDB) UpdateChannelInformation(guildID, channelID uint64, name string, updateName bool, metadata *harmonytypesv1.Metadata, updateMetadata bool) error {
	panic("unimplemented")
}

func (d DummyDB) AddMemberToGuild(userID, guildID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) AddChannelToGuild(guildID uint64, channelName string, previous, next uint64, category bool, md *harmonytypesv1.Metadata) (queries.Channel, error) {
	panic("unimplemented")
}

func (d DummyDB) DeleteChannelFromGuild(guildID, channelID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64, md *harmonytypesv1.Metadata) (*queries.Message, error) {
	panic("unimplemented")
}

func (d DummyDB) DeleteMessage(messageID, channelID, guildID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) GetMessageOwner(messageID uint64) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) ResolveGuildID(inviteID string) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) IncrementInvite(inviteID string) error {
	panic("unimplemented")
}

func (d DummyDB) DeleteInvite(inviteID string) error {
	panic("unimplemented")
}

func (d DummyDB) SessionToUserID(session string) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) UserInGuild(userID, guildID uint64) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) GetMessageDate(messageID uint64) (time.Time, error) {
	panic("unimplemented")
}

func (d DummyDB) GetMessages(guildID, channelID uint64) ([]queries.Message, error) {
	panic("unimplemented")
}

func (d DummyDB) GetMessagesBefore(guildID, channelID uint64, date time.Time) ([]queries.Message, error) {
	panic("unimplemented")
}

func (d DummyDB) UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) error {
	panic("unimplemented")
}

func (d DummyDB) GetGuildPicture(guildID uint64) (string, error) {
	panic("unimplemented")
}

func (d DummyDB) GetInvites(guildID uint64) ([]queries.Invite, error) {
	panic("unimplemented")
}

func (d DummyDB) DeleteMember(guildID, userID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) GetLocalGuilds(userID uint64) ([]uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) ChannelsForGuild(guildID uint64) ([]queries.Channel, error) {
	panic("unimplemented")
}

func (d DummyDB) MembersInGuild(guildID uint64) ([]uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) CountMembersInGuild(guildID uint64) (int64, error) {
	panic("unimplemented")
}

func (d DummyDB) GetMessage(messageID uint64) (queries.Message, error) {
	panic("unimplemented")
}

func (d DummyDB) GetUserByEmail(email string) (queries.GetUserByEmailRow, error) {
	panic("unimplemented")
}

func (d DummyDB) GetUserByID(userID uint64) (queries.GetUserRow, error) {
	panic("unimplemented")
}

func (d DummyDB) AddSession(userID uint64, session string) error {
	panic("unimplemented")
}

func (d DummyDB) GetLocalUserForForeignUser(userID uint64, homeserver string) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) AddLocalUser(userID uint64, email, username string, passwordHash []byte) error {
	panic("unimplemented")
}

func (d DummyDB) AddForeignUser(homeServer string, userID, localUserID uint64, username, avatar string) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) EmailExists(email string) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) ExpireSessions() error {
	panic("unimplemented")
}

func (d DummyDB) UpdateUsername(userID uint64, username string) error {
	panic("unimplemented")
}

func (d DummyDB) GetAvatar(userID uint64) (sql.NullString, error) {
	panic("unimplemented")
}

func (d DummyDB) UpdateAvatar(userID uint64, avatar string) error {
	panic("unimplemented")
}

func (d DummyDB) HasGuildWithID(guildID uint64) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) HasChannelWithID(guildID, channelID uint64) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) HasMessageWithID(guildID, channelID, messageID uint64) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) GetGuildByID(guildID uint64) (queries.Guild, error) {
	panic("unimplemented")
}

func (d DummyDB) UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string, metadata *harmonytypesv1.Metadata, updateMetadata bool) (time.Time, error) {
	panic("unimplemented")
}

func (d DummyDB) SetStatus(userID uint64, status harmonytypesv1.UserStatus) error {
	panic("unimplemented")
}

func (d DummyDB) SetUsername(userID uint64, username string) error {
	panic("unimplemented")
}

func (d DummyDB) SetAvatar(userID uint64, avatar string) error {
	panic("unimplemented")
}

func (d DummyDB) SetIsBot(userID uint64, isBot bool) error {
	panic("unimplemented")
}

func (d DummyDB) GetUserMetadata(userID uint64, appID string) (string, error) {
	panic("unimplemented")
}

func (d DummyDB) GetNonceInfo(nonce string) (queries.GetNonceInfoRow, error) {
	panic("unimplemented")
}

func (d DummyDB) AddNonce(nonce string, userID uint64, homeServer string) error {
	panic("unimplemented")
}

func (d DummyDB) GetGuildList(userID uint64) ([]queries.GetGuildListRow, error) {
	panic("unimplemented")
}

func (d DummyDB) GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error) {
	panic("unimplemented")
}

func (d DummyDB) AddGuildToList(userID, guildID uint64, homeServer string) error {
	panic("unimplemented")
}

func (d DummyDB) MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error {
	panic("unimplemented")
}

func (d DummyDB) GetChannelListPosition(guildID, channelID uint64) (string, error) {
	panic("unimplemented")
}

func (d DummyDB) MoveChannel(guildID, channelID, previousID, nextID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) RemoveGuildFromList(userID, guildID uint64, homeServer string) error {
	panic("unimplemented")
}

func (d DummyDB) UserIsLocal(userID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) CreateEmotePack(userID, packID uint64, packName string) error {
	panic("unimplemented")
}

func (d DummyDB) IsPackOwner(userID, packID uint64) (bool, error) {
	panic("unimplemented")
}

func (d DummyDB) AddEmoteToPack(packID uint64, imageID string, name string) error {
	panic("unimplemented")
}

func (d DummyDB) DeleteEmoteFromPack(packID uint64, imageID string) error {
	panic("unimplemented")
}

func (d DummyDB) DeleteEmotePack(packID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) GetEmotePacks(userID uint64) ([]queries.GetEmotePacksRow, error) {
	panic("unimplemented")
}

func (d DummyDB) GetEmotePackEmotes(packID uint64) ([]queries.GetEmotePackEmotesRow, error) {
	panic("unimplemented")
}

func (d DummyDB) DequipEmotePack(userID, packID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) AddRoleToGuild(guildID uint64, role *chatv1.Role) error {
	panic("unimplemented")
}

func (d DummyDB) RemoveRoleFromGuild(guildID, roleID uint64) error {
	panic("unimplemented")
}

func (d DummyDB) GetRolePositions(guildID, before, previous uint64) (pos string, retErr error) {
	panic("unimplemented")
}

func (d DummyDB) MoveRole(guildID, roleID, beforeRole, previousRole uint64) error {
	panic("unimplemented")
}

func (d DummyDB) GetGuildRoles(guildID uint64) ([]*chatv1.Role, error) {
	panic("unimplemented")
}

func (d DummyDB) SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []PermissionsNode) error {
	panic("unimplemented")
}

func (d DummyDB) GetPermissions(guildID uint64, channelID uint64, roleID uint64) (permissions []PermissionsNode, err error) {
	panic("unimplemented")
}

func (d DummyDB) GetPermissionsData(guildID uint64) (PermissionsData, error) {
	panic("unimplemented")
}

func (d DummyDB) RolesForUser(guildID, userID uint64) ([]uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) error {
	panic("unimplemented")
}

func (d DummyDB) ModifyRole(guildID, roleID uint64, name string, color int32, hoist, pingable, updateName, updateColor, updateHoist, updatePingable bool) error {
	panic("unimplemented")
}

func (d DummyDB) DeleteFileMeta(fileID string) error {
	panic("unimplemented")
}

func (d DummyDB) GetFileIDByHash(hash []byte) (string, error) {
	panic("unimplemented")
}

func (d DummyDB) AddFileHash(fileID string, hash []byte) error {
	panic("unimplemented")
}

func (d DummyDB) SetFileMetadata(fileID string, contentType, name string, size int32) error {
	panic("unimplemented")
}

func (d DummyDB) GetFileMetadata(fileID string) (queries.GetFileMetadataRow, error) {
	panic("unimplemented")
}

func (d DummyDB) GetFirstChannel(guildID uint64) (uint64, error) {
	panic("unimplemented")
}

func (d DummyDB) ExtendSession(session string) error {
	panic("unimplemented")
}
