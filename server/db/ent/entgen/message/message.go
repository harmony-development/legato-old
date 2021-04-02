// Code generated by entc, DO NOT EDIT.

package message

import (
	"time"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedat holds the string denoting the createdat field in the database.
	FieldCreatedat = "createdat"
	// FieldEditedat holds the string denoting the editedat field in the database.
	FieldEditedat = "editedat"
	// FieldActions holds the string denoting the actions field in the database.
	FieldActions = "actions"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldOverrides holds the string denoting the overrides field in the database.
	FieldOverrides = "overrides"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeChannel holds the string denoting the channel edge name in mutations.
	EdgeChannel = "channel"
	// EdgeOverride holds the string denoting the override edge name in mutations.
	EdgeOverride = "override"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeReplies holds the string denoting the replies edge name in mutations.
	EdgeReplies = "replies"
	// EdgeTextmessage holds the string denoting the textmessage edge name in mutations.
	EdgeTextmessage = "textmessage"
	// EdgeFilemessage holds the string denoting the filemessage edge name in mutations.
	EdgeFilemessage = "filemessage"
	// EdgeEmbedmessage holds the string denoting the embedmessage edge name in mutations.
	EdgeEmbedmessage = "embedmessage"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "messages"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_message"
	// ChannelTable is the table the holds the channel relation/edge.
	ChannelTable = "messages"
	// ChannelInverseTable is the table name for the Channel entity.
	// It exists in this package in order to avoid circular dependency with the "channel" package.
	ChannelInverseTable = "channels"
	// ChannelColumn is the table column denoting the channel relation/edge.
	ChannelColumn = "channel_message"
	// OverrideTable is the table the holds the override relation/edge.
	OverrideTable = "overrides"
	// OverrideInverseTable is the table name for the Override entity.
	// It exists in this package in order to avoid circular dependency with the "override" package.
	OverrideInverseTable = "overrides"
	// OverrideColumn is the table column denoting the override relation/edge.
	OverrideColumn = "message_override"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "messages"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "message_replies"
	// RepliesTable is the table the holds the replies relation/edge.
	RepliesTable = "messages"
	// RepliesColumn is the table column denoting the replies relation/edge.
	RepliesColumn = "message_replies"
	// TextmessageTable is the table the holds the textmessage relation/edge.
	TextmessageTable = "text_messages"
	// TextmessageInverseTable is the table name for the TextMessage entity.
	// It exists in this package in order to avoid circular dependency with the "textmessage" package.
	TextmessageInverseTable = "text_messages"
	// TextmessageColumn is the table column denoting the textmessage relation/edge.
	TextmessageColumn = "message_textmessage"
	// FilemessageTable is the table the holds the filemessage relation/edge.
	FilemessageTable = "messages"
	// FilemessageInverseTable is the table name for the FileMessage entity.
	// It exists in this package in order to avoid circular dependency with the "filemessage" package.
	FilemessageInverseTable = "file_messages"
	// FilemessageColumn is the table column denoting the filemessage relation/edge.
	FilemessageColumn = "message_filemessage"
	// EmbedmessageTable is the table the holds the embedmessage relation/edge.
	EmbedmessageTable = "messages"
	// EmbedmessageInverseTable is the table name for the EmbedMessage entity.
	// It exists in this package in order to avoid circular dependency with the "embedmessage" package.
	EmbedmessageInverseTable = "embed_messages"
	// EmbedmessageColumn is the table column denoting the embedmessage relation/edge.
	EmbedmessageColumn = "message_embedmessage"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldCreatedat,
	FieldEditedat,
	FieldActions,
	FieldMetadata,
	FieldOverrides,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"channel_message",
	"message_replies",
	"message_filemessage",
	"message_embedmessage",
	"user_message",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedat holds the default value on creation for the "createdat" field.
	DefaultCreatedat func() time.Time
)
