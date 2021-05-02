package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/lexorank"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *DB) AddRoleToGuild(guildID, roleID uint64, name string, color int, hoist, pingable bool) (err error) {
	defer doRecovery(&err)
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

func (d *DB) GetGuildRoles(guildID uint64) (roles []*types.RoleData, err error) {
	defer doRecovery(&err)
	data := d.Guild.GetX(ctx, guildID).QueryRole().AllX(ctx)
	roles = make([]*types.RoleData, len(data))
	for i, entry := range data {
		roles[i] = &types.RoleData{
			ID:       entry.ID,
			Name:     entry.Name,
			Position: entry.Position,
			Color:    entry.Color,
			Hoist:    entry.Hoist,
			Pingable: entry.Pingable,
		}
	}
	return
}

func (d *DB) GetPermissions(roleID uint64) (permissions []types.PermissionsNode, err error) {
	defer doRecovery(&err)
	nodes := d.Role.GetX(ctx, roleID).QueryPermissionNode().AllX(ctx)
	for _, node := range nodes {
		permissions = append(permissions, types.PermissionsNode{
			Node:  node.Node,
			Allow: node.Allow,
		})
	}
	return
}

func (d *DB) SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []types.PermissionsNode) (err error) {
	defer doRecovery(&err)
	tx := d.TxX()
	nodes := make([]*entgen.PermissionNodeCreate, len(permissions))
	for i, node := range permissions {
		nodes[i] = tx.PermissionNode.Create().SetNode(node.Node).SetAllow(node.Allow)
	}
	savedNodes := d.PermissionNode.
		CreateBulk(nodes...).
		SaveX(ctx)
	if roleID != 0 {
		d.Role.
			UpdateOneID(roleID).
			AddPermissionNode(
				savedNodes...,
			).
			ExecX(ctx)
	} else if channelID != 0 {
		d.Channel.
			UpdateOneID(channelID).
			AddPermissionNode(
				savedNodes...,
			).
			ExecX(ctx)
	} else if guildID != 0 {
		d.Guild.
			UpdateOneID(guildID).
			AddPermissionNode(
				savedNodes...,
			).
			ExecX(ctx)
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}
	return
}

func (d *DB) GetPermissionsData(guildID uint64) (data types.PermissionsData, err error) {
	defer doRecovery(&err)
	roles := d.Guild.GetX(ctx, guildID).QueryRole().WithPermissionNode().AllX(ctx)
	for _, role := range roles {
		if perms, err := d.GetPermissions(role.ID); err != nil {
			panic(err)
		} else {
			data.Roles[role.ID] = perms
		}
	}
	chans := d.Guild.Query().Where(guild.ID(guildID)).QueryChannel().Order(entgen.Asc(channel.FieldPosition)).WithRole().AllX(ctx)
	var category uint64 = 0
	data.Channels = make(map[uint64]map[uint64][]types.PermissionsNode)
	for _, c := range chans {
		if c.Kind == uint64(types.ChannelKindCategory) {
			category = c.ID
		} else if category != 0 {
			categoryData, ok := data.Categories[category]
			_ = ok
			data.Categories[category] = append(categoryData, c.ID)
		}
		for _, role := range roles {
			if perms, err := d.GetPermissions(role.ID); err != nil {
				panic(err)
			} else {
				data.Channels[c.ID][role.ID] = perms
			}
		}
	}
	return
}

func (d *DB) MoveRole(guildID, roleID, previousRole, nextRole uint64) (err error) {
	defer doRecovery(&err)
	previousPos := d.Role.GetX(ctx, previousRole).Position
	nextPos := d.Role.GetX(ctx, nextRole).Position
	d.Role.
		UpdateOneID(roleID).
		SetPosition(
			lexorank.Rank(previousPos, nextPos),
		).
		ExecX(ctx)
	return
}

func (d *DB) ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) (err error) {
	defer doRecovery(&err)
	d.User.
		UpdateOneID(userID).
		AddRoleIDs(addRoles...).
		RemoveRoleIDs(removeRoles...).
		ExecX(ctx)
	return
}

func (d *DB) ModifyRole(roleID uint64, name string, color int, hoist, pingable, updateName, updateColor, updateHoist, updatePingable bool) (err error) {
	defer doRecovery(&err)
	update := d.Role.UpdateOneID(roleID)
	if updateName {
		update.SetName(name)
	}
	if updateColor {
		update.SetColor(color)
	}
	if updateHoist {
		update.SetHoist(hoist)
	}
	if updatePingable {
		update.SetPingable(pingable)
	}
	update.ExecX(ctx)
	return
}

func (d *DB) RemoveRoleFromGuild(guildID, roleID uint64) (err error) {
	defer doRecovery(&err)
	d.Guild.
		UpdateOneID(guildID).
		RemoveRoleIDs(roleID).
		ExecX(ctx)
	return
}

func (d *DB) RolesForUser(guildID, userID uint64) (roles []uint64, err error) {
	defer doRecovery(&err)
	roles = d.Role.
		Query().
		Where(
			role.And(
				role.HasMembersWith(
					user.ID(userID),
				),
				role.HasGuildWith(
					guild.ID(guildID),
				),
			),
		).
		IDsX(ctx)
	return
}
