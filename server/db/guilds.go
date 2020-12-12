package db

import (
	"database/sql"

	"github.com/harmony-development/legato/server/db/queries"
)

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

// UserInGuild checks whether a user is in a guild
func (db *HarmonyDB) UserInGuild(userID, guildID uint64) (bool, error) {
	id, err := db.queries.UserInGuild(ctx, queries.UserInGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	db.Logger.CheckException(err)
	return id == userID, err
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

// MembersInGuild lists the members of a guild
func (db *HarmonyDB) MembersInGuild(guildID uint64) ([]uint64, error) {
	return db.queries.GetGuildMembers(ctx, guildID)
}

func (db *HarmonyDB) HasGuildWithID(guildID uint64) (bool, error) {
	count, err := db.queries.GuildWithIDExists(ctx, guildID)
	return count, err
}

func (db *HarmonyDB) GetGuildByID(guildID uint64) (queries.Guild, error) {
	return db.queries.GetGuildData(ctx, guildID)
}
