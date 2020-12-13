package db

import (
	"database/sql"
	"time"

	"github.com/harmony-development/legato/server/db/queries"
)

// AddMessage adds a message to a channel
func (db *HarmonyDB) AddMessage(channelID, guildID, userID, messageID uint64, message string, attachments []string, embeds, actions, overrides []byte, replyTo sql.NullInt64) (*queries.Message, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
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
	db.Logger.CheckException(err)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message
func (db *HarmonyDB) DeleteMessage(messageID, channelID, guildID uint64) error {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	tq := db.queries.WithTx(tx)
	numRows, err := tq.DeleteMessage(ctx, queries.DeleteMessageParams{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	})
	db.Logger.CheckException(err)
	if err != nil {
		return err
	}
	if numRows > 1 { // JUST IN CASE the delete query deletes too much
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		db.Logger.CheckException(err)
		return err
	}
	return nil
}

// GetMessageOwner gets the owner of a messageID
func (db *HarmonyDB) GetMessageOwner(messageID uint64) (uint64, error) {
	owner, err := db.queries.GetMessageAuthor(ctx, messageID)
	db.Logger.CheckException(err)
	return owner, err
}

// GetMessageDate gets the date for a message
func (db *HarmonyDB) GetMessageDate(messageID uint64) (time.Time, error) {
	msgDate, err := db.queries.GetMessageDate(ctx, messageID)
	db.Logger.CheckException(err)
	return msgDate, err
}

// GetMessages gets the newest messages from a guild
func (db *HarmonyDB) GetMessages(guildID, channelID uint64) ([]queries.Message, error) {
	msgs, err := db.GetMessagesBefore(guildID, channelID, time.Now())
	db.Logger.CheckException(err)
	return msgs, err
}

// GetMessagesBefore gets messages before a given message in a guild
func (db *HarmonyDB) GetMessagesBefore(guildID, channelID uint64, date time.Time) ([]queries.Message, error) {
	msgsBefore, err := db.queries.GetMessages(ctx, queries.GetMessagesParams{
		Guildid:   guildID,
		Channelid: channelID,
		Before:    date,
		Max:       int32(db.Config.Server.GetMessageCount),
	})
	db.Logger.CheckException(err)
	return msgsBefore, err
}

// GetMessage gets the data of a message
func (db *HarmonyDB) GetMessage(messageID uint64) (queries.Message, error) {
	return db.queries.GetMessage(ctx, messageID)
}

func (db *HarmonyDB) UpdateMessage(messageID uint64, content *string, embeds, actions, overrides *[]byte, attachments *[]string) (time.Time, error) {
	tx, err := db.Begin()
	db.Logger.CheckException(err)
	if err != nil {
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
			editedAt = data.EditedAt.Time
			return err
		})
	}
	if overrides != nil {
		e.Execute(func() error {
			return tq.UpdateMessageOverrides(ctx, queries.UpdateMessageOverridesParams{
				MessageID: messageID,
				Overrides: *overrides,
			})
		})
	}
	if attachments != nil {
		e.Execute(func() error {
			return tq.UpdateMessageAttachments(ctx, queries.UpdateMessageAttachmentsParams{
				MessageID:   messageID,
				Attachments: *attachments,
			})
		})
	}
	if e.err != nil {
		if err := tx.Rollback(); err != nil {
			return time.Time{}, err
		}
		return time.Time{}, e.err
	}
	if err := tx.Commit(); err != nil {
		return time.Time{}, err
	}
	return editedAt, nil
}

func (db HarmonyDB) HasMessageWithID(guildID, channelID, messageID uint64) (bool, error) {
	return db.queries.MessageWithIDExists(ctx, queries.MessageWithIDExistsParams{
		GuildID:   guildID,
		ChannelID: channelID,
		MessageID: messageID,
	})
}
