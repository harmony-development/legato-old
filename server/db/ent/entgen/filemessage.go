// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/filemessage"
)

// FileMessage is the model entity for the FileMessage schema.
type FileMessage struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileMessageQuery when eager-loading is set.
	Edges FileMessageEdges `json:"edges"`
}

// FileMessageEdges holds the relations/edges for other nodes in the graph.
type FileMessageEdges struct {
	// File holds the value of the file edge.
	File []*File `json:"file,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FileOrErr returns the File value or an error if the edge
// was not loaded in eager-loading.
func (e FileMessageEdges) FileOrErr() ([]*File, error) {
	if e.loadedTypes[0] {
		return e.File, nil
	}
	return nil, &NotLoadedError{edge: "file"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileMessage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case filemessage.FieldID:
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type FileMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileMessage fields.
func (fm *FileMessage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case filemessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fm.ID = int(value.Int64)
		}
	}
	return nil
}

// QueryFile queries the "file" edge of the FileMessage entity.
func (fm *FileMessage) QueryFile() *FileQuery {
	return (&FileMessageClient{config: fm.config}).QueryFile(fm)
}

// Update returns a builder for updating this FileMessage.
// Note that you need to call FileMessage.Unwrap() before calling this method if this FileMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (fm *FileMessage) Update() *FileMessageUpdateOne {
	return (&FileMessageClient{config: fm.config}).UpdateOne(fm)
}

// Unwrap unwraps the FileMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fm *FileMessage) Unwrap() *FileMessage {
	tx, ok := fm.config.driver.(*txDriver)
	if !ok {
		panic("entgen: FileMessage is not a transactional entity")
	}
	fm.config.driver = tx.drv
	return fm
}

// String implements the fmt.Stringer.
func (fm *FileMessage) String() string {
	var builder strings.Builder
	builder.WriteString("FileMessage(")
	builder.WriteString(fmt.Sprintf("id=%v", fm.ID))
	builder.WriteByte(')')
	return builder.String()
}

// FileMessages is a parsable slice of FileMessage.
type FileMessages []*FileMessage

func (fm FileMessages) config(cfg config) {
	for _i := range fm {
		fm[_i].config = cfg
	}
}
