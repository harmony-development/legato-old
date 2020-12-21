package db

import (
	"database/sql"

	"github.com/harmony-development/legato/server/db/queries"
	"github.com/ztrue/tracerr"
)

// AddChannelToGuild adds a new channel to a guild
func (db *HarmonyDB) AddChannelToGuild(guildID uint64, channelName string, before, previous uint64, category bool, kind string) (queries.Channel, error) {
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
	channel, err := db.queries.CreateChannel(ctx, queries.CreateChannelParams{
		GuildID:     toSqlInt64(guildID),
		ChannelID:   chanID,
		ChannelName: channelName,
		Position:    pos,
		Category:    category,
		Kind: sql.NullString{
			String: kind,
			Valid:  kind != "",
		},
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return channel, err
}

// DeleteChannelFromGuild removes a channel from a guild
func (db *HarmonyDB) DeleteChannelFromGuild(guildID, channelID uint64) error {
	err := db.queries.DeleteChannel(ctx, queries.DeleteChannelParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
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

func (db *HarmonyDB) HasChannelWithID(guildID, channelID uint64) (bool, error) {
	count, err := db.queries.NumChannelsWithID(ctx, queries.NumChannelsWithIDParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
	err = tracerr.Wrap(err)
	return count != 0, err
}

func (db *HarmonyDB) GetChannelListPosition(guildID, channelID uint64) (string, error) {
	position, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		GuildID:   toSqlInt64(guildID),
		ChannelID: channelID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return position, err
}

func (db *HarmonyDB) GetChannelPositions(guildID, before, previous uint64) (pos string, retErr error) {
	nextPos, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		ChannelID: before,
		GuildID:   toSqlInt64(guildID),
	})
	err = tracerr.Wrap(err)
	if err != nil && err != sql.ErrNoRows {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	prevPos, err := db.queries.GetChannelPosition(ctx, queries.GetChannelPositionParams{
		ChannelID: previous,
		GuildID:   toSqlInt64(guildID),
	})
	err = tracerr.Wrap(err)
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
		err = tracerr.Wrap(err)
		return err
	}
	err = db.queries.MoveChannel(ctx, queries.MoveChannelParams{
		Position:  pos,
		ChannelID: channelID,
		GuildID:   toSqlInt64(guildID),
	})
	err = tracerr.Wrap(err)

	db.Logger.CheckException(err)

	return err
}
