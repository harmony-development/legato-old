package test

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/db/types"
)

type User struct {
	id       uint64
	email    string
	username string
	password []byte
}

type Guild struct {
	id       uint64
	owner    uint64
	name     string
	picture  string
	members  map[uint64]struct{}
	channels map[uint64]struct{}
	bans     map[uint64]struct{}
}

type Channel struct {
	id       uint64
	name     string
	position string
	category bool
	messages []uint64
}

type Invite struct {
	name         string
	guildid      uint64
	possibleUses int32
	uses         int32
}

type Message struct {
	id          uint64
	author      uint64
	content     string
	attachments []string
	embeds      []byte
	actions     []byte
	overrides   []byte
	replyTo     int64
}

type MockDB struct {
	users         map[uint64]*User
	userByEmail   map[string]*User
	userBySession map[string]uint64
	guilds        map[uint64]*Guild
	channels      map[uint64]*Channel
	invites       map[string]*Invite
	messages      map[uint64]*Message
}

func NewMockDB() *MockDB {
	return &MockDB{
		users:         map[uint64]*User{},
		userByEmail:   map[string]*User{},
		userBySession: map[string]uint64{},
		guilds:        map[uint64]*Guild{},
		channels:      map[uint64]*Channel{},
		invites:       map[string]*Invite{},
		messages:      map[uint64]*Message{},
	}
}

func (d MockDB) Migrate() error {
	panic("unimplemented")
}

func (d MockDB) SessionExpireRoutine() {
	panic("unimplemented")
}

func (d MockDB) CreateGuild(owner, id, channelID uint64, guildName, picture string) (*queries.Guild, error) {
	d.channels[channelID] = &Channel{
		id:       channelID,
		name:     "general",
		messages: []uint64{},
	}
	d.guilds[id] = &Guild{
		id:       id,
		owner:    owner,
		name:     guildName,
		picture:  picture,
		members:  map[uint64]struct{}{owner: {}},
		channels: map[uint64]struct{}{channelID: {}},
		bans:     map[uint64]struct{}{},
	}
	return &queries.Guild{
		GuildID:    id,
		OwnerID:    owner,
		GuildName:  guildName,
		PictureUrl: picture,
		Metadata:   []byte{},
	}, nil
}

func (d MockDB) DeleteGuild(guildID uint64) error {
	panic("unimplemented")
}

func (d MockDB) GetOwner(guildID uint64) (uint64, error) {
	panic("unimplemented")
}

func (d MockDB) IsOwner(guildID, userID uint64) (bool, error) {
	if guild, ok := d.guilds[guildID]; !ok {
		return false, errors.New("guild not found")
	} else {
		return guild.owner == userID, nil
	}
}

func (d MockDB) CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error) {
	d.invites[name] = &Invite{
		name:         name,
		guildid:      guildID,
		possibleUses: possibleUses,
		uses:         0,
	}

	return queries.Invite{
		InviteID:     name,
		Uses:         0,
		PossibleUses: sql.NullInt32{Int32: possibleUses},
		GuildID:      guildID,
	}, nil
}

func (d MockDB) UpdateChannelInformation(guildID, channelID uint64, name string, updateName bool, metadata *harmonytypesv1.Metadata, updateMetadata bool) error {
	panic("unimplemented")
}

func (d MockDB) AddMemberToGuild(userID, guildID uint64) error {
	guild, ok := d.guilds[guildID]
	if !ok {
		return errors.New("guild not found")
	}
	guild.members[userID] = struct{}{}
	return nil
}

func (d MockDB) AddChannelToGuild(guildID uint64, channelName string, previous, next uint64, category bool, md *harmonytypesv1.Metadata) (queries.Channel, error) {
	panic("unimplemented")
}

func (d MockDB) DeleteChannelFromGuild(guildID, channelID uint64) error {
	panic("unimplemented")
}

