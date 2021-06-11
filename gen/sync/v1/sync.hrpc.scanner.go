package v1

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (x *Event) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *Event) Scan(src interface{}) error {
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

func (x *Ack) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *Ack) Scan(src interface{}) error {
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

func (x *Syn) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *Syn) Scan(src interface{}) error {
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
