package ent_shared

import (
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/lexorank"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *DB) AddChannelToGuild(guildID, channelID uint64, channelName string, previous, next *uint64, kind types.ChannelKind, md *harmonytypesv1.Metadata) (c types.ChannelData, err error) {
	defer doRecovery(&err)

	previousChannelPos := ""
	nextChannelPos := ""

	if previous != nil && *previous != 0 {
		previousChannelPos = d.Channel.GetX(ctx, *previous).Position
	}
	if next != nil && *next != 0 {
		nextChannelPos = d.Channel.GetX(ctx, *next).Position
	}

	d.Channel.Create().
		SetID(channelID).
		SetGuildID(guildID).
		SetPosition(lexorank.Rank(previousChannelPos, nextChannelPos)).
		SetName(channelName).
		SetKind(uint64(kind)).
		SetMetadata(md).
		SaveX(ctx)

	c.ID = channelID
	c.Metadata = md
	c.Name = channelName

	return
}

func (d *DB) DeleteChannelFromGuild(guildID, channelID uint64) (err error) {
	defer doRecovery(&err)

	d.Guild.UpdateOneID(guildID).RemoveChannelIDs(channelID).ExecX(ctx)

	return
}

func (d *DB) ChannelsForGuild(guildID uint64) (chans []*types.ChannelData, err error) {
	defer doRecovery(&err)

	channels := d.Channel.Query().Where(channel.HasGuildWith(guild.ID(guildID))).AllX(ctx)

	for _, c := range channels {
		chans = append(chans, &types.ChannelData{
			ID:       c.ID,
			Name:     c.Name,
			Metadata: c.Metadata,
			Position: c.Position,
			Kind:     c.Kind,
		})
	}

	return
}

func (d *DB) HasChannelWithID(guildID, channelID uint64) (hasChannel bool, err error) {
	defer doRecovery(&err)
	hasChannel = d.Channel.Query().Where(channel.ID(channelID)).ExistX(ctx)
	return
}

func (d *DB) GetChannelListPosition(channelID uint64) (pos string, err error) {
	defer doRecovery(&err)
	d.Channel.GetX(ctx, channelID)
	return
}

func (d *DB) MoveChannel(channelID uint64, previousID, nextID *uint64) (err error) {
	defer doRecovery(&err)
	previousChannelPos := ""
	nextChannelPos := ""
	if previousID != nil {
		previousChannelPos = d.Channel.GetX(ctx, *previousID).Position
	}
	if nextID != nil {
		nextChannelPos = d.Channel.GetX(ctx, *nextID).Position
	}
	d.Channel.UpdateOneID(channelID).SetPosition(lexorank.Rank(previousChannelPos, nextChannelPos))
	return
}

func (d *DB) GetFirstChannel(guildID uint64) (channelID uint64, err error) {
	defer doRecovery(&err)
	channelID = d.Guild.GetX(ctx, guildID).QueryChannel().FirstIDX(ctx)
	return
}

func (d *DB) UpdateChannelInformation(guildID, channelID uint64, name *string, metadata *harmonytypesv1.Metadata) (err error) {
	defer doRecovery(&err)
	update := d.Channel.UpdateOneID(channelID)
	if name != nil {
		update.SetName(*name)
	}
	if metadata != nil {
		update.SetMetadata(metadata)
	}
	return
}
