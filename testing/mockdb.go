package testing

import (
	"database/sql"
	"time"

	"harmony-server/server/db/queries"
)

type MockDB struct {
}

func (db *MockDB) Migrate() error {
	panic("implement me")
}

func (db *MockDB) SessionExpireRoutine() {
	panic("implement me")
}

func (db *MockDB) DeleteGuild(guildID uint64) error {
	panic("implement me")
}

func (db *MockDB) GetOwner(guildID uint64) (uint64, error) {
	panic("implement me")
}

func (db *MockDB) IsOwner(guildID uint64, userID uint64) (bool, error) {
	panic("implement me")
}

func (db *MockDB) CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error) {
	panic("implement me")
}

func (db *MockDB) AddMemberToGuild(userID uint64, guildID uint64) error {
	panic("implement me")
}

func (db *MockDB) AddChannelToGuild(guildID uint64, channelName string) (queries.Channel, error) {
	panic("implement me")
}

func (db *MockDB) DeleteChannelFromGuild(guildID, channelID uint64) error {
	panic("implement me")
}

func (db *MockDB) AddMessage(channelID, guildID, userID uint64, message string, attachments []string, embeds, actions [][]byte) (*queries.Message, error) {
	panic("implement me")
}

func (db *MockDB) DeleteMessage(messageID uint64, channelID uint64, guildID uint64) error {
	panic("implement me")
}

func (db *MockDB) GetMessageOwner(messageID uint64) (uint64, error) {
	panic("implement me")
}

func (db *MockDB) ResolveGuildID(inviteID string) (uint64, error) {
	panic("implement me")
}

func (db *MockDB) IncrementInvite(inviteID string) error {
	panic("implement me")
}

func (db *MockDB) DeleteInvite(inviteID string) error {
	panic("implement me")
}

func (db *MockDB) SessionToUserID(session string) (uint64, error) {
	panic("implement me")
}

func (db *MockDB) UserInGuild(userID uint64, guildID uint64) (bool, error) {
	panic("implement me")
}

func (db *MockDB) GetAttachments(messageID uint64) ([]string, error) {
	panic("implement me")
}

func (db *MockDB) GetMessageDate(messageID uint64) (time.Time, error) {
	panic("implement me")
}

func (db *MockDB) GetMessages(guildID uint64, channelID uint64) ([]queries.Message, error) {
	panic("implement me")
}

func (db *MockDB) GetMessagesBefore(guildID uint64, channelID uint64, date time.Time) ([]queries.Message, error) {
	panic("implement me")
}

func (db *MockDB) UpdateGuildName(guildID uint64, newName string) error {
	panic("implement me")
}

func (db *MockDB) GetGuildPicture(guildID uint64) (string, error) {
	panic("implement me")
}

func (db *MockDB) SetGuildPicture(guildID uint64, pictureURL string) error {
	panic("implement me")
}

func (db *MockDB) AddAttachments(messageID uint64, attachments []string) error {
	panic("implement me")
}

func (db *MockDB) GetInvites(guildID uint64) ([]queries.Invite, error) {
	panic("implement me")
}

func (db *MockDB) DeleteMember(guildID, userID uint64) error {
	panic("implement me")
}

func (db *MockDB) GuildsForUser(userID uint64) ([]uint64, error) {
	panic("implement me")
}

func (db *MockDB) ChannelsForGuild(guildID uint64) ([]queries.Channel, error) {
	panic("implement me")
}

func (db *MockDB) MembersInGuild(guildID uint64) ([]uint64, error) {
	panic("implement me")
}

func (db *MockDB) GetMessage(messageID uint64) (queries.Message, error) {
	panic("implement me")
}

func (db *MockDB) GetUserByEmail(email string) (queries.GetUserByEmailRow, error) {
	panic("implement me")
}

func (db *MockDB) GetUserByID(userID uint64) (queries.GetUserRow, error) {
	panic("implement me")
}

func (db *MockDB) AddSession(userID uint64, session string) error {
	return nil
}

func (db *MockDB) GetLocalUserForForeignUser(userID uint64, homeserver string) (uint64, error) {
	panic("implement me")
}

func (db *MockDB) AddLocalUser(userID uint64, email, username string, passwordHash []byte) error {
	return nil
}

func (db *MockDB) AddForeignUser(homeServer string, userID, localUserID uint64, username, avatar string) (uint64, error) {
	panic("implement me")
}

func (db *MockDB) EmailExists(email string) bool {
	return false
}

func (db *MockDB) ExpireSessions() error {
	panic("implement me")
}

func (db *MockDB) UpdateUsername(userID uint64, username string) error {
	panic("implement me")
}

func (db *MockDB) GetAvatar(userID uint64) (sql.NullString, error) {
	panic("implement me")
}

func (db *MockDB) UpdateAvatar(userID uint64, avatar string) error {
	panic("implement me")
}

func (db *MockDB) HasGuildWithID(guildID uint64) (bool, error) {
	panic("implement me")
}

func (db *MockDB) HasChannelWithID(guildID, channelID uint64) (bool, error) {
	panic("implement me")
}

func (db *MockDB) AddFileHash(fileID string, hash []byte) error {
	panic("implement me")
}

func (db *MockDB) GetFileIDFromHash(hash []byte) (string, error) {
	panic("implement me")
}

func (db *MockDB) GetGuildByID(guildID uint64) (queries.Guild, error) {
	panic("implement me")
}

func (db *MockDB) UpdateMessage(messageID uint64, content *string, embeds, actions *[][]byte) (time.Time, error) {
	panic("implement me")
}

func (db *MockDB) SetStatus(userID uint64, status queries.Userstatus) error {
	panic("implement me")
}

// CreateGuild creates a standard guild
func (db *MockDB) CreateGuild(owner, id uint64, guildName string, picture string) (*queries.Guild, error) {
	return &queries.Guild{
		OwnerID:    owner,
		GuildID:    id,
		GuildName:  guildName,
		PictureUrl: picture,
	}, nil
}
