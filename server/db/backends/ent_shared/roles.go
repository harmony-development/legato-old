package ent_shared

import "github.com/harmony-development/legato/server/db/ent/entgen"

func (d *database) AddRoleToGuild(guildID, roleID uint64, name string, color int, hoist, pingable bool) (err error) {
	doRecovery(&err)
	d.Guild.
		UpdateOneID(guildID).
		AddRole(
			d.Role.
				Create().
				SetName(name).
				SetColor(color).
				SetHoist(hoist).
				SetPingable(pingable).
				SaveX(ctx),
		).
		ExecX(ctx)
	return
}

func (d *database) GetGuildRoles(guildID uint64) (roles []*entgen.Role, err error) {
	doRecovery(&err)
	roles = d.Guild.GetX(ctx, guildID).QueryRole().AllX(ctx)
	return
}
