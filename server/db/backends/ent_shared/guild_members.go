package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
)

func (d *DB) GetOwner(guildID uint64) (userID uint64, err error) {
	defer doRecovery(&err)

	userID = d.Guild.QueryOwner(d.Guild.GetX(ctx, guildID)).OnlyIDX(ctx)

	return
}

func (d *DB) IsOwner(guildID, userID uint64) (isOwner bool, err error) {
	defer doRecovery(&err)

	if owner, err := d.GetOwner(guildID); err != nil {
		panic(err)
	} else {
		isOwner = userID == owner
	}
	return
}

func (d *DB) AddMemberToGuild(userID, guildID uint64) (err error) {
	defer doRecovery(&err)

	d.Guild.UpdateOneID(guildID).AddUserIDs(userID).ExecX(ctx)

	return
}

func (d *DB) DeleteMember(guildID, userID uint64) (err error) {
	defer doRecovery(&err)

	d.Guild.
		Update().
		Where(guild.ID(guildID)).
		RemoveUserIDs(userID)

	return
}

func (d *DB) MembersInGuild(guildID uint64) (users []uint64, err error) {
	defer doRecovery(&err)

	users = d.Guild.
		Query().
		Where(guild.ID(guildID)).
		QueryUser().
		IDsX(ctx)

	return
}

func (d *DB) CountMembersInGuild(guildID uint64) (memberCount int64, err error) {
	defer doRecovery(&err)

	memberCount = int64(d.Guild.
		Query().
		Where(guild.ID(guildID)).
		QueryUser().
		CountX(ctx),
	)

	return
}
