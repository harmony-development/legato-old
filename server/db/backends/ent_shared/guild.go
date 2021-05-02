package ent_shared

import (
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *DB) CreateGuild(owner, id, channelID uint64, guildName, picture string) (guild *types.GuildData, err error) {
	defer doRecovery(&err)
	guild = &types.GuildData{}
	data := d.Guild.Create().
		SetID(id).
		SetOwnerID(owner).
		SetName(guildName).
		SetPicture(picture).
		SetMetadata(&harmonytypesv1.Metadata{}).
		SaveX(ctx)

	channel := d.Channel.
		Create().
		SetKind(0).
		SetID(channelID).
		SetName("general").
		SetPosition("").
		SetMetadata(&harmonytypesv1.Metadata{}).
		SaveX(ctx)

	data.Update().AddUserIDs(owner).AddChannel(channel).ExecX(ctx)

	guild.ID = data.ID
	guild.Name = data.Name
	guild.Owner = owner
	guild.Picture = data.Picture
	return
}

func (d *DB) DeleteGuild(guildID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.
		Delete().
		Where(guild.ID(guildID)).
		ExecX(ctx)
	return
}

func (d *DB) BanUser(guildID, userID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.UpdateOneID(guildID).AddBanIDs(userID).ExecX(ctx)
	return
}

func (d *DB) UnbanUser(guildID, userID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.UpdateOneID(guildID).RemoveBanIDs(userID).ExecX(ctx)
	return
}

func (d *DB) IsBanned(guildID, userID uint64) (banned bool, err error) {
	defer doRecovery(&err)
	banned = d.Guild.Query().Where(guild.ID(guildID)).QueryBans().Where(user.ID(userID)).ExistX(ctx)
	return
}

func (d *DB) GetGuildByID(guildID uint64) (g *types.GuildData, err error) {
	defer doRecovery(&err)

	data := d.Guild.Query().WithOwner().Where(guild.ID(guildID)).OnlyX(ctx)

	g = &types.GuildData{
		ID:      data.ID,
		Owner:   data.Edges.Owner.ID,
		Name:    data.Name,
		Picture: data.Picture,
	}
	return
}

func (d *DB) GetGuildPicture(guildID uint64) (picture string, err error) {
	defer doRecovery(&err)
	picture = d.Guild.GetX(ctx, guildID).Picture
	return
}

func (d *DB) GetLocalGuilds(userID uint64) (guilds []uint64, err error) {
	defer doRecovery(&err)
	guilds = d.User.GetX(ctx, userID).QueryGuild().IDsX(ctx)
	return
}

func (d *DB) HasGuildWithID(guildID uint64) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.Guild.Query().Where(guild.ID(guildID)).ExistX(ctx)
	return
}

func (d *DB) UserInGuild(userID, guildID uint64) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.Guild.
		Query().
		Where(
			guild.ID(guildID),
		).
		QueryUser().Where(
		user.ID(userID),
	).ExistX(ctx)
	return
}

func (d *DB) UpdateGuildInformation(guildID uint64, name, picture string, metadata *harmonytypesv1.Metadata, updateName, updatePicture, updateMetadata bool) (err error) {
	defer doRecovery(&err)
	update := d.Guild.
		UpdateOneID(guildID)
	if updateName {
		update.SetName(name)
	}
	if updatePicture {
		update.SetPicture(picture)
	}
	if updateMetadata {
		update.SetMetadata(metadata)
	}
	update.ExecX(ctx)
	return
}
