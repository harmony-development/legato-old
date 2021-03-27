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

// AddChannelToGuild adds a new channel to a guild
func (db *database) AddChannelToGuild(guildID uint64, channelName string, before, previous uint64, category bool, metadata *harmonytypesv1.Metadata) (queries.Channel, error) {
	pos, err := db.GetChannelPositions(guildID, before, previous)
	err = tracerr.Wrap(err)
	if err != nil {
		return queries.Channel{}, err
	}
	chanID, err := db.Sonyflake.NextID()
	err = tracerr.Wrap(err)
	if err != nil {
		return queries.Channel{}, err
	}
	md, err := utilities.SerializeMetadata(metadata)
	if err != nil {
		return queries.Channel{}, err
	}
	channel, err := db.queries.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     utilities.ToSqlInt64(guildID),
		ChannelID:   chanID,
		ChannelName: channelName,
		Position:    pos,
		Category:    category,
		Metadata:    md,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return channel, err
}

// DeleteChannelFromGuild removes a channel from a guild
func (db *database) DeleteChannelFromGuild(guildID, channelID uint64) error {
	err := db.queries.DeleteChannel(ctx, queries.DeleteChannelParams{
		GuildID:   utilities.ToSqlInt64(guildID),
		ChannelID: channelID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db *database) UpdateChannelInformation(guildID, channelID uint64, name string, updateName bool, metadata *harmonytypesv1.Metadata, updateMetadata bool) error {
	tx, err := db.Begin()
	if err != nil {
		return tracerr.Wrap(err)
	}
	tq := db.queries.WithTx(tx)

	e := utilities.Executor{}
	if updateName {
		e.Execute(func() error {
			return tq.UpdateChannelName(ctx, queries.UpdateChannelNameParams{
				ChannelName: name,
				GuildID:     utilities.ToSqlInt64(guildID),
				ChannelID:   channelID,
			})
		})
	}
	if updateMetadata {
		e.Execute(func() error {
			data, err := utilities.SerializeMetadata(metadata)
			if err != nil {
				return err
			}
			return tq.UpdateChannelMetadata(ctx, queries.UpdateChannelMetadataParams{
				Metadata:  data,
				GuildID:   utilities.ToSqlInt64(guildID),
				ChannelID: channelID,
			})
		})
	}

	return e.Err
}

// UpdateChannelName sets the name of a channel
func (db *database) SetChannelName(guildID, channelID uint64, name string) error {
	return db.queries.UpdateChannelName(ctx, queries.UpdateChannelNameParams{
		ChannelName: name,
		GuildID:     utilities.ToSqlInt64(guildID),
		ChannelID:   channelID,
	})
}

// ChannelsForGuild gets the channels for a guild
func (db *database) ChannelsForGuild(guildID uint64) ([]queries.Channel, error) {
	return db.queries.GetChannels(ctx, utilities.ToSqlInt64(guildID))
}

func (db *database) HasChannelWithID(guildID, channelID uint64) (bool, error) {
	count, err := db.queries.NumChannelsWithID(ctx, queries.NumChannelsWithIDParams{
		GuildID:   utilities.ToSqlInt64(guildID),
		ChannelID: channelID,
	})
	err = tracerr.Wrap(err)
	return count != 0, err
}

func (db *database) GetChannelListPosition(guildID, channelID uint64) (string, error) {
	position, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		GuildID:   utilities.ToSqlInt64(guildID),
		ChannelID: channelID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return position, err
}

func (db *database) GetChannelPositions(guildID, before, previous uint64) (pos string, retErr error) {
	nextPos, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		ChannelID: before,
		GuildID:   utilities.ToSqlInt64(guildID),
	})
	err = tracerr.Wrap(err)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	prevPos, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		ChannelID: previous,
		GuildID:   utilities.ToSqlInt64(guildID),
	})
	err = tracerr.Wrap(err)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	pos = Rank(prevPos, nextPos)
	return
}

func (db *database) MoveChannel(guildID, channelID, previousID, nextID uint64) error {
	pos, err := db.GetChannelPositions(guildID, previousID, nextID)
	if err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	err = db.queries.MoveChannel(ctx, queries.MoveChannelParams{
		Position:  pos,
		ChannelID: channelID,
		GuildID:   utilities.ToSqlInt64(guildID),
	})
	err = tracerr.Wrap(err)

	db.Logger.CheckException(err)

	return err
}
