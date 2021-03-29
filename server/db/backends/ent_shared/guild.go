package ent_shared

import (
	"database/sql"

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
			d.Channel.Create().SaveX(ctx),
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

func (d *database) GetOwner(guildID uint64) (userID uint64, err error) {
	defer doRecovery(&err)
	userID = d.Guild.
		Query().
		Where(guild.ID(guildID)).
		OnlyX(ctx).Owner
	return
}

func (d *database) IsOwner(guildID, userID uint64) (isOwner bool, err error) {
	defer doRecovery(&err)

	if owner, err := d.GetOwner(guildID); err != nil {
		panic(err)
	} else {
		isOwner = userID == owner
	}
	return
}

func (d *database) CreateInvite(guildID uint64, possibleUses int32, name string) (inv queries.Invite, err error) {
	defer doRecovery(&err)

	savedInvite := d.Invite.
		Create().
		SetCode(name).
		SetGuildID(guildID).
		SetPossibleUses(int64(possibleUses)).
		SaveX(ctx)

	inv = queries.Invite{
		InviteID: savedInvite.Code,
		Uses:     int32(savedInvite.Uses),
		PossibleUses: sql.NullInt32{
			Int32: int32(savedInvite.PossibleUses),
			Valid: false,
		},
		GuildID: guildID,
	}

	return
}

func (d *database) AddMemberToGuild(userID, guildID uint64) (err error) {
	defer doRecovery(&err)

	d.Guild.UpdateOneID(guildID).AddUserIDs(userID).SaveX(ctx)

	return
}
