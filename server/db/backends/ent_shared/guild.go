package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/queries"
)

func (d *database) CreateGuild(owner, id, channelID uint64, guildName, picture string) (guild *queries.Guild, err error) {
	defer doRecovery(&err)
	g := d.Guild.Create().
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
	guild = &queries.Guild{
		GuildID:    g.ID,
		OwnerID:    g.Owner,
		GuildName:  g.Name,
		PictureUrl: g.Picture,
		Metadata:   g.Metadata,
	}
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
