package ent_shared

import (
	"time"

	proto "github.com/golang/protobuf/proto"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/textmessage"
)

func mustBytes(m proto.Message) []byte {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return data
}

// TODO: overrides, actions
func (d *database) addMessageStem(channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata) *entgen.MessageCreate {
	msg := d.Message.Create().
		SetID(messageID).
		SetChannelID(channelID).
		SetUserID(authorID).
		SetMetadata(metadata).
		SetActions(actions).
		SetOverrides(mustBytes(overrides))

	if replyTo != nil {
		msg.AddReplyIDs(*replyTo)
	}

	return msg
}

func (d *database) AddTextMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, content string) (t time.Time, e error) {
	defer doRecovery(&e)

	msg := d.addMessageStem(channelID, messageID, authorID, actions, overrides, replyTo, metadata).SaveX(ctx)
	msg.Update().SetTextmessage(d.TextMessage.Create().SetMessage(msg).SetContent(content).SaveX(ctx)).SaveX(ctx)

	return msg.Createdat, nil
}

func (d *database) AddFilesMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, files []*harmonytypesv1.Attachment) (t time.Time, e error) {
	panic("unimplemented")
}
func (d *database) AddEmbedMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, embeds []*harmonytypesv1.Embed) (t time.Time, e error) {
	panic("unimplemented")
}

func (d *database) DeleteMessage(messageID uint64) (err error) {
	defer doRecovery(&err)

	d.Message.DeleteOneID(messageID).ExecX(ctx)

	return
}

func (d *database) GetMessage(messageID uint64) (msg *entgen.Message, err error) {
	defer doRecovery(&err)
	msg = d.Message.GetX(ctx, messageID)
	return
}

func (d *database) GetMessages(channelID uint64) (msgs []*entgen.Message, err error) {
	defer doRecovery(&err)
	msgs = d.Channel.
		GetX(ctx, channelID).
		QueryMessage().
		Limit(50).
		AllX(ctx)
	return
}

func (d *database) GetMessagesBefore(channelID uint64, date time.Time) (msgs []*entgen.Message, err error) {
	defer doRecovery(&err)
	msgs = d.Message.
		Query().
		Limit(50).
		Where(
			message.And(
				message.CreatedatLT(date),
				message.HasChannelWith(
					channel.ID(channelID),
				),
			),
		).
		AllX(ctx)
	return
}

func (d *database) HasMessageWithID(messageID uint64) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.Message.Query().Where(message.ID(messageID)).ExistX(ctx)
	return
}

func (d *database) UpdateTextMessage(messageID uint64, content string) (t time.Time, err error) {
	defer doRecovery(&err)
	d.TextMessage.
		Update().
		Where(
			textmessage.HasMessageWith(
				message.ID(messageID),
			),
		).
		SetContent(content).
		ExecX(ctx)
	return
}
