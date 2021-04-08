package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/invite"
)

func (d *DB) CreateInvite(guildID uint64, possibleUses int32, name string) (inv *entgen.Invite, err error) {
	defer doRecovery(&err)

	inv = d.Invite.
		Create().
		SetID(name).
		SetGuildID(guildID).
		SetPossibleUses(int64(possibleUses)).
		SaveX(ctx)

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

func (d *DB) GetInvites(guildID uint64) (invites []*entgen.Invite, err error) {
	defer doRecovery(&err)
	queriedInvites := d.Guild.Query().Where(guild.ID(guildID)).QueryInvite().AllX(ctx)
	for _, inv := range queriedInvites {
		invites = append(invites, inv)
	}
	return
}

func (d *DB) ResolveGuildID(inviteID string) (guildID uint64, err error) {
	defer doRecovery(&err)
	guildID = d.Invite.GetX(ctx, inviteID).QueryGuild().OnlyX(ctx).ID
	return
}
