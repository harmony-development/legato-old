package ent_shared

import (
	"github.com/golang/protobuf/proto"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

func (d *DB) CreateGuild(owner, id, channelID uint64, guildName, picture string) (guild *entgen.Guild, err error) {
	defer doRecovery(&err)
	guild = d.Guild.Create().
		SetID(id).
		SetOwner(owner).
		SetName(guildName).
		SetPicture(picture).
		AddChannel(
			d.Channel.
				Create().
				SetKind(0).
				SetID(channelID).
				SetName("general").
				SaveX(ctx),
		).
		SaveX(ctx)
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

func (d *DB) GetGuildByID(guildID uint64) (guild *entgen.Guild, err error) {
	defer doRecovery(&err)
	guild = d.Guild.GetX(ctx, guildID)
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
		marshalled, err := proto.Marshal(metadata)
		if err != nil {
			panic(err)
		}
		update.SetMetadata(marshalled)
	}
	update.ExecX(ctx)
	return
}
