package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
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
	tq := db.Queries.WithTx(tx)
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
	return db.Queries.DeleteGuild(ctx, guildID)
}

// GetOwner gets the owner of a guild
func (db *HarmonyDB) GetOwner(guildID int64) (int64, error) {
	return db.Queries.GetGuildOwner(ctx, guildID)
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
	return db.Queries.CreateGuildInvite(ctx, queries.CreateGuildInviteParams{
		Name:         name,
		PossibleUses: sql.NullInt32{Int32: possibleUses, Valid: true},
		GuildID:      guildID,
	})
}

// AddMemberToGuild adds a new member to a guild
func (db *HarmonyDB) AddMemberToGuild(userID int64, guildID int64) error {
	return db.Queries.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
}

// AddChannelToGuild adds a new channel to a guild
func (db *HarmonyDB) AddChannelToGuild(guildID int64, channelName string) (queries.Channel, error) {
	return db.Queries.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guildID),
		ChannelName: channelName,
	})
}

// DeleteChannelFromGuild removes a channel from a guild
func (db *HarmonyDB) DeleteChannelFromGuild(guildID, channelID int64) error {
	return db.Queries.DeleteChannel(ctx, queries.DeleteChannelParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
}

// AddMessage adds a message to a channel
func (db *HarmonyDB) AddMessage(messageID, channelID, guildID, userID int64, message string, attachments []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	tq := db.Queries.WithTx(tx)
	if _, err := tq.AddMessage(ctx, queries.AddMessageParams{
		MessageID: messageID,
		GuildID:   guildID,
		ChannelID: channelID,
		UserID:    userID,
		CreatedAt: time.Now(),
		EditedAt:  sql.NullTime{},
		Content:   message,
	}); err != nil {
		return err
	}
	for _, attachment := range attachments {
		if err := tq.AddAttachment(ctx, queries.AddAttachmentParams{
			MessageID:     messageID,
			AttachmentUrl: attachment,
		}); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID int64, channelID int64, guildID int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	tq := db.Queries.WithTx(tx)
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
	return db.Queries.GetMessageAuthor(ctx, messageID)
}

// InviteToGuild
func (db *HarmonyDB) ResolveGuildID(inviteID int64) (int64, error) {
	return db.Queries.ResolveGuildID(ctx, inviteID)
}

// IncrementInvite adds to the invite counter in a DB
func (db *HarmonyDB) IncrementInvite(inviteID int64) error {
	return db.Queries.IncrementInvite(ctx, inviteID)
}

// DeleteInvite removes an invite from the DB
func (db *HarmonyDB) DeleteInvite(inviteID int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	tq := db.Queries.WithTx(tx)
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
		return db.Queries.SessionToUserID(ctx, session)
	}
	return s, nil
}

// UserInGuild checks whether a user is in a guild
func (db *HarmonyDB) UserInGuild(userID int64, guildID int64) error {
	_, err := db.Queries.UserInGuild(ctx, queries.UserInGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	return err
}

// GetAttachments gets attachments for a message
func (db *HarmonyDB) GetAttachments(messageID int64) ([]string, error) {
	return db.Queries.GetAttachments(ctx, messageID)
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID int64) (time.Time, error) {
	return db.Queries.GetMessageDate(ctx, messageID)
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID string, channelID string) ([]Message, error) {
	return db.GetMessagesBefore(guildID, channelID, 0)
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID int64, channelID int64, date time.Time) ([]Message, error) {
	return db.Queries.GetMessages(ctx, queries.GetMessagesParams{
		Guildid:   guildID,
		Channelid: channelID,
		Before:    date,
		Max:       int32(db.Config.Server.GetMessageCount),
	})
}

// UpdateGuildName updates the guild name
func (db *HarmonyDB) UpdateGuildName(guildID string, newName string) error {
	_, err := db.Exec("UPDATE guilds SET guildname=$1 WHERE guildid=$2", newName, guildID)
	return err
}

// GetGuildPicture gets the picture for a given guild
func (db *HarmonyDB) GetGuildPicture(guildID string) (*string, error) {
	var picture string
	err := db.QueryRow("SELECT picture FROM guilds WHERE guildid=$1", guildID).Scan(&picture)
	return &picture, err
}

// SetGuildPicture sets the picture for a given guild
func (db *HarmonyDB) SetGuildPicture(guildID string, pictureID string) error {
	_, err := db.Exec("UPDATE guilds SET picture=$1 WHERE guildid=$2", pictureID, guildID)
	return err
}

// AddAttachments adds attachments to a message
func (db *HarmonyDB) AddAttachments(messageID string, attachments []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, a := range attachments {
		if _, err := tx.Exec("INSERT INTO attachments(messageid, attachment) VALUES($1, $2)", messageID, a); err != nil {
			return err
		}
	}
	return nil
}

// GetInvites gets open invites for a guild
func (db *HarmonyDB) GetInvites(guildID string) ([]Invite, error) {
	res, err := db.Query("SElECT inviteid, invitecount FROM invites WHERE guildid=$1 ORDER BY invitecount", guildID)
	if err != nil {
		return nil, err
	}
	var invites []Invite
	for res.Next() {
		var invite Invite
		if err := res.Scan(&invite.ID, &invite.Uses); err != nil {
			return nil, err
		}
		invites = append(invites, invite)
	}
	return invites, nil
}

// DeleteMember deletes a member from a guild
func (db *HarmonyDB) DeleteMember(guildID string, userID string) error {
	_, err := db.Exec("DELETE FROM guildmembers WHERE guildid=$1 AND userid=$2", guildID, userID)
	return err
}
