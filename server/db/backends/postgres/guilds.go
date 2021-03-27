// +build ignore

package postgres

import (
	"database/sql"
	"errors"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/db/utilities"
	"github.com/ztrue/tracerr"
)

// CreateGuild creates a standard guild
func (db *database) CreateGuild(owner, id, channelID uint64, guildName, picture string) (*queries.Guild, error) {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	guild, err := tq.CreateGuild(ctx, queries.CreateGuildParams{
		GuildID:    id,
		OwnerID:    owner,
		GuildName:  guildName,
		PictureUrl: picture,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	err = tq.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  owner,
		GuildID: guild.GuildID,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	_, err = tq.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     utilities.ToSqlInt64(guild.GuildID),
		ChannelID:   channelID,
		ChannelName: "general",
		Position:    "",
	})
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	return &guild, nil
}

// DeleteGuild deletes a guild with an ID
func (db *database) DeleteGuild(guildID uint64) error {
	err := db.queries.DeleteGuild(ctx, guildID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

// GetOwner gets the owner of a guild
func (db *database) GetOwner(guildID uint64) (uint64, error) {
	owner, err := db.queries.GetGuildOwner(ctx, guildID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return owner, err
}

// IsOwner returns whether the user is the guild owner
func (db *database) IsOwner(guildID, userID uint64) (bool, error) {
	owner, err := db.GetOwner(guildID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	if err != nil {
		return false, err
	}
	return owner == userID, nil
}

// AddInvite inserts a new invite to the DB
func (db *database) CreateInvite(guildID uint64, possibleUses int32, name string) (queries.Invite, error) {
	inv, err := db.queries.CreateGuildInvite(ctx, queries.CreateGuildInviteParams{
		InviteID:     name,
		PossibleUses: sql.NullInt32{Int32: possibleUses, Valid: true},
		GuildID:      guildID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return inv, err
}

// AddMemberToGuild adds a new member to a guild
func (db *database) AddMemberToGuild(userID, guildID uint64) error {
	err := db.queries.AddUserToGuild(ctx, queries.AddUserToGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

// InviteToGuild
func (db *database) ResolveGuildID(inviteID string) (uint64, error) {
	id, err := db.queries.ResolveGuildID(ctx, inviteID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return id, err
}

// IncrementInvite adds to the invite counter in a DB
func (db *database) IncrementInvite(inviteID string) error {
	err := db.queries.IncrementInvite(ctx, inviteID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

// DeleteInvite removes an invite from the DB
func (db *database) DeleteInvite(inviteID string) error {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	tq := db.queries.WithTx(tx)
	rows, err := tq.DeleteInvite(ctx, inviteID)
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	if rows > 1 {
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

// UserInGuild checks whether a user is in a guild
func (db *database) UserInGuild(userID, guildID uint64) (bool, error) {
	id, err := db.queries.UserInGuild(ctx, queries.UserInGuildParams{
		UserID:  userID,
		GuildID: guildID,
	})
	err = tracerr.Wrap(err)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	db.Logger.CheckException(err)
	return id == userID, err
}

// UpdateGuildName updates the guild name
func (db *database) UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) error {
	e := utilities.Executor{}
	if updateName {
		e.Execute(func() error {
			return db.queries.SetGuildName(ctx, queries.SetGuildNameParams{
				GuildName: name,
				GuildID:   guildID,
			})
		})
	}
	if updatePicture {
		e.Execute(func() error {
			return db.queries.SetGuildPicture(ctx, queries.SetGuildPictureParams{
				PictureUrl: picture,
				GuildID:    guildID,
			})
		})
	}
	if updateMetadata {
		e.Execute(func() error {
			data, err := utilities.SerializeMetadata(metadata)
			if err != nil {
				return err
			}
			return db.queries.SetGuildMetadata(ctx, queries.SetGuildMetadataParams{
				Metadata: data,
				GuildID:  guildID,
			})
		})
	}

	return e.Err
}

// GetGuildPicture gets the picture for a given guild
func (db *database) GetGuildPicture(guildID uint64) (string, error) {
	pic, err := db.queries.GetGuildPicture(ctx, guildID)
	if err != nil {
		err = tracerr.Wrap(err)
		return "", err
	}
	return pic, err
}

func (db *database) CountMembersInGuild(guildID uint64) (int64, error) {
	return db.queries.CountMembersInGuild(ctx, guildID)
}

// SetGuildPicture sets the picture for a given guild
func (db *database) SetGuildPicture(guildID uint64, pictureURL string) error {
	err := db.queries.SetGuildPicture(ctx, queries.SetGuildPictureParams{
		GuildID:    guildID,
		PictureUrl: pictureURL,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

// GetInvites gets open invites for a guild
func (db *database) GetInvites(guildID uint64) ([]queries.Invite, error) {
	invites, err := db.queries.OpenInvites(ctx, guildID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return invites, err
}

// DeleteMember deletes a member from a guild
func (db *database) DeleteMember(guildID, userID uint64) error {
	err := db.queries.RemoveUserFromGuild(ctx, queries.RemoveUserFromGuildParams{
		GuildID: guildID,
		UserID:  userID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

// MembersInGuild lists the members of a guild
func (db *database) MembersInGuild(guildID uint64) ([]uint64, error) {
	data, err := db.queries.GetGuildMembers(ctx, guildID)
	err = tracerr.Wrap(err)
	return data, err
}

func (db *database) HasGuildWithID(guildID uint64) (bool, error) {
	count, err := db.queries.GuildWithIDExists(ctx, guildID)
	err = tracerr.Wrap(err)
	return count, err
}

func (db *database) GetGuildByID(guildID uint64) (queries.Guild, error) {
	data, err := db.queries.GetGuildData(ctx, guildID)
	err = tracerr.Wrap(err)
	return data, err
}

func (db *database) GetFirstChannel(guildID uint64) (uint64, error) {
	data, err := db.queries.GetFirstChannel(ctx, utilities.ToSqlInt64(guildID))
	return data, err
}

func (db *database) BanUser(guildID, userID uint64) error {
	return db.queries.BanUser(ctx, queries.BanUserParams{
		GuildID: guildID,
		UserID:  userID,
	})
}

func (db *database) UnbanUser(guildID, userID uint64) error {
	return db.queries.UnbanUser(ctx, queries.UnbanUserParams{
		GuildID: guildID,
		UserID:  userID,
	})
}

func (db *database) IsBanned(guildID, userID uint64) (bool, error) {
	return db.queries.IsBanned(ctx, queries.IsBannedParams{
		GuildID: guildID,
		UserID:  userID,
	})
}
