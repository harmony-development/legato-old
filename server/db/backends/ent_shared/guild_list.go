package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/guildlistentry"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/lexorank"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *DB) GetGuildList(userID uint64) (guilds []*types.GuildListEntryData, err error) {
	defer doRecovery(&err)
	data := d.User.GetX(ctx, userID).QueryListentry().AllX(ctx)
	guilds = make([]*types.GuildListEntryData, len(data))
	for i, entry := range data {
		guilds[i] = &types.GuildListEntryData{
			ID:       entry.ID,
			Host:     entry.Host,
			Position: entry.Position,
		}
	}
	return
}

func (d *DB) GetGuildListPosition(userID, guildID uint64, host string) (pos string, err error) {
	defer doRecovery(&err)
	d.GuildListEntry.Query().Where(
		guildlistentry.And(
			guildlistentry.HasUserWith(user.ID(userID)),
			guildlistentry.Host(host),
			guildlistentry.ID(guildID),
		),
	).OnlyX(ctx)
	return
}

func (d *DB) AddGuildToList(userID, guildID uint64, homeServer string) (err error) {
	defer doRecovery(&err)
	tx := d.TxX()
	tx.User.UpdateOneID(userID).AddListentry(
		tx.GuildListEntry.Create().
			SetID(guildID).SetHost(homeServer).
			SetUserID(userID).
			SaveX(ctx),
	).ExecX(ctx)
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return
}

func (d *DB) MoveGuild(userID, guildID uint64, host string, nextGuildID, prevGuildID uint64, nextHost, prevHost string) (err error) {
	defer doRecovery(&err)
	prevPos, err := d.GetGuildListPosition(userID, prevGuildID, prevHost)
	chk(err)
	nextPos, err := d.GetGuildListPosition(userID, nextGuildID, nextHost)
	chk(err)
	d.GuildListEntry.Update().Where(
		guildlistentry.And(
			guildlistentry.HasUserWith(
				user.ID(userID),
			),
			guildlistentry.ID(guildID),
			guildlistentry.Host(host),
		),
	).SetPosition(
		lexorank.Rank(prevPos, nextPos),
	).ExecX(ctx)
	return
}

func (d *DB) RemoveGuildFromList(userID, guildID uint64, host string) (err error) {
	defer doRecovery(&err)
	d.GuildListEntry.Delete().Where(guildlistentry.And(
		guildlistentry.HasUserWith(
			user.ID(userID),
		),
		guildlistentry.ID(guildID),
		guildlistentry.Host(host),
	),
	).ExecX(ctx)
	return
}
