package ent_shared

import (
	"time"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/textmessage"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (d *DB) AddTextMessage(guildID, channelID, messageID uint64, authorID uint64, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, content string) (t time.Time, e error) {
	defer doRecovery(&e)
	msg :=
		d.addMessageStem(channelID, messageID, authorID, replyTo, overrides, metadata).
			SetTextMessage(
				d.TextMessage.
					Create().
					SetContent(content).
					SaveX(ctx),
			).
			SaveX(ctx)

	t = msg.Createdat

	return
}

func (d *DB) AddFilesMessage(guildID, channelID, messageID uint64, authorID uint64, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, files []string) (t time.Time, e error) {
	defer doRecovery(&e)
	msg :=
		d.addMessageStem(channelID, messageID, authorID, replyTo, overrides, metadata).
			SetFileMessage(
				d.FileMessage.
					Create().
					AddFileIDs(files...).
					SaveX(ctx),
			).
			SaveX(ctx)

	t = msg.Createdat

	return
}

func (d *DB) AddEmbedMessage(guildID, channelID, messageID uint64, authorID uint64, overrides *harmonytypesv1.Override, replyTo *uint64, metadata *harmonytypesv1.Metadata, embed *harmonytypesv1.Embed) (t time.Time, e error) {
	defer doRecovery(&e)
	msg :=
		d.addMessageStem(channelID, messageID, authorID, replyTo, overrides, metadata).
			SetEmbedMessage(
				d.EmbedMessage.Create().
					SetData(embed).
					SaveX(ctx),
			).
			SaveX(ctx)
	t = msg.Createdat
	return
}

func (d *DB) DeleteMessage(messageID uint64) (err error) {
	defer doRecovery(&err)

	d.Message.DeleteOneID(messageID).ExecX(ctx)

	return
}

func (d *DB) GetMessage(messageID uint64) (msg *harmonytypesv1.Message, err error) {
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
		WithTextMessage().
		WithEmbedMessage().
		WithFileMessage(func(fmq *entgen.FileMessageQuery) {
			fmq.WithFile()
		}).
		OnlyX(ctx)

	msg = &harmonytypesv1.Message{
		GuildId:   data.Edges.Channel.Edges.Guild.ID,
		ChannelId: data.Edges.Channel.ID,
		MessageId: data.ID,
		AuthorId:  data.Edges.User.ID,
		CreatedAt: timestamppb.New(data.Createdat),
		EditedAt:  timestamppb.New(data.Editedat),
		InReplyTo: data.Edges.Parent.ID,
	}

	if data.Edges.TextMessage != nil {
		msg.Content = &harmonytypesv1.Content{
			Content: &harmonytypesv1.Content_TextMessage{
				TextMessage: &harmonytypesv1.ContentText{
					Content: data.Edges.TextMessage.Content,
				},
			},
		}
	} else if data.Edges.FileMessage != nil {
		msg.Content = &harmonytypesv1.Content{
			Content: &harmonytypesv1.Content_FilesMessage{
				FilesMessage: &harmonytypesv1.ContentFiles{
					Attachments: data.Edges.FileMessage.Edges.File,
				},
			},
		}
	}

	return
}

func (d *DB) GetMessages(channelID uint64) (msgs []*harmonytypesv1.Message, err error) {
	defer doRecovery(&err)
	msgs = d.Channel.
		GetX(ctx, channelID).
		QueryMessage().
		Limit(50).
		AllX(ctx)
	return
}

func (d *DB) GetMessagesBefore(channelID uint64, date time.Time) (msgs []*entgen.Message, err error) {
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

func (d *DB) HasMessageWithID(messageID uint64) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.Message.Query().Where(message.ID(messageID)).ExistX(ctx)
	return
}

func (d *DB) UpdateTextMessage(messageID uint64, content string) (t time.Time, err error) {
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
