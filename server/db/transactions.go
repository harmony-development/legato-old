package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/db/queries"
)

func toSqlString(input string) sql.NullString {
	return sql.NullString{String: input, Valid: true}
}

func toSqlInt64(input uint64) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(input), Valid: true}
}

type executor struct {
	err error
}

func (e *executor) Execute(f func() error) {
	if e.err != nil {
		return
	}
	e.err = f()
}

var ctx = context.Background()

// CreateGuild creates a standard guild
func (db *HarmonyDB) CreateGuild(owner, id uint64, guildName string, picture string) (*queries.Guild, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	guild, err := tq.CreateGuild(ctx, queries.CreateGuildParams{
		GuildID:    id,
		OwnerID:    owner,
		GuildName:  guildName,
		PictureUrl: picture,
	})
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	err = tq.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  owner,
		GuildID: guild.GuildID,
	})
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	_, err = tq.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guild.GuildID),
		ChannelName: "general",
	})
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return nil, err
	}
	return &guild, nil
}

// DeleteGuild deletes a guild with an ID
func (db *HarmonyDB) DeleteGuild(guildID uint64) error {
	err := db.queries.DeleteGuild(ctx, guildID)
	db.Logger.CheckException(err)
	return err
}

// GetOwner gets the owner of a guild
func (db *HarmonyDB) GetOwner(guildID uint64) (uint64, error) {
	owner, err := db.queries.GetGuildOwner(ctx, guildID)
	db.Logger.CheckException(err)
	return owner, err
}

// IsOwner returns whether the user is the guild owner
func (db *HarmonyDB) IsOwner(guildID uint64, userID uint64) (bool, error) {
	owner, err := db.GetOwner(guildID)
	db.Logger.CheckException(err)
	if err != nil {
		return false, err
	}
	return owner == userID, nil
}

// AddInvite inserts a new invite to the DB
func (db *HarmonyDB) CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error) {
	inv, err := db.queries.CreateGuildInvite(ctx, queries.CreateGuildInviteParams{
		InviteID:     name,
		PossibleUses: sql.NullInt32{Int32: possibleUses, Valid: true},
		GuildID:      guildID,
	})
	db.Logger.CheckException(err)
	return inv, err
}

// AddMemberToGuild adds a new member to a guild
func (db *HarmonyDB) AddMemberToGuild(userID uint64, guildID uint64) error {
	err := db.queries.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	db.Logger.CheckException(err)
	return err
}

// AddChannelToGuild adds a new channel to a guild
func (db *HarmonyDB) AddChannelToGuild(guildID uint64, channelName string) (queries.Channel, error) {
	channel, err := db.queries.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guildID),
		ChannelName: channelName,
	})
	db.Logger.CheckException(err)
	return channel, err
}

