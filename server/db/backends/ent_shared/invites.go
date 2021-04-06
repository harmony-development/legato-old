package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/invite"
)

func (d *database) CreateInvite(guildID uint64, possibleUses int32, name string) (inv *entgen.Invite, err error) {
	defer doRecovery(&err)

	inv = d.Invite.
		Create().
		SetCode(name).
		SetGuildID(guildID).
		SetPossibleUses(int64(possibleUses)).
		SaveX(ctx)

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

func (d *database) GetInvites(guildID uint64) (invites []*entgen.Invite, err error) {
	defer doRecovery(&err)
	queriedInvites := d.Guild.Query().Where(guild.ID(guildID)).QueryInvite().AllX(ctx)
	for _, inv := range queriedInvites {
		invites = append(invites, inv)
	}
	return
}
