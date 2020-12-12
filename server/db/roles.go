package db

import (
	"database/sql"
	"encoding/json"

	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/db/queries"
)

func (db HarmonyDB) AddRoleToGuild(guildID uint64, role *corev1.Role) error {
	_, err := db.queries.CreateRole(ctx, queries.CreateRoleParams{
		GuildID:  guildID,
		RoleID:   role.RoleId,
		Name:     role.Name,
		Color:    role.Color,
		Hoist:    role.Hoist,
		Pingable: role.Pingable,
	})
	return err
}

func (db HarmonyDB) RemoveRoleFromGuild(guildID, roleID uint64) error {
	return db.queries.DeleteRole(ctx, queries.DeleteRoleParams{
		GuildID: guildID,
		RoleID:  roleID,
	})
}

func (db *HarmonyDB) GetRolePositions(guildID, before, previous uint64) (pos string, retErr error) {
	nextPos, err := db.queries.GetRolePosition(ctx, queries.GetRolePositionParams{
		GuildID: guildID,
		RoleID:  before,
	})
	if err != nil && err != sql.ErrNoRows {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	prevPos, err := db.queries.GetRolePosition(ctx, queries.GetRolePositionParams{
		GuildID: guildID,
		RoleID:  previous,
	})
	if err != nil && err != sql.ErrNoRows {
		db.Logger.Exception(err)
		retErr = err
		return
	}
	pos = Rank(prevPos, nextPos)
	return
}

func (db HarmonyDB) MoveRole(guildID, roleID, beforeRole, previousRole uint64) error {
	pos, err := db.GetRolePositions(guildID, beforeRole, previousRole)
	if err != nil {
		return err
	}
	err = db.queries.MoveRole(ctx, queries.MoveRoleParams{
		Position: pos,
		RoleID:   roleID,
		GuildID:  guildID,
	})

	db.Logger.CheckException(err)

	return err
}

func (db HarmonyDB) GetGuildRoles(guildID uint64) (ret []*corev1.Role, err error) {
	roles, err := db.queries.GetRolesForGuild(ctx, guildID)

	for _, role := range roles {
		ret = append(ret, &corev1.Role{
			Name:     role.Name,
			RoleId:   role.RoleID,
			Color:    role.Color,
			Hoist:    role.Hoist,
			Pingable: role.Pingable,
		})
	}

	return
}

func (db HarmonyDB) SetPermissions(guildID uint64, channelID uint64, roleID uint64, permissions []PermissionsNode) error {
	var ln int
	for _, perm := range permissions {
		if perm.Node == "" {
			continue
		}
		permissions[ln] = perm
		ln++
	}
	permissions = permissions[:ln]

	var exists bool
	var err error

	if channelID == 0 {
		exists, err = db.queries.PermissionExistsWithoutChannel(ctx, queries.PermissionExistsWithoutChannelParams{
			GuildID: guildID,
			RoleID:  roleID,
		})
	} else {
		exists, err = db.queries.PermissionsExists(ctx, queries.PermissionsExistsParams{
			GuildID: guildID,
			ChannelID: sql.NullInt64{
				Int64: int64(channelID),
				Valid: true,
			},
			RoleID: roleID,
		})
	}

	if err != nil {
		return err
	}

	if exists {
		if channelID == 0 {
			return db.queries.UpdatePermissionsWithoutChannel(ctx, queries.UpdatePermissionsWithoutChannelParams{
				GuildID: guildID,
				RoleID:  roleID,
				Nodes:   mustSerialize(permissions),
			})
		}
		return db.queries.UpdatePermissions(ctx, queries.UpdatePermissionsParams{
			GuildID: guildID,
			ChannelID: sql.NullInt64{
				Int64: int64(channelID),
				Valid: true,
			},
			RoleID: roleID,
			Nodes:  mustSerialize(permissions),
		})
	}
	return db.queries.SetPermissions(ctx, queries.SetPermissionsParams{
		GuildID: guildID,
		ChannelID: sql.NullInt64{
			Int64: int64(channelID),
			Valid: channelID != 0,
		},
		RoleID: roleID,
		Nodes:  mustSerialize(permissions),
	})
}

func (db HarmonyDB) GetPermissions(guildID uint64, channelID uint64, roleID uint64) (permissions []PermissionsNode, err error) {
	var data json.RawMessage

	if channelID == 0 {
		data, err = db.queries.GetPermissionsWithoutChannel(ctx, queries.GetPermissionsWithoutChannelParams{
			GuildID: guildID,
			RoleID:  roleID,
		})
		println(string(data))
	} else {
		data, err = db.queries.GetPermissions(ctx, queries.GetPermissionsParams{
			GuildID: guildID,
			ChannelID: sql.NullInt64{
				Int64: int64(channelID),
				Valid: true,
			},
			RoleID: roleID,
		})
	}

	if err != nil && err != sql.ErrNoRows {
		return
	} else if len(data) == 0 {
		d := "[]"
		data = json.RawMessage(d)
	}

	mustDeserialize(data, &permissions)

	return
}

func (db HarmonyDB) GetPermissionsData(guildID uint64) (ret PermissionsData, err error) {
	ret.Roles = make(map[uint64][]PermissionsNode)

	roles, err := db.queries.GetRolesForGuild(ctx, guildID)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	for _, role := range roles {
		perms, err := db.GetPermissions(guildID, 0, role.RoleID)
		if err != nil && err != sql.ErrNoRows {
			return PermissionsData{}, err
		}
		ret.Roles[role.RoleID] = perms
	}

	ret.Categories = make(map[uint64][]uint64)
	chans, err := db.ChannelsForGuild(guildID)
	if err != nil {
		return
	}

	cat := uint64(0)
	for _, channel := range chans {
		if channel.Category {
			cat = channel.ChannelID
		} else if cat != 0 {
			data, _ := ret.Categories[cat]
			ret.Categories[cat] = append(data, channel.ChannelID)
		}
	}

	ret.Channels = make(map[uint64]map[uint64][]PermissionsNode)
	for _, channel := range chans {
		ret.Channels[channel.ChannelID] = make(map[uint64][]PermissionsNode)
		for _, role := range roles {
			perms, err := db.GetPermissions(guildID, channel.ChannelID, role.RoleID)
			if err != nil && err != sql.ErrNoRows {
				return PermissionsData{}, err
			}
			ret.Channels[channel.ChannelID][role.RoleID] = perms
		}
	}

	return
}

func (db HarmonyDB) RolesForUser(guildID, userID uint64) (ret []uint64, err error) {
	return db.queries.RolesForUser(ctx, queries.RolesForUserParams{
		GuildID:  guildID,
		MemberID: userID,
	})
}

func (db HarmonyDB) ModifyRole(guildID, roleID uint64, name string, color int32, hoist, pingable, updateName, updateColor, updateHoist, updatePingable bool) error {
	tx, err := db.Begin()
	if err != nil {
		db.Logger.CheckException(err)
		return err
	}

	quer := db.queries.WithTx(tx)

	if updateName {
		err = quer.SetRoleName(ctx, queries.SetRoleNameParams{
			GuildID: guildID,
			RoleID:  roleID,
			Name:    name,
		})
	}
	if updateColor && err == nil {
		err = quer.SetRoleColor(ctx, queries.SetRoleColorParams{
			GuildID: guildID,
			RoleID:  roleID,
			Color:   color,
		})
	}
	if updateHoist && err == nil {
		err = quer.SetRoleHoist(ctx, queries.SetRoleHoistParams{
			GuildID: guildID,
			RoleID:  roleID,
			Hoist:   hoist,
		})
	}
	if updatePingable && err == nil {
		err = quer.SetRolePingable(ctx, queries.SetRolePingableParams{
			GuildID:  guildID,
			RoleID:   roleID,
			Pingable: pingable,
		})
	}
	if err != nil {
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

func (db HarmonyDB) ManageRoles(guildID, userID uint64, addRoles, removeRoles []uint64) error {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}

	quer := db.queries.WithTx(tx)

	for _, add := range addRoles {
		err = quer.AddUserToRole(ctx, queries.AddUserToRoleParams{
			GuildID:  guildID,
			MemberID: userID,
			RoleID:   add,
		})
		if err != nil {
			return err
		}
	}
	for _, remove := range removeRoles {
		err = quer.RemoveUserFromRole(ctx, queries.RemoveUserFromRoleParams{
			GuildID:  guildID,
			MemberID: userID,
			RoleID:   remove,
		})
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return err
	}

	return nil
}
