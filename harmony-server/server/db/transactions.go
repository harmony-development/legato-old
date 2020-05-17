package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"harmony-server/server/db/queries"
)

func toSqlString(input string) sql.NullString {
	return sql.NullString{String: input, Valid: true}
}

func toSqlInt64(input int64) sql.NullInt64 {
	return sql.NullInt64{Int64: input, Valid: true}
}

var ctx = context.Background()

// CreateGuild creates a standard guild
func (db *HarmonyDB) CreateGuild(owner int64, guildName string, picture string) (*queries.Guild, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	guild, err := tq.CreateGuild(ctx, queries.CreateGuildParams{
		OwnerID:    owner,
		GuildName:  guildName,
		PictureUrl: picture,
	})
	if err != nil {
		return nil, err
	}
	err = tq.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  owner,
		GuildID: guild.GuildID,
	})
	if err != nil {
		return nil, err
	}
	_, err = tq.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guild.GuildID),
		ChannelName: "general",
	})
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &guild, nil
}

// DeleteGuild deletes a guild with an ID
func (db *HarmonyDB) DeleteGuild(guildID int64) error {
	return db.queries.DeleteGuild(ctx, guildID)
}

// GetOwner gets the owner of a guild
func (db *HarmonyDB) GetOwner(guildID int64) (int64, error) {
	return db.queries.GetGuildOwner(ctx, guildID)
}

// IsOwner returns whether the user is the guild owner
func (db *HarmonyDB) IsOwner(guildID int64, userID int64) (bool, error) {
	str, err := db.GetOwner(guildID)
	if err != nil {
		return false, err
	}
	return str == userID, nil
}

// AddInvite inserts a new invite to the DB
func (db *HarmonyDB) CreateInvite(guildID int64, possibleUses int32, name string) (queries.Invite, error) {
	return db.queries.CreateGuildInvite(ctx, queries.CreateGuildInviteParams{
		Name:         name,
		PossibleUses: sql.NullInt32{Int32: possibleUses, Valid: true},
		GuildID:      guildID,
	})
}

// AddMemberToGuild adds a new member to a guild
func (db *HarmonyDB) AddMemberToGuild(userID int64, guildID int64) error {
	return db.queries.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
}

// AddChannelToGuild adds a new channel to a guild
func (db *HarmonyDB) AddChannelToGuild(guildID int64, channelName string) (queries.Channel, error) {
	return db.queries.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guildID),
		ChannelName: channelName,
	})
}

