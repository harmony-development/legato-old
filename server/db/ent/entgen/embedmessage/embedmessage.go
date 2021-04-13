// Code generated by entc, DO NOT EDIT.

package embedmessage

const (
	// Label holds the string label denoting the embedmessage type in the database.
	Label = "embed_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// EdgeEmbedField holds the string denoting the embed_field edge name in mutations.
	EdgeEmbedField = "embed_field"
	// EdgeMessage holds the string denoting the message edge name in mutations.
	EdgeMessage = "message"
	// Table holds the table name of the embedmessage in the database.
	Table = "embed_messages"
	// EmbedFieldTable is the table the holds the embed_field relation/edge.
	EmbedFieldTable = "embed_fields"
	// EmbedFieldInverseTable is the table name for the EmbedField entity.
	// It exists in this package in order to avoid circular dependency with the "embedfield" package.
	EmbedFieldInverseTable = "embed_fields"
	// EmbedFieldColumn is the table column denoting the embed_field relation/edge.
	EmbedFieldColumn = "embed_message_embed_field"
	// MessageTable is the table the holds the message relation/edge.
	MessageTable = "embed_messages"
	// MessageInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessageInverseTable = "messages"
	// MessageColumn is the table column denoting the message relation/edge.
	MessageColumn = "message_embed_message"
)

// Columns holds all SQL columns for embedmessage fields.
var Columns = []string{
	FieldID,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "embed_messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"message_embed_message",
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
