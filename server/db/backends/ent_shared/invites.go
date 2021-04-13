package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/invite"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *DB) CreateInvite(guildID uint64, possibleUses int32, name string) (inv *types.InviteData, err error) {
	defer doRecovery(&err)

	data := d.Invite.
		Create().
		SetID(name).
		SetGuildID(guildID).
		SetPossibleUses(int64(possibleUses)).
		SaveX(ctx)

	inv = &types.InviteData{
		ID:           data.ID,
		PossibleUses: int32(data.PossibleUses),
		Uses:         int32(data.Uses),
	}

	return
}

func (d *DB) IncrementInvite(inviteID string) (err error) {
	defer doRecovery(&err)
	d.Invite.Update().Where(invite.ID(inviteID)).AddUses(1).ExecX(ctx)
	return
}

func (d *DB) DeleteInvite(inviteID string) (err error) {
	defer doRecovery(&err)
	d.Invite.Delete().Where(invite.ID(inviteID)).ExecX(ctx)
	return
}

func (d *DB) GetInvites(guildID uint64) (invites []*types.InviteData, err error) {
	defer doRecovery(&err)
	queriedInvites := d.Guild.Query().Where(guild.ID(guildID)).QueryInvite().AllX(ctx)
	for _, inv := range queriedInvites {
		invites = append(invites, &types.InviteData{
			ID:           inv.ID,
			PossibleUses: int32(inv.PossibleUses),
			Uses:         int32(inv.Uses),
		})
	}
	return
}

func (d *DB) ResolveGuildID(inviteID string) (guildID uint64, err error) {
	defer doRecovery(&err)
	guildID = d.Invite.GetX(ctx, inviteID).QueryGuild().OnlyX(ctx).ID
	return
}