func (d MockDB) AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64, md *harmonytypesv1.Metadata) (*queries.Message, error) {
	if channel, ok := d.channels[channelID]; !ok {
		return nil, errors.New("channel not found")
	} else {
		channel.messages = append(channel.messages, messageID)
		d.messages[messageID] = &Message{
			id:          messageID,
			author:      userID,
			content:     message,
			attachments: attachments,
			embeds:      embeds,
			actions:     actions,
			overrides:   overrides,
			replyTo:     replyTo.Int64,
		}
		return &queries.Message{
			MessageID:   messageID,
			GuildID:     guildID,
			ChannelID:   channelID,
			UserID:      userID,
			CreatedAt:   time.Now(),
			EditedAt:    sql.NullTime{},
			Content:     message,
			Embeds:      embeds,
			Actions:     actions,
			Overrides:   overrides,
			ReplyToID:   sql.NullInt64{},
			Attachments: attachments,
			Metadata:    []byte{},
		}, nil
	}
}

func (d MockDB) DeleteMessage(messageID, channelID, guildID uint64) error {
	panic("unimplemented")
}

func (d MockDB) GetMessageOwner(messageID uint64) (uint64, error) {
	panic("unimplemented")
}

func (d MockDB) ResolveGuildID(inviteID string) (uint64, error) {
	inv, ok := d.invites[inviteID]
	if !ok {
		return 0, errors.New("invite not found")
	}
	return inv.guildid, nil
}

func (d MockDB) IncrementInvite(inviteID string) error {
	if inv, ok := d.invites[inviteID]; !ok {
		return errors.New("invite not found")
	} else {
		inv.uses++
		return nil
	}
}

func (d MockDB) DeleteInvite(inviteID string) error {
	panic("unimplemented")
}

func (d MockDB) SessionToUserID(session string) (uint64, error) {
	if d.userBySession[session] == 0 {
		return 0, errors.New("session does not exist")
	}
	return d.userBySession[session], nil
}

func (d MockDB) UserInGuild(userID, guildID uint64) (bool, error) {
	guild, ok := d.guilds[guildID]
	if !ok {
		return false, errors.New("guild does not exist")
	}
	_, hasMember := guild.members[userID]
	return hasMember, nil
}

func (d MockDB) GetMessageDate(messageID uint64) (time.Time, error) {
	panic("unimplemented")
}

func (d MockDB) GetMessages(guildID, channelID uint64) ([]queries.Message, error) {
	panic("unimplemented")
}

func (d MockDB) GetMessagesBefore(guildID, channelID uint64, date time.Time) ([]queries.Message, error) {
	panic("unimplemented")
}

func (d MockDB) UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) error {
	panic("unimplemented")
}

func (d MockDB) GetGuildPicture(guildID uint64) (string, error) {
	panic("unimplemented")
}

func (d MockDB) GetInvites(guildID uint64) ([]queries.Invite, error) {
	panic("unimplemented")
}

func (d MockDB) DeleteMember(guildID, userID uint64) error {
	if guild, ok := d.guilds[guildID]; !ok {
		return errors.New("guild not found")
	} else {
		delete(guild.members, userID)
		return nil
	}
}

func (d MockDB) GetLocalGuilds(userID uint64) ([]uint64, error) {
	panic("unimplemented")
}

func (d MockDB) ChannelsForGuild(guildID uint64) ([]queries.Channel, error) {
	guild, ok := d.guilds[guildID]
	if !ok {
		return nil, errors.New("guild does not exist")
	}
	channels := []queries.Channel{}
	for channelID := range guild.channels {
		c, ok := d.channels[channelID]
		if !ok {
			return nil, fmt.Errorf("channel %d does not exist on %d", channelID, guildID)
		}
		channels = append(channels, queries.Channel{
			ChannelID:   channelID,
			GuildID:     sql.NullInt64{Int64: int64(guildID)},
			ChannelName: c.name,
			Position:    c.position,
			Category:    c.category,
			Metadata:    []byte{},
		})
	}
	return channels, nil
}

func (d MockDB) MembersInGuild(guildID uint64) ([]uint64, error) {
	panic("unimplemented")
}

func (d MockDB) CountMembersInGuild(guildID uint64) (int64, error) {
	panic("unimplemented")
}

func (d MockDB) GetMessage(messageID uint64) (queries.Message, error) {
	panic("unimplemented")
}

func (d MockDB) GetUserByEmail(email string) (queries.GetUserByEmailRow, error) {
	user := d.userByEmail[email]
	if user == nil {
		return queries.GetUserByEmailRow{}, errors.New("user does not exist")
	}
	return queries.GetUserByEmailRow{
		UserID:   user.id,
		Username: user.username,
		Avatar:   sql.NullString{},
		Status:   0,
		Password: user.password,
	}, nil
}

