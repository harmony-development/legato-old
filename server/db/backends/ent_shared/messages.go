package ent_shared

import (
	"time"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/types"
)

// TODO: overrides, actions
func (d *DB) addMessageStem(channelID, messageID uint64, authorID uint64, replyTo *uint64, override *harmonytypesv1.Override, metadata *harmonytypesv1.Metadata) *entgen.MessageCreate {
	msg := d.Message.Create().
		SetID(messageID).
		SetChannelID(channelID).
		SetUserID(authorID).
		SetMetadata(metadata).
		SetCreatedat(time.Now())

	if override != nil {
		msg.SetOverride(override)
	}

	if replyTo != nil {
		msg.AddReplyIDs(*replyTo)
	}

	return msg
}

func (d *DB) AddMessage(guildID, channelID, messageID uint64, authorID uint64, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, content *harmonytypesv1.Content) (t time.Time, e error) {
	defer doRecovery(&e)
	msg :=
		d.addMessageStem(channelID, messageID, authorID, replyTo, overrides, metadata).
			SetContent(content).
			SaveX(ctx)

	t = msg.Createdat

	return
}

func (d *DB) DeleteMessage(messageID uint64) (err error) {
	defer doRecovery(&err)

	d.Message.DeleteOneID(messageID).ExecX(ctx)

	return
}

func (d *DB) GetMessage(messageID uint64) (msg *types.MessageData, err error) {
	defer doRecovery(&err)
	data := d.Message.
		Query().
		Where(
			message.ID(messageID),
		).
		WithChannel(func(cq *entgen.ChannelQuery) {
			cq.WithGuild()
		}).
		WithUser().
		WithParent().
		OnlyX(ctx)

	return types.Into(data), nil
}

func (d *DB) GetMessages(channelID uint64) (msgs []*types.MessageData, err error) {
	defer doRecovery(&err)
	messages := d.Channel.
		GetX(ctx, channelID).
		QueryMessage().
		Limit(50).
		AllX(ctx)

	return types.IntoMany(messages), nil
}

func (d *DB) GetMessagesBefore(channelID uint64, date time.Time) (msgs []*types.MessageData, err error) {
	defer doRecovery(&err)
	messages := d.Message.
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
	return types.IntoMany(messages), nil
}

func (d *DB) HasMessageWithID(messageID uint64) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.Message.Query().Where(message.ID(messageID)).ExistX(ctx)
	return
}

func (d *DB) UpdateTextMessage(messageID uint64, content string) (t time.Time, err error) {
	defer doRecovery(&err)

	data := d.Message.GetX(ctx, messageID)

	v := data.Content.Content.(*harmonytypesv1.Content_TextMessage)
	v.TextMessage.Content = content
	data.Content.Content = v

	ret := data.Update().SetContent(data.Content).SetEditedat(time.Now()).SaveX(ctx)

	return ret.Editedat, nil
}

func (d *DB) GetMessageOwner(messageID uint64) (userID uint64, err error) {
	defer doRecovery(&err)
	userID = d.Message.
		Query().
		Where(
			message.ID(messageID),
			message.HasUserWith(
				user.ID(userID),
			),
		).
		OnlyX(ctx).
		ID
	return
}
