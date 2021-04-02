package ent_shared

import (
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
)

// TODO: finish
func (d *database) AddChannelToGuild(guildID uint64, channelName string, previous, next uint64, category bool, md *harmonytypesv1.Metadata) (q queries.Channel, err error) {
	defer doRecovery(&err)

	d.Guild.UpdateOneID(guildID).AddChannel(
		d.Channel.Create().
			SetName(channelName).
			SaveX(ctx),
	)

	return
}
