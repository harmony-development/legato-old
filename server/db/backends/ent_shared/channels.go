package ent_shared

import (
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
)

func (d *database) AddChannelToGuild(guildID uint64, channelID uint64, channelName string, previous, next uint64, category bool, md *harmonytypesv1.Metadata) (channel queries.Channel, err error) {
	defer doRecovery(&err)

	// TODO uhh add the channel
	d.Guild.UpdateOneID(guildID).AddChannel()

	return
}
