package ent_shared

import (
	"database/sql"

	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/invite"
	"github.com/harmony-development/legato/server/db/queries"
)

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

func (d *database) IncrementInvite(inviteID string) (err error) {
	defer doRecovery(&err)
	d.Invite.Update().Where(invite.Code(inviteID)).AddUses(1).ExecX(ctx)
	return
}

func (d *database) DeleteInvite(inviteID string) (err error) {
	defer doRecovery(&err)
	d.Invite.Delete().Where(invite.Code(inviteID)).ExecX(ctx)
	return
}

func (d *database) GetInvites(guildID uint64) (invites []queries.Invite, err error) {
	defer doRecovery(&err)
	queriedInvites := d.Guild.Query().Where(guild.ID(guildID)).QueryInvite().AllX(ctx)
	for _, inv := range queriedInvites {
		invites = append(invites, queries.Invite{
			InviteID: inv.Code,
			Uses: int32(inv.Uses),
			PossibleUses: sql.NullInt32{
				Int32: int32(inv.PossibleUses),
			},
			GuildID: guildID,
		})
	}
	return
}