package db

import (
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/ztrue/tracerr"
	"google.golang.org/protobuf/proto"
)

// size is roughly 150kb
const size = 18750

func serializeMetadata(md *harmonytypesv1.Metadata) ([]byte, error) {
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