// DeleteChannelFromGuild removes a channel from a guild
func (db *HarmonyDB) DeleteChannelFromGuild(guildID, channelID int64) error {
	return db.queries.DeleteChannel(ctx, queries.DeleteChannelParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
}

// AddMessage adds a message to a channel
func (db *HarmonyDB) AddMessage(channelID, guildID, userID int64, message string, attachments []string, embeds, actions [][]byte) (*queries.Message, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	var rawEmbeds, rawActions []json.RawMessage
	for _, embed := range embeds {
		rawEmbeds = append(rawEmbeds, json.RawMessage(embed))
	}
	for _, action := range actions {
		rawActions = append(rawActions, json.RawMessage(action))
	}
	msg, err := tq.AddMessage(ctx, queries.AddMessageParams{
		GuildID:   guildID,
		ChannelID: channelID,
		UserID:    userID,
		Content:   message,
		Embeds:    rawEmbeds,
		Actions:   rawActions,
	})
	if err != nil {
		return nil, err
	}
	for _, attachment := range attachments {
		if err := tq.AddAttachment(ctx, queries.AddAttachmentParams{
			MessageID:     msg.MessageID,
			AttachmentUrl: attachment,
		}); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID int64, channelID int64, guildID int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	tq := db.queries.WithTx(tx)
	numRows, err := tq.DeleteMessage(ctx, queries.DeleteMessageParams{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	})
	if err != nil {
		return err
	}
	if numRows > 1 { // JUST IN CASE the delete query deletes too much
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// GetMessageOwner gets the owner of a messageID
func (db *HarmonyDB) GetMessageOwner(messageID int64) (int64, error) {
	return db.queries.GetMessageAuthor(ctx, messageID)
}

// InviteToGuild
func (db *HarmonyDB) ResolveGuildID(inviteID int64) (int64, error) {
	return db.queries.ResolveGuildID(ctx, inviteID)
}

// IncrementInvite adds to the invite counter in a DB
func (db *HarmonyDB) IncrementInvite(inviteID int64) error {
	return db.queries.IncrementInvite(ctx, inviteID)
}

// DeleteInvite removes an invite from the DB
func (db *HarmonyDB) DeleteInvite(inviteID int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	tq := db.queries.WithTx(tx)
	rows, err := tq.DeleteInvite(ctx, inviteID)
	if err != nil {
		return err
	}
	if rows > 1 {
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// SessionToUserID gets the user ID from a session
func (db *HarmonyDB) SessionToUserID(session string) (int64, error) {
	userID, exists := db.SessionCache.Get(session)
	s, ok := userID.(int64)
	if !exists || !ok {
		return db.queries.SessionToUserID(ctx, session)
	}
	return s, nil
}

// UserInGuild checks whether a user is in a guild
func (db *HarmonyDB) UserInGuild(userID int64, guildID int64) (bool, error) {
	id, err := db.queries.UserInGuild(ctx, queries.UserInGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	return id == userID, err
}

// GetAttachments gets attachments for a message
func (db *HarmonyDB) GetAttachments(messageID int64) ([]string, error) {
	return db.queries.GetAttachments(ctx, messageID)
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID int64) (time.Time, error) {
	return db.queries.GetMessageDate(ctx, messageID)
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID int64, channelID int64) ([]queries.GetMessagesRow, error) {
	return db.GetMessagesBefore(guildID, channelID, time.Now())
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID int64, channelID int64, date time.Time) ([]queries.GetMessagesRow, error) {
	return db.queries.GetMessages(ctx, queries.GetMessagesParams{
		Guildid:   guildID,
		Channelid: channelID,
		Before:    date,
		Max:       int32(db.Config.Server.GetMessageCount),
	})
}

// UpdateGuildName updates the guild name
func (db *HarmonyDB) UpdateGuildName(guildID int64, newName string) error {
	return db.queries.SetGuildName(ctx, queries.SetGuildNameParams{
		GuildName: newName,
		GuildID:   guildID,
	})
}

// GetGuildPicture gets the picture for a given guild
func (db *HarmonyDB) GetGuildPicture(guildID int64) (string, error) {
	return db.queries.GetGuildPicture(ctx, guildID)
}

// SetGuildPicture sets the picture for a given guild
func (db *HarmonyDB) SetGuildPicture(guildID int64, pictureURL string) error {
	return db.queries.SetGuildPicture(ctx, queries.SetGuildPictureParams{
		GuildID:    guildID,
		PictureUrl: pictureURL,
	})
}

// AddAttachments adds attachments to a message
func (db *HarmonyDB) AddAttachments(messageID int64, attachments []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	withTX := db.queries.WithTx(tx)
	for _, a := range attachments {
		err = withTX.AddAttachment(ctx, queries.AddAttachmentParams{
			MessageID:     messageID,
			AttachmentUrl: a,
		})
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()
	return nil
}

// GetInvites gets open invites for a guild
func (db *HarmonyDB) GetInvites(guildID int64) ([]queries.Invite, error) {
	return db.queries.OpenInvites(ctx, guildID)
}

// DeleteMember deletes a member from a guild
func (db *HarmonyDB) DeleteMember(guildID, userID int64) error {
	return db.queries.RemoveUserFromGuild(ctx, queries.RemoveUserFromGuildParams{
		GuildID: guildID,
		UserID:  userID,
	})
}

// GuildsForUser gets the guilds a user is in
func (db *HarmonyDB) GuildsForUser(userID int64) ([]int64, error) {
	return db.queries.GuildsForUser(ctx, userID)
}

// ChannelsForGuild gets the channels for a guild
func (db *HarmonyDB) ChannelsForGuild(guildID int64) ([]queries.Channel, error) {
	return db.queries.GetChannels(ctx, toSqlInt64(guildID))
}

// MembersInGuild lists the members of a guild
func (db *HarmonyDB) MembersInGuild(guildID int64) ([]queries.GuildMember, error) {
	return db.queries.GetGuildMembers(ctx, guildID)
}

// GetMessage gets the data of a message
func (db *HarmonyDB) GetMessage(messageID int64) (queries.Message, error) {
	return db.queries.GetMessage(ctx, messageID)

// GetUser gets a user with their email
func (db *HarmonyDB) GetUser(email string) (queries.User, error) {
	return db.queries.GetUser(ctx, email)
}

func (db *HarmonyDB) AddSession(userID int64, session string) error {
	db.SessionCache.Add(session, userID)
	return db.queries.AddSession(ctx, queries.AddSessionParams{
		UserID:  userID,
		Session: session,
	})
}
