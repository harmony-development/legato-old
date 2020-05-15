package db

//go:generate sqlc generate

import (
	"database/sql"
	"time"

	"context"
)

func toSqlString(input string) sql.NullString {
	return sql.NullString{String: input, Valid: true}
}

func toSqlInt64(input int64) sql.NullInt64 {
	return sql.NullInt64{Int64: input, Valid: true}
}

var ctx = context.Background()

// CreateGuild creates a standard guild
func (db *HarmonyDB) CreateGuild(owner string, guildName string, picture string) (*Guild, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	queries := db.Queries.WithTx(tx)
	guild, err := queries.CreateGuild(ctx, CreateGuildParams{
		OwnerID:    owner,
		GuildName:  guildName,
		PictureUrl: picture,
	})
	if err != nil {
		return nil, err
	}
	err = queries.AddUserToGuild(ctx, AddUserToGuildParams{
		UserID:  owner,
		GuildID: guild.GuildID,
	})
	if err != nil {
		return nil, err
	}
	_, err = queries.CreateChannel(ctx, CreateChannelParams{
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
func (db *HarmonyDB) GetOwner(guildID int64) (string, error) {
	return db.Queries.GetGuildOwner(ctx, guildID)
}

// IsOwner returns whether the user is the guild owner
func (db *HarmonyDB) IsOwner(guildID int64, userID string) (bool, error) {
	str, err := db.GetOwner(guildID)
	if err != nil {
		return false, err
	}
	return str == userID, nil
}

// AddInvite inserts a new invite to the DB
func (db *HarmonyDB) CreateInvite(guildID int64, possibleUses int32, name string) (Invite, error) {
	return db.Queries.CreateGuildInvite(ctx, CreateGuildInviteParams{
		Name:         name,
		PossibleUses: sql.NullInt32{Int32: possibleUses, Valid: true},
		GuildID:      toSqlInt64(guildID),
	})
}

// AddMemberToGuild adds a new member to a guild
func (db *HarmonyDB) AddMemberToGuild(userID string, guildID int64) error {
	return db.Queries.AddUserToGuild(ctx, AddUserToGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
}

// AddChannelToGuild adds a new channel to a guild
func (db *HarmonyDB) AddChannelToGuild(guildID int64, channelName string) (Channel, error) {
	return db.Queries.CreateChannel(ctx, CreateChannelParams{
		GuildID:     toSqlInt64(guildID),
		ChannelName: channelName,
	})
}

// DeleteChannelFromGuild removes a channel from a guild
func (db *HarmonyDB) DeleteChannelFromGuild(guildID, channelID int64) error {
	return db.Queries.DeleteChannel(ctx, DeleteChannelParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
}

// AddMessage adds a message to a channel
func (db *HarmonyDB) AddMessage(messageID string, channelID string, guildID string, userID string, message string, attachments []string) error {
	if _, err := db.Exec(`
	INSERT INTO messages(messageid, channelid, guildid, userid, createdat, message) 
	VALUES($1, $2, $3, $4, $5, $6)
	`, messageID, channelID, guildID, userID, time.Now().UTC().Unix(), message); err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, attachment := range attachments {
		if _, err := tx.Exec(`
			INSERT INTO attachments(messageid, attachment)
			VALUES($1, $2)
		`, messageID, attachment); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID string, channelID string, guildID string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(`
		DELETE FROM messages WHERE messageid=$1 AND channelid=$2 AND guildID=$3
	`, messageID, channelID, guildID)
	if err != nil {
		return err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCount != 1 {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return nil
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// GetMessageOwner gets the owner of a messageID
func (db *HarmonyDB) GetMessageOwner(messageID string) (*string, error) {
	var userID string
	err := db.QueryRow("SELECT userid FROM messages WHERE messageid=$1", messageID).Scan(&userID)
	return &userID, err
}

// ResolveInvite gets a guildID from an invite code
func (db *HarmonyDB) ResolveInvite(inviteID string) (*string, error) {
	var guildID string
	if err := db.QueryRow(`SELECT guildid FROM invites WHERE inviteid=$1`, inviteID).Scan(&guildID); err != nil {
		return nil, err
	}
	return &guildID, nil
}

// IncrementInvite adds to the invite use counter
func (db *HarmonyDB) IncrementInvite(inviteID string) error {
	_, err := db.Exec(`UPDATE invites SET uses=uses+1 WHERE inviteid=$1`, inviteID)
	return err
}

// DeleteInvite deletes an invite
func (db *HarmonyDB) DeleteInvite(inviteID string, guildID string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(`
		DELETE FROM invites WHERE inviteid=$1 AND guildid=$2
	`, inviteID, guildID)
	if err != nil {
		return err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCount != 1 {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return nil
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// SessionToUserID gets the user ID from a session
func (db *HarmonyDB) SessionToUserID(session string) (*string, error) {
	userID, exists := db.SessionCache.Get(session)
	s := userID.(string)
	if !exists {
		if err := db.QueryRow("SELECT userid FROM sessions WHERE session=$1", session).Scan(&s); err != nil {
			return nil, err
		}
		return &s, nil
	}
	return &s, nil
}

// UserInGuild checks whether a user is in a guild
func (db *HarmonyDB) UserInGuild(userID string, guildID string) (bool, error) {
	ok, err := db.ContainsRow("SELECT userid FROM guildmembers WHERE userid=$1 AND guildid=$2", userID, guildID)
	if !ok || err != nil {
		return false, err
	}
	return ok, nil
}

// GetAttachments gets attachments for a message
func (db *HarmonyDB) GetAttachments(messageID string) ([]string, error) {
	res, err := db.Query("SELECT attachment FROM attachments WHERE messageid=$1", messageID)
	if err != nil {
		return nil, err
	}
	var attachments []string
	for res.Next() {
		var a string
		if err := res.Scan(&a); err != nil {
			return nil, err
		}
		attachments = append(attachments, a)
	}
	return attachments, nil
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID string) (*int, error) {
	var createdAt int
	if err := db.QueryRow("SELECT createdat FROM messages WHERE messageid=$1", messageID).Scan(&createdAt); err != nil {
		return nil, err
	}
	return &createdAt, nil
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID string, channelID string) ([]Message, error) {
	return db.GetMessagesBefore(guildID, channelID, 0)
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID string, channelID string, date int) ([]Message, error) {
	res, err := db.Query(
		`SELECT messageid, userid, message, createdat FROM messages 
		WHERE guildid=$1 AND 
		      channelid=$2 AND 
		      createdat<$3 
		ORDER BY createdat DESC LIMIT $4`,
		guildID, channelID, date, db.Config.Server.GetMessageCount,
	)
	if err != nil {
		return nil, err
	}
	var messages []Message
	for res.Next() {
		var messageID, userID, message string
		var createdAt int
		if err := res.Scan(&messageID, &userID, &message, &createdAt); err != nil {
			return nil, err
		}
		attachments, err := db.GetAttachments(messageID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, Message{
			UserID:      userID,
			MessageID:   messageID,
			Message:     message,
			Attachments: attachments,
			CreatedAt:   createdAt,
		})
	}
	return messages, nil
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
