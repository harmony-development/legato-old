package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
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
func (db *HarmonyDB) CreateGuild(owner, id, channelID uint64, guildName, picture string) (*queries.Guild, error) {
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
		ChannelID:   channelID,
		ChannelName: "general",
		Position:    "",
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
func (db *HarmonyDB) IsOwner(guildID, userID uint64) (bool, error) {
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
func (db *HarmonyDB) AddMemberToGuild(userID, guildID uint64) error {
	err := db.queries.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	db.Logger.CheckException(err)
	return err
}

// AddChannelToGuild adds a new channel to a guild
func (db *HarmonyDB) AddChannelToGuild(guildID uint64, channelName string, before, previous uint64, category bool) (queries.Channel, error) {
	pos, err := db.GetChannelPositions(guildID, before, previous)
	if err != nil {
		return queries.Channel{}, err
	}
	chanID, err := db.Sonyflake.NextID()
	if err != nil {
		return queries.Channel{}, err
	}
	channel, err := db.queries.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guildID),
		ChannelID:   chanID,
		ChannelName: channelName,
		Position:    pos,
		Category:    category,
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
func (db *HarmonyDB) AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64) (*queries.Message, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	msg, err := tq.AddMessage(ctx, queries.AddMessageParams{
		GuildID:     guildID,
		ChannelID:   channelID,
		UserID:      userID,
		MessageID:   messageID,
		Content:     message,
		Embeds:      embeds,
		Actions:     actions,
		Overrides:   overrides,
		Attachments: attachments,
		ReplyToID:   replyTo,
	})
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID, channelID, guildID uint64) error {
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
func (db *HarmonyDB) UserInGuild(userID, guildID uint64) (bool, error) {
	id, err := db.queries.UserInGuild(ctx, queries.UserInGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	db.Logger.CheckException(err)
	return id == userID, err
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID uint64) (time.Time, error) {
	msgDate, err := db.queries.GetMessageDate(ctx, messageID)
	db.Logger.CheckException(err)
	return msgDate, err
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID, channelID uint64) ([]queries.Message, error) {
	msgs, err := db.GetMessagesBefore(guildID, channelID, time.Now())
	db.Logger.CheckException(err)
	return msgs, err
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID, channelID uint64, date time.Time) ([]queries.Message, error) {
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

// GetLocalGuilds gets the guilds a user is in
func (db *HarmonyDB) GetLocalGuilds(userID uint64) ([]uint64, error) {
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

func (db *HarmonyDB) UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte) (time.Time, error) {
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
			data, err := tq.UpdateMessageEmbeds(ctx, queries.UpdateMessageEmbedsParams{
				MessageID: messageID,
				Embeds:    *embeds,
			})
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if actions != nil {
		e.Execute(func() error {
			data, err := tq.UpdateMessageActions(ctx, queries.UpdateMessageActionsParams{
				MessageID: messageID,
				Actions:   *actions,
			})
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if overrides != nil {
		e.Execute(func() error {
			return tq.UpdateMessageOverrides(ctx, queries.UpdateMessageOverridesParams{
				MessageID: messageID,
				Overrides: *overrides,
			})
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

func (db *HarmonyDB) GetChannelListPosition(guildID, channelID uint64) (string, error) {
	position, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
	db.Logger.CheckException(err)
	return position, err
}

func (db *HarmonyDB) GetChannelPositions(guildID, before, previous uint64) (pos string, retErr error) {
	nextPos, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		ChannelID: before,
		GuildID:   toSqlInt64(guildID),
	})
	if err != nil && err != sql.ErrNoRows {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	prevPos, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		ChannelID: previous,
		GuildID:   toSqlInt64(guildID),
	})
	if err != nil && err != sql.ErrNoRows {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	pos = Rank(prevPos, nextPos)
	return
}

func (db *HarmonyDB) MoveChannel(guildID, channelID, previousID, nextID uint64) error {
	pos, err := db.GetChannelPositions(guildID, previousID, nextID)
	if err != nil {
		return err
	}
	err = db.queries.MoveChannel(ctx, queries.MoveChannelParams{
		Position:  pos,
		ChannelID: channelID,
		GuildID:   toSqlInt64(guildID),
	})

	db.Logger.CheckException(err)

	return err
}

func (db *HarmonyDB) MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error {
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

func (db HarmonyDB) HasMessageWithID(guildID, channelID, messageID uint64) (bool, error) {
	return db.queries.MessageWithIDExists(ctx, queries.MessageWithIDExistsParams{
		GuildID:   guildID,
		ChannelID: channelID,
		MessageID: messageID,
	})
}

func (db HarmonyDB) UserIsLocal(userID uint64) error {
	ok, err := db.queries.UserIsLocal(ctx, userID)
	if err == nil && !ok {
		err = ErrNotLocal
	}
	return err
}

func (db HarmonyDB) CreateEmotePack(userID, packID uint64, packName string) error {
	err := db.queries.CreateEmotePack(ctx, queries.CreateEmotePackParams{
		UserID:   userID,
		PackID:   packID,
		PackName: packName,
	})
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) IsPackOwner(userID, packID uint64) (bool, error) {
	owner, err := db.queries.GetPackOwner(ctx, packID)
	if err != nil {
		return false, err
	}
	return owner == userID, nil
}

func (db HarmonyDB) AddEmoteToPack(packID uint64, imageID string, name string) error {
	err := db.queries.AddEmoteToPack(ctx, queries.AddEmoteToPackParams{
		PackID:    packID,
		ImageID:   imageID,
		EmoteName: name,
	})
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) DeleteEmoteFromPack(packID uint64, imageID string) error {
	err := db.queries.DeleteEmoteFromPack(ctx, queries.DeleteEmoteFromPackParams{
		PackID:  packID,
		ImageID: imageID,
	})
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) DeleteEmotePack(packID uint64) error {
	err := db.queries.DeleteEmotePack(ctx, queries.DeleteEmotePackParams{
		PackID: packID,
	})
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) GetEmotePacks(userID uint64) ([]queries.GetEmotePacksRow, error) {
	emotes, err := db.queries.GetEmotePacks(ctx, userID)
	if err != nil {
		db.Logger.CheckException(err)
		return nil, err
	}
	return emotes, nil
}

func (db HarmonyDB) GetEmotePackEmotes(packID uint64) ([]queries.GetEmotePackEmotesRow, error) {
	return db.queries.GetEmotePackEmotes(ctx, packID)
}

func (db HarmonyDB) DequipEmotePack(userID, packID uint64) error {
	return db.queries.DequipEmotePack(ctx, queries.DequipEmotePackParams{
		PackID: packID,
		UserID: userID,
	})
}