// DeleteChannelFromGuild removes a channel from a guild
func (db *HarmonyDB) DeleteChannelFromGuild(guildID, channelID uint64) error {
	err := db.queries.DeleteChannel(ctx, queries.DeleteChannelParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
	db.Logger.CheckException(err)
	return err
}

// AddMessage adds a message to a channel
func (db *HarmonyDB) AddMessage(channelID, guildID, userID uint64, message string, attachments []string, embeds, actions [][]byte) (*queries.Message, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
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
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	for _, attachment := range attachments {
		if err := tq.AddAttachment(ctx, queries.AddAttachmentParams{
			MessageID:  msg.MessageID,
			Attachment: attachment,
		}); err != nil {
			_ = tx.Rollback()
			db.Logger.CheckException(err)
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID uint64, channelID uint64, guildID uint64) error {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	tq := db.queries.WithTx(tx)
	numRows, err := tq.DeleteMessage(ctx, queries.DeleteMessageParams{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	})
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	if numRows > 1 { // JUST IN CASE the delete query deletes too much
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

// GetMessageOwner gets the owner of a messageID
func (db *HarmonyDB) GetMessageOwner(messageID uint64) (uint64, error) {
	owner, err := db.queries.GetMessageAuthor(ctx, messageID)
	db.Logger.CheckException(err)
	return owner, err
}

// InviteToGuild
func (db *HarmonyDB) ResolveGuildID(inviteID string) (uint64, error) {
	id, err := db.queries.ResolveGuildID(ctx, inviteID)
	db.Logger.CheckException(err)
	return id, err
}

// IncrementInvite adds to the invite counter in a DB
func (db *HarmonyDB) IncrementInvite(inviteID string) error {
	err := db.queries.IncrementInvite(ctx, inviteID)
	db.Logger.CheckException(err)
	return err
}

// DeleteInvite removes an invite from the DB
func (db *HarmonyDB) DeleteInvite(inviteID string) error {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	tq := db.queries.WithTx(tx)
	rows, err := tq.DeleteInvite(ctx, inviteID)
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	if rows > 1 {
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

// SessionToUserID gets the user ID from a session
func (db *HarmonyDB) SessionToUserID(session string) (uint64, error) {
	userID, exists := db.SessionCache.Get(session)
	s, ok := userID.(uint64)
	if !exists || !ok {
		userID, err := db.queries.SessionToUserID(ctx, session)
		if err != nil {
			db.Logger.CheckException(err)
		}
		return userID, err
	}
	return s, nil
}

// UserInGuild checks whether a user is in a guild
func (db *HarmonyDB) UserInGuild(userID uint64, guildID uint64) (bool, error) {
	id, err := db.queries.UserInGuild(ctx, queries.UserInGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	db.Logger.CheckException(err)
	return id == userID, err
}

// GetAttachments gets attachments for a message
func (db *HarmonyDB) GetAttachments(messageID uint64) ([]string, error) {
	attachments, err := db.queries.GetAttachments(ctx, messageID)
	db.Logger.CheckException(err)
	return attachments, err
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID uint64) (time.Time, error) {
	msgDate, err := db.queries.GetMessageDate(ctx, messageID)
	db.Logger.CheckException(err)
	return msgDate, err
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID uint64, channelID uint64) ([]queries.Message, error) {
	msgs, err := db.GetMessagesBefore(guildID, channelID, time.Now())
	db.Logger.CheckException(err)
	return msgs, err
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID uint64, channelID uint64, date time.Time) ([]queries.Message, error) {
	msgsBefore, err := db.queries.GetMessages(ctx, queries.GetMessagesParams{
		Guildid:   guildID,
		Channelid: channelID,
		Before:    date,
		Max:       int32(db.Config.Server.GetMessageCount),
	})
	db.Logger.CheckException(err)
	return msgsBefore, err
}

// UpdateGuildName updates the guild name
func (db *HarmonyDB) UpdateGuildName(guildID uint64, newName string) error {
	err := db.queries.SetGuildName(ctx, queries.SetGuildNameParams{
		GuildName: newName,
		GuildID:   guildID,
	})
	db.Logger.CheckException(err)
	return err
}

// GetGuildPicture gets the picture for a given guild
func (db *HarmonyDB) GetGuildPicture(guildID uint64) (string, error) {
	pic, err := db.queries.GetGuildPicture(ctx, guildID)
	if err != nil {
		return "", err
	}
	return pic, err
}

// SetGuildPicture sets the picture for a given guild
func (db *HarmonyDB) SetGuildPicture(guildID uint64, pictureURL string) error {
	err := db.queries.SetGuildPicture(ctx, queries.SetGuildPictureParams{
		GuildID:    guildID,
		PictureUrl: pictureURL,
	})
	db.Logger.CheckException(err)
	return err
}

// AddAttachments adds attachments to a message
func (db *HarmonyDB) AddAttachments(messageID uint64, attachments []string) error {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	withTX := db.queries.WithTx(tx)
	for _, a := range attachments {
		err = withTX.AddAttachment(ctx, queries.AddAttachmentParams{
			MessageID:  messageID,
			Attachment: a,
		})
		if err != nil {
			_ = tx.Rollback()
			db.Logger.CheckException(err)
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

// GetInvites gets open invites for a guild
func (db *HarmonyDB) GetInvites(guildID uint64) ([]queries.Invite, error) {
	invites, err := db.queries.OpenInvites(ctx, guildID)
	db.Logger.CheckException(err)
	return invites, err
}

// DeleteMember deletes a member from a guild
func (db *HarmonyDB) DeleteMember(guildID, userID uint64) error {
	err := db.queries.RemoveUserFromGuild(ctx, queries.RemoveUserFromGuildParams{
		GuildID: guildID,
		UserID:  userID,
	})
	db.Logger.CheckException(err)
	return err
}

// GuildsForUser gets the guilds a user is in
func (db *HarmonyDB) GuildsForUser(userID uint64) ([]uint64, error) {
	return db.queries.GuildsForUser(ctx, userID)
}

// UpdateChannelName sets the name of a channel
func (db *HarmonyDB) SetChannelName(guildID, channelID uint64, name string) error {
	return db.queries.UpdateChannelName(ctx, queries.UpdateChannelNameParams{
		ChannelName: name,
		GuildID:     toSqlInt64(guildID),
		ChannelID:   channelID,
	})
}

// GuildsForUser gets the guilds a user is in with additional data
func (db *HarmonyDB) GuildsForUserWithData(userID uint64) (ret []queries.Guild, err error) {
	var data []queries.GuildsForUserWithDataRow
	data, err = db.queries.GuildsForUserWithData(ctx, userID)
	for _, guild := range data {
		ret = append(ret, queries.Guild{
			GuildID:    guild.GuildID,
			OwnerID:    guild.OwnerID,
			GuildName:  guild.GuildName,
			PictureUrl: guild.PictureUrl,
		})
	}
	return
}

// ChannelsForGuild gets the channels for a guild
func (db *HarmonyDB) ChannelsForGuild(guildID uint64) ([]queries.Channel, error) {
	return db.queries.GetChannels(ctx, toSqlInt64(guildID))
}

// MembersInGuild lists the members of a guild
func (db *HarmonyDB) MembersInGuild(guildID uint64) ([]uint64, error) {
	return db.queries.GetGuildMembers(ctx, guildID)
}

// GetMessage gets the data of a message
func (db *HarmonyDB) GetMessage(messageID uint64) (queries.Message, error) {
	return db.queries.GetMessage(ctx, messageID)
}

// GetUser gets a user with their email
func (db *HarmonyDB) GetUserByEmail(email string) (queries.GetUserByEmailRow, error) {
	return db.queries.GetUserByEmail(ctx, email)
}

// GetUserByID gets a user with their ID and their home server
func (db *HarmonyDB) GetUserByID(userID uint64) (queries.GetUserRow, error) {
	return db.queries.GetUser(ctx, userID)
}

// AddSession persists a session into the DB
func (db *HarmonyDB) AddSession(userID uint64, session string) error {
	db.SessionCache.Add(session, userID)
	return db.queries.AddSession(ctx, queries.AddSessionParams{
		UserID:     userID,
		Session:    session,
		Expiration: time.Now().UTC().Add(db.Config.Server.SessionDuration).Unix(),
	})
}

// GetLocalUserForForeignUser gets a local user from the foreign users database
func (db *HarmonyDB) GetLocalUserForForeignUser(userID uint64, homeserver string) (uint64, error) {
	return db.queries.GetLocalUserID(ctx, queries.GetLocalUserIDParams{
		UserID:     userID,
		HomeServer: homeserver,
	})
}

// AddLocalUser adds a user to the DB that contains login information (not foreign)
func (db *HarmonyDB) AddLocalUser(userID uint64, email, username string, passwordHash []byte) error {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	tq := db.queries.WithTx(tx)
	if err := tq.AddUser(ctx, userID); err != nil {
		return err
	}
	if err := tq.AddLocalUser(ctx, queries.AddLocalUserParams{
		UserID:   userID,
		Email:    email,
		Password: passwordHash,
	}); err != nil {
		return err
	}
	if err := tq.AddProfile(ctx, queries.AddProfileParams{
		UserID:   userID,
		Username: username,
		Avatar:   sql.NullString{},
		Status:   int16(profilev1.UserStatus_USER_STATUS_OFFLINE),
	}); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return tx.Rollback()
	}
	return nil
}

// AddForeignUser inserts
func (db *HarmonyDB) AddForeignUser(homeServer string, userID, localUserID uint64, username, avatar string) (uint64, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return 0, err
	}
	tq := db.queries.WithTx(tx)
	if err := tq.AddUser(ctx, localUserID); err != nil {
		return 0, err
	}
	if err := tq.AddProfile(ctx, queries.AddProfileParams{
		UserID:   localUserID,
		Username: username,
		Avatar:   toSqlString(avatar),
		Status:   int16(profilev1.UserStatus_USER_STATUS_OFFLINE),
	}); err != nil {
		return 0, err
	}
	if userID, err = tq.AddForeignUser(ctx, queries.AddForeignUserParams{
		UserID:      userID,
		HomeServer:  homeServer,
		LocalUserID: localUserID,
	}); err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}
	return userID, nil
}

func (db *HarmonyDB) EmailExists(email string) (bool, error) {
	count, err := db.queries.EmailExists(ctx, email)
	return count > 0, err
}

func (db *HarmonyDB) ExpireSessions() error {
	if err := db.queries.ExpireSessions(ctx, time.Now().UTC().Unix()); err != nil {
		return err
	}
	return nil
}

func (db *HarmonyDB) UpdateUsername(userID uint64, username string) error {
	return db.queries.UpdateUsername(ctx, queries.UpdateUsernameParams{
		Username: username,
		UserID:   userID,
	})
}

func (db *HarmonyDB) GetAvatar(userID uint64) (sql.NullString, error) {
	return db.queries.GetAvatar(ctx, userID)
}

func (db *HarmonyDB) UpdateAvatar(userID uint64, avatar string) error {
	return db.queries.UpdateAvatar(ctx, queries.UpdateAvatarParams{
		Avatar: toSqlString(avatar),
		UserID: userID,
	})
}

func (db *HarmonyDB) HasGuildWithID(guildID uint64) (bool, error) {
	count, err := db.queries.GuildWithIDExists(ctx, guildID)
	return count, err
}

func (db *HarmonyDB) HasChannelWithID(guildID, channelID uint64) (bool, error) {
	count, err := db.queries.NumChannelsWithID(ctx, queries.NumChannelsWithIDParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
	return count != 0, err
}

func (db *HarmonyDB) AddFileHash(fileID string, hash []byte) error {
	return db.queries.AddFileHash(ctx, queries.AddFileHashParams{
		Hash:   hash,
		FileID: fileID,
	})
}

func (db *HarmonyDB) GetFileIDFromHash(hash []byte) (string, error) {
	return db.queries.GetFileByHash(ctx, hash)
}

func (db *HarmonyDB) GetGuildByID(guildID uint64) (queries.Guild, error) {
	return db.queries.GetGuildData(ctx, guildID)
}

func (db *HarmonyDB) UpdateMessage(messageID uint64, content *string, embeds, actions *[][]byte) (time.Time, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return time.Time{}, err
	}
	tq := db.queries.WithTx(tx)
	var editedAt time.Time
	e := executor{}
	if content != nil {
		e.Execute(func() error {
			data, err := tq.UpdateMessageContent(ctx, queries.UpdateMessageContentParams{
				MessageID: messageID,
				Content:   *content,
			})
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if embeds != nil {
		e.Execute(func() error {
			var rawEmbeds []json.RawMessage
			for _, embed := range *embeds {
				rawEmbeds = append(rawEmbeds, embed)
			}
			data, err := tq.UpdateMessageEmbeds(ctx, queries.UpdateMessageEmbedsParams{
				MessageID: messageID,
				Embeds:    rawEmbeds,
			})
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if actions != nil {
		e.Execute(func() error {
			var rawActions []json.RawMessage
			for _, action := range *actions {
				rawActions = append(rawActions, action)
			}
			data, err := tq.UpdateMessageActions(ctx, queries.UpdateMessageActionsParams{
				MessageID: messageID,
				Actions:   rawActions,
			})
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if e.err != nil {
		if err := tx.Rollback(); err != nil {
			return time.Time{}, err
		}
		return time.Time{}, e.err
	}
	if err := tx.Commit(); err != nil {
		return time.Time{}, err
	}
	return editedAt, nil
}

func (db *HarmonyDB) SetStatus(userID uint64, status profilev1.UserStatus) error {
	return db.queries.SetStatus(ctx, queries.SetStatusParams{
		Status: int16(status), // lol shut up it's an int16
		UserID: userID,
	})
}

func (db *HarmonyDB) GetUserMetadata(userID uint64, appID string) (string, error) {
	metadata, err := db.queries.GetUserMetadata(ctx, queries.GetUserMetadataParams{
		UserID: userID,
		AppID:  appID,
	})
	db.Logger.CheckException(err)
	return metadata, err
}

func (db *HarmonyDB) GetNonceInfo(nonce string) (queries.GetNonceInfoRow, error) {
	info, err := db.queries.GetNonceInfo(ctx, nonce)
	db.Logger.CheckException(err)
	return info, err
}

func (db *HarmonyDB) AddNonce(nonce string, userID uint64, homeServer string) error {
	err := db.queries.AddNonce(ctx, queries.AddNonceParams{
		Nonce:      nonce,
		UserID:     userID,
		HomeServer: homeServer,
	})
	db.Logger.CheckException(err)
	return err
}

func (db *HarmonyDB) GetGuildList(userID uint64) ([]queries.GetGuildListRow, error) {
	guilds, err := db.queries.GetGuildList(ctx, userID)
	db.Logger.CheckException(err)
	return guilds, err
}

func (db *HarmonyDB) GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error) {
	position, err := db.queries.GetGuildListPosition(ctx, queries.GetGuildListPositionParams{
		UserID:     userID,
		GuildID:    guildID,
		HomeServer: homeServer,
	})
	db.Logger.CheckException(err)
	return position, err
}

func (db *HarmonyDB) AddGuildToList(userID, guildID uint64, homeServer string) error {
	pos, err := db.queries.GetLastGuildPositionInList(ctx, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			pos = ""
		} else {
			db.Logger.Exception(err)
			return err
		}
	}

	err = db.queries.AddToGuildList(ctx, queries.AddToGuildListParams{
		UserID:     userID,
		GuildID:    guildID,
		HomeServer: homeServer,
		Position:   Rank(pos, ""),
	})

	db.Logger.CheckException(err)
	return err
}

func (db *HarmonyDB) MoveGuild(userID uint64, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error {
	nextPos, err := db.queries.GetGuildListPosition(ctx, queries.GetGuildListPositionParams{
		UserID:     userID,
		GuildID:    nextGuildID,
		HomeServer: nextHomeServer,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			nextPos = ""
		} else {
			db.Logger.Exception(err)
			return err
		}
	}

	prevPos, err := db.queries.GetGuildListPosition(ctx, queries.GetGuildListPositionParams{
		UserID:     userID,
		GuildID:    prevGuildID,
		HomeServer: prevHomeServer,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			nextPos = ""
		} else {
			db.Logger.Exception(err)
			return err
		}
	}

	err = db.queries.MoveGuild(ctx, queries.MoveGuildParams{
		Position:   Rank(prevPos, nextPos),
		GuildID:    guildID,
		HomeServer: homeServer,
	})

	db.Logger.CheckException(err)

	return err
}

func (db HarmonyDB) RemoveGuildFromList(userID, guildID uint64, homeServer string) error {
	err := db.queries.RemoveGuildFromList(ctx, queries.RemoveGuildFromListParams{
		UserID:     userID,
		GuildID:    guildID,
		HomeServer: homeServer,
	})
	db.Logger.CheckException(err)
	return err
}