func (d MockDB) GetUserByID(userID uint64) (queries.GetUserRow, error) {
	user := d.users[userID]
	if user == nil {
		return queries.GetUserRow{}, errors.New("user does not exist")
	}
	return queries.GetUserRow{
		UserID:   user.id,
		Username: user.username,
		Avatar:   sql.NullString{},
		IsBot:    false,
		Status:   0,
	}, nil
}

func (d MockDB) AddSession(userID uint64, session string) error {
	d.userBySession[session] = userID
	return nil
}

func (d MockDB) GetLocalUserForForeignUser(userID uint64, homeserver string) (uint64, error) {
	panic("unimplemented")
}

func (d MockDB) AddLocalUser(userID uint64, email, username string, passwordHash []byte) error {
	u := &User{
		email:    email,
		username: username,
		password: passwordHash,
		id:       userID,
	}
	d.users[userID] = u
	d.userByEmail[email] = u
	return nil
}

func (d MockDB) AddForeignUser(homeServer string, userID, localUserID uint64, username, avatar string) (uint64, error) {
	panic("unimplemented")
}

func (d MockDB) EmailExists(email string) (bool, error) {
	_, exists := d.userByEmail[email]
	return exists, nil
}

func (d MockDB) ExpireSessions() error {
	panic("unimplemented")
}

func (d MockDB) UpdateUsername(userID uint64, username string) error {
	panic("unimplemented")
}

func (d MockDB) GetAvatar(userID uint64) (sql.NullString, error) {
	panic("unimplemented")
}

func (d MockDB) UpdateAvatar(userID uint64, avatar string) error {
	panic("unimplemented")
}

func (d MockDB) HasGuildWithID(guildID uint64) (bool, error) {
	panic("unimplemented")
}

func (d MockDB) HasChannelWithID(guildID, channelID uint64) (bool, error) {
	panic("unimplemented")
}

func (d MockDB) HasMessageWithID(guildID, channelID, messageID uint64) (bool, error) {
	panic("unimplemented")
}

func (d MockDB) GetGuildByID(guildID uint64) (queries.Guild, error) {
	guild, exists := d.guilds[guildID]
	if !exists {
		return queries.Guild{}, errors.New("guild doesn't exist")
	}
	return queries.Guild{
		GuildID:    guild.id,
		PictureUrl: guild.picture,
		GuildName:  guild.name,
		Metadata:   []byte{},
	}, nil
}

func (d MockDB) UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string, metadata *harmonytypesv1.Metadata, updateMetadata bool) (time.Time, error) {
	panic("unimplemented")
}

func (d MockDB) SetStatus(userID uint64, status harmonytypesv1.UserStatus) error {
	panic("unimplemented")
}

func (d MockDB) SetUsername(userID uint64, username string) error {
	panic("unimplemented")
}

func (d MockDB) SetAvatar(userID uint64, avatar string) error {
	panic("unimplemented")
}

func (d MockDB) SetIsBot(userID uint64, isBot bool) error {
	panic("unimplemented")
}

func (d MockDB) GetUserMetadata(userID uint64, appID string) (string, error) {
	panic("unimplemented")
}

func (d MockDB) GetNonceInfo(nonce string) (queries.GetNonceInfoRow, error) {
	panic("unimplemented")
}

func (d MockDB) AddNonce(nonce string, userID uint64, homeServer string) error {
	panic("unimplemented")
}

func (d MockDB) GetGuildList(userID uint64) ([]queries.GetGuildListRow, error) {
	panic("unimplemented")
}

func (d MockDB) GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error) {
	panic("unimplemented")
}

func (d MockDB) AddGuildToList(userID, guildID uint64, homeServer string) error {
	panic("unimplemented")
}

func (d MockDB) MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error {
	panic("unimplemented")
}

func (d MockDB) GetChannelListPosition(guildID, channelID uint64) (string, error) {
	panic("unimplemented")
}

func (d MockDB) MoveChannel(guildID, channelID, previousID, nextID uint64) error {
	panic("unimplemented")
}

