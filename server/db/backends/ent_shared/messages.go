package ent_shared

import (
	"database/sql"
	"time"

	proto "github.com/golang/protobuf/proto"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
)

func mustBytes(m proto.Message) []byte {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return data
}

// TODO: overrides
func (d *database) addMessageStem(channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo sql.NullInt64, metadata *harmonytypesv1.Metadata) *entgen.MessageCreate {
	foo := d.Message.Create().
		SetID(messageID).
		SetChannelID(channelID).
		SetUserID(authorID).
		SetMetadata(metadata).
		SetActions(actions).
		SetOverrides(mustBytes(overrides))

	if replyTo.Valid {
		foo = foo.AddReplyIDs(uint64(replyTo.Int64))
	}

	return foo
}

func (d *database) AddTextMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo sql.NullInt64, metadata *harmonytypesv1.Metadata, content string) (t time.Time, e error) {
	defer doRecovery(&e)

	msg := d.addMessageStem(channelID, messageID, authorID, actions, overrides, replyTo, metadata).SaveX(ctx)
	msg.Update().SetTextmessage(d.TextMessage.Create().SetMessage(msg).SetContent(content).SaveX(ctx)).SaveX(ctx)

	return msg.Createdat, nil
}

func (d *database) AddFilesMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo sql.NullInt64, metadata *harmonytypesv1.Metadata, files []*harmonytypesv1.Attachment) (t time.Time, e error) {
	panic("unimplemented")
}
func (d *database) AddEmbedMessage(guildID, channelID, messageID uint64, authorID uint64, actions []*harmonytypesv1.Action, overrides *harmonytypesv1.Override, replyTo sql.NullInt64, metadata *harmonytypesv1.Metadata, embeds []*harmonytypesv1.Embed) (t time.Time, e error) {
	panic("unimplemented")
}
