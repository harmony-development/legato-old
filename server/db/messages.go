package db

import (
	"database/sql"
	"time"

	"github.com/harmony-development/legato/server/db/queries"
	"github.com/ztrue/tracerr"
)

// AddMessage adds a message to a channel
func (db *HarmonyDB) AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64) (*queries.Message, error) {
	tx, err := db.Begin()
	if err != nil {
		tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	tq := db.queries.WithTx(tx)
	msg, err := tq.AddMessage(ctx, queries.AddMessageParams{
		GuildID:     guildID,
		ChannelID:   channelID,
		UserID:      userID,
		MessageID:   messageID,
		Content:     message,
		Embeds:      embeds,
		Actions:     actions,
		Overrides:   overrides,
		Attachments: attachments,
		ReplyToID:   replyTo,
	})
	if err != nil {
		tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID, channelID, guildID uint64) (err error) {
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
func (db *HarmonyDB) GetMessageOwner(messageID uint64) (owner uint64, err error) {
	owner, err = db.queries.GetMessageAuthor(ctx, messageID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return owner, err
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID uint64) (t time.Time, err error) {
	msgDate, err := db.queries.GetMessageDate(ctx, messageID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return msgDate, err
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID, channelID uint64) (r []queries.Message, err error) {
	msgs, err := db.GetMessagesBefore(guildID, channelID, time.Now())
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return msgs, err
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID, channelID uint64, date time.Time) (r []queries.Message, err error) {
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
func (db *HarmonyDB) GetMessage(messageID uint64) (r queries.Message, err error) {
	r, err = db.queries.GetMessage(ctx, messageID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return
}

func (db *HarmonyDB) UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string) (r time.Time, err error) {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return time.Time{}, err
	}
	tq := db.queries.WithTx(tx)
	var editedAt time.Time
	e := executor{}
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
	if e.err != nil {
		if err := tx.Rollback(); err != nil {
			err = tracerr.Wrap(err)
			return time.Time{}, err
		}
		err = tracerr.Wrap(err)
		return time.Time{}, e.err
	}
	if err := tx.Commit(); err != nil {
		err = tracerr.Wrap(err)
		return time.Time{}, err
	}
	return editedAt, nil
}

func (db HarmonyDB) HasMessageWithID(guildID, channelID, messageID uint64) (r bool, err error) {
	r, err = db.queries.MessageWithIDExists(ctx, queries.MessageWithIDExistsParams{
		GuildID:   guildID,
		ChannelID: channelID,
		MessageID: messageID,
	})
	err = tracerr.Wrap(err)
	return
}