func (d MockDB) RemoveGuildFromList(userID, guildID uint64, homeServer string) error {
	// TODO: add guild list to mock DB
	return nil
}

func (d MockDB) UserIsLocal(userID uint64) error {
	panic("unimplemented")
}

func (d MockDB) CreateEmotePack(userID, packID uint64, packName string) error {
	panic("unimplemented")
}

func (d MockDB) IsPackOwner(userID, packID uint64) (bool, error) {
	panic("unimplemented")
}

func (d MockDB) AddEmoteToPack(packID uint64, imageID string, name string) error {
	panic("unimplemented")
}

func (d MockDB) DeleteEmoteFromPack(packID uint64, imageID string) error {
	panic("unimplemented")
}

func (d MockDB) DeleteEmotePack(packID uint64) error {
	panic("unimplemented")
}

func (d MockDB) GetEmotePacks(userID uint64) ([]queries.GetEmotePacksRow, error) {
	panic("unimplemented")
}

func (d MockDB) GetEmotePackEmotes(packID uint64) ([]queries.GetEmotePackEmotesRow, error) {
	panic("unimplemented")
}

func (d MockDB) DequipEmotePack(userID, packID uint64) error {
	panic("unimplemented")
}

func (d MockDB) AddRoleToGuild(guildID uint64, role *chatv1.Role) error {
	panic("unimplemented")
}

func (d MockDB) RemoveRoleFromGuild(guildID, roleID uint64) error {
	panic("unimplemented")
}

func (d MockDB) GetRolePositions(guildID, before, previous uint64) (pos string, retErr error) {
	panic("unimplemented")
}

func (d MockDB) MoveRole(guildID, roleID, beforeRole, previousRole uint64) error {
	panic("unimplemented")
}

func (d MockDB) GetGuildRoles(guildID uint64) ([]*chatv1.Role, error) {
	panic("unimplemented")
}

func (d MockDB) SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []types.PermissionsNode) error {
	panic("unimplemented")
}

func (d MockDB) GetPermissions(guildID uint64, channelID uint64, roleID uint64) (permissions []types.PermissionsNode, err error) {
	panic("unimplemented")
}

func (d MockDB) GetPermissionsData(guildID uint64) (types.PermissionsData, error) {
	panic("unimplemented")
}

func (d MockDB) RolesForUser(guildID, userID uint64) ([]uint64, error) {
	panic("unimplemented")
}

func (d MockDB) ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) error {
	panic("unimplemented")
}

func (d MockDB) ModifyRole(guildID, roleID uint64, name string, color int32, hoist, pingable, updateName, updateColor, updateHoist, updatePingable bool) error {
	panic("unimplemented")
}

func (d MockDB) DeleteFileMeta(fileID string) error {
	panic("unimplemented")
}

func (d MockDB) GetFileIDByHash(hash []byte) (string, error) {
	panic("unimplemented")
}

func (d MockDB) AddFileHash(fileID string, hash []byte) error {
	panic("unimplemented")
}

func (d MockDB) SetFileMetadata(fileID string, contentType, name string, size int32) error {
	panic("unimplemented")
}

func (d MockDB) GetFileMetadata(fileID string) (queries.GetFileMetadataRow, error) {
	panic("unimplemented")
}

func (d MockDB) GetFirstChannel(guildID uint64) (uint64, error) {
	if guild, ok := d.guilds[guildID]; !ok {
		return 0, errors.New("guild not found")
	} else {
		for channel := range guild.channels {
			return channel, nil
		}
	}
	return 0, errors.New("guild has no channels")
}

func (d MockDB) ExtendSession(session string) error {
	panic("unimplemented")
}

func (d MockDB) BanUser(guildID, userID uint64) error {
	if guild, ok := d.guilds[guildID]; !ok {
		return errors.New("guild not found")
	} else {
		guild.bans[userID] = struct{}{}
		return nil
	}
}

func (d MockDB) IsBanned(guildID, userID uint64) (bool, error) {
	guild, ok := d.guilds[guildID]
	if !ok {
		return false, errors.New("guild not found")
	}
	_, ok = guild.bans[userID]
	return ok, nil
}

func (d MockDB) UnbanUser(guildID, userID uint64) error {
	panic("unimplemented")
}
