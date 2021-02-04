package utilities

import (
	"database/sql"
	"encoding/json"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/ztrue/tracerr"
	"google.golang.org/protobuf/proto"
)

// size is roughly 150kb
const size = 18750

func SerializeMetadata(md *harmonytypesv1.Metadata) ([]byte, error) {
	data, err := proto.Marshal(md)
	err = tracerr.Wrap(err)
	if len(data) > size {
		return nil, tracerr.New("serializeMetadata: metadata too large")
	}
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return data, nil
}

// DeserializeMetadata deserializes metadata from the database to protobuf format metadata
func DeserializeMetadata(data []byte) *harmonytypesv1.Metadata {
	md := harmonytypesv1.Metadata{}
	err := tracerr.Wrap(proto.Unmarshal(data, &md))
	if err != nil {
		panic(err)
	}
	return &md
}

func ToSqlString(input string) sql.NullString {
	return sql.NullString{String: input, Valid: true}
}

func ToSqlInt64(input uint64) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(input), Valid: true}
}

type Executor struct {
	Err error
}

func (e *Executor) Execute(f func() error) {
	if e.Err != nil {
		return
	}
	e.Err = f()
}

func MustDeserialize(d json.RawMessage, v interface{}) {
	err := json.Unmarshal(d, v)
	if err != nil {
		panic(err)
	}
}

func MustSerialize(v interface{}) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}
