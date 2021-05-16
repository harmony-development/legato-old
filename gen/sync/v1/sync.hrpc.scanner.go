package v1

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (x *SyncRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *SyncRequest) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if b, ok := src.([]byte); ok {
		if err := proto.Unmarshal(b, x); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unexpected type %T", src)
}

func (x *PostBoxEvent) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *PostBoxEvent) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if b, ok := src.([]byte); ok {
		if err := proto.Unmarshal(b, x); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unexpected type %T", src)
}

func (x *PostEventRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *PostEventRequest) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if b, ok := src.([]byte); ok {
		if err := proto.Unmarshal(b, x); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unexpected type %T", src)
}
