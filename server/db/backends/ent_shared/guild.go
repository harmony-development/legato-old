package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
)

func (d *database) CreateGuild(owner, id, channelID uint64, guildName, picture string) (guild *entgen.Guild, err error) {
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

func (d *database) DeleteGuild(guildID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.
		Delete().
		Where(guild.ID(guildID)).
		ExecX(ctx)
	return
}

func (d *database) BanUser(guildID, userID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.UpdateOneID(guildID).AddBanIDs(userID).ExecX(ctx)
	return
}

func (d *database) UnbanUser(guildID, userID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.UpdateOneID(guildID).RemoveBanIDs(userID).ExecX(ctx)
	return
}

func (d *database) GetGuildByID(guildID uint64) (guild *entgen.Guild, err error) {
	defer doRecovery(&err)
	guild = d.Guild.GetX(ctx, guildID)
	return
}

func (d *database) GetGuildPicture(guildID uint64) (picture string, err error) {
	defer doRecovery(&err)
	picture = d.Guild.GetX(ctx, guildID).Picture
	return
}

func (d *database) GetLocalGuilds(userID uint64) (guilds []uint64, err error) {
	defer doRecovery(&err)
	guilds = d.User.GetX(ctx, userID).QueryGuild().IDsX(ctx)
	return
}
