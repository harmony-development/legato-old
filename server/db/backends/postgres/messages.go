package postgres

import (
	"database/sql"
	"time"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/db/utilities"
	"github.com/ztrue/tracerr"
	"google.golang.org/protobuf/proto"
)

type MessageKind int32

const (
	MessageKindText = iota
	MessageKindPhoto
	MessageKindEmbed
	MessageKindFiles
)

// AddMessage adds a message to a channel
func (db *database) AddMessage(guildID, channelID, userID, messageID uint64, replyTo sql.NullInt64, metadata *harmonytypesv1.Metadata, message *harmonytypesv1.Content) (*queries.Message, error) {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	data, err := utilities.SerializeMetadata(metadata)
	if err != nil {
		return nil, err
	}
	var messageData []byte
	var kind MessageKind
	switch v := message.Content.(type) {
	case *harmonytypesv1.Content_TextMessage:
		kind = MessageKindText
		messageData, err = proto.Marshal(v.TextMessage)
		if err != nil {
			return nil, err
		}
	case *harmonytypesv1.Content_PhotoMessage:
		kind = MessageKindPhoto
		messageData, err = proto.Marshal(v.PhotoMessage)
		if err != nil {
			return nil, err
		}
	case *harmonytypesv1.Content_EmbedMessage:
		kind = MessageKindEmbed
		messageData, err = proto.Marshal(v.EmbedMessage)
		if err != nil {
			return nil, err
		}
	case *harmonytypesv1.Content_FilesMessage:
		kind = MessageKindFiles
		messageData, err = proto.Marshal(v.FilesMessage)
		if err != nil {
			return nil, err
		}
	}

	msg, err := tq.AddMessage(ctx, queries.AddMessageParams{
		GuildID:   guildID,
		ChannelID: channelID,
		UserID:    userID,
		MessageID: messageID,
		Kind:      int32(kind),
		Content:   messageData,
		ReplyToID: replyTo,
		Metadata:  data,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message
func (db *database) DeleteMessage(messageID, channelID, guildID uint64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	tq := db.queries.WithTx(tx)
	numRows, err := tq.DeleteMessage(ctx, queries.DeleteMessageParams{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	if numRows > 1 { // JUST IN CASE the delete query deletes too much
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

// GetMessageOwner gets the owner of a messageID
func (db *database) GetMessageOwner(messageID uint64) (owner uint64, err error) {
	owner, err = db.queries.GetMessageAuthor(ctx, messageID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return owner, err
}

// GetMessageDate gets the date for a message
func (db *database) GetMessageDate(messageID uint64) (t time.Time, err error) {
	msgDate, err := db.queries.GetMessageDate(ctx, messageID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return msgDate, err
}

// GetMessages gets the newest messages from a guild
func (db *database) GetMessages(guildID, channelID uint64) (r []queries.Message, err error) {
	msgs, err := db.GetMessagesBefore(guildID, channelID, time.Now())
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return msgs, err
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *database) GetMessagesBefore(guildID, channelID uint64, date time.Time) (r []queries.Message, err error) {
	msgsBefore, err := db.queries.GetMessages(ctx, queries.GetMessagesParams{
		Guildid:   guildID,
		Channelid: channelID,
		Before:    date,
		Max:       int32(db.Config.Server.Policies.APIs.Messages.MaximumGetAmount),
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return msgsBefore, err
}

// GetMessage gets the data of a message
func (db *database) GetMessage(messageID uint64) (r queries.Message, err error) {
	r, err = db.queries.GetMessage(ctx, messageID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return
}

func (db *database) UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string, metadata *harmonytypesv1.Metadata, updateMetadata bool) (time.Time, error) {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return time.Time{}, err
	}
	tq := db.queries.WithTx(tx)
	var editedAt time.Time
	e := utilities.Executor{}
	if content != nil {
		e.Execute(func() error {
			data, err := tq.UpdateMessageContent(ctx, queries.UpdateMessageContentParams{
				MessageID: messageID,
				Content:   *content,
			})
			err = tracerr.Wrap(err)
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if embeds != nil {
		e.Execute(func() error {
			data, err := tq.UpdateMessageEmbeds(ctx, queries.UpdateMessageEmbedsParams{
				MessageID: messageID,
				Embeds:    *embeds,
			})
			err = tracerr.Wrap(err)
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if actions != nil {
		e.Execute(func() error {
			data, err := tq.UpdateMessageActions(ctx, queries.UpdateMessageActionsParams{
				MessageID: messageID,
				Actions:   *actions,
			})
			err = tracerr.Wrap(err)
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if overrides != nil {
		e.Execute(func() error {
			return tracerr.Wrap(tq.UpdateMessageOverrides(ctx, queries.UpdateMessageOverridesParams{
				MessageID: messageID,
				Overrides: *overrides,
			}))
		})
	}
	if attachments != nil {
		e.Execute(func() error {
			return tracerr.Wrap(tq.UpdateMessageAttachments(ctx, queries.UpdateMessageAttachmentsParams{
				MessageID:   messageID,
				Attachments: *attachments,
			}))
		})
	}
	if updateMetadata {
		e.Execute(func() error {
			data, err := utilities.SerializeMetadata(metadata)
			if err != nil {
				return err
			}
			return tracerr.Wrap(tq.UpdateMessageMetadata(ctx, queries.UpdateMessageMetadataParams{
				MessageID: messageID,
				Metadata:  data,
			}))
		})
	}
	if e.Err != nil {
		if err := tx.Rollback(); err != nil {
			err = tracerr.Wrap(err)
			return time.Time{}, err
		}
		err = tracerr.Wrap(e.Err)
		return time.Time{}, err
	}
	if err := tx.Commit(); err != nil {
		err = tracerr.Wrap(err)
		return time.Time{}, err
	}
	return editedAt, nil
}

func (db database) HasMessageWithID(guildID, channelID, messageID uint64) (r bool, err error) {
	r, err = db.queries.MessageWithIDExists(ctx, queries.MessageWithIDExistsParams{
		GuildID:   guildID,
		ChannelID: channelID,
		MessageID: messageID,
	})
	err = tracerr.Wrap(err)
	return
}
