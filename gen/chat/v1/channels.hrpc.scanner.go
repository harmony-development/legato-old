package v1

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (x *CreateChannelRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *CreateChannelRequest) Scan(src interface{}) error {
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

func (x *CreateChannelResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *CreateChannelResponse) Scan(src interface{}) error {
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

func (x *GetGuildChannelsRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetGuildChannelsRequest) Scan(src interface{}) error {
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

func (x *GetGuildChannelsResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetGuildChannelsResponse) Scan(src interface{}) error {
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

func (x *UpdateChannelInformationRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *UpdateChannelInformationRequest) Scan(src interface{}) error {
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

func (x *UpdateChannelOrderRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *UpdateChannelOrderRequest) Scan(src interface{}) error {
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

func (x *DeleteChannelRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *DeleteChannelRequest) Scan(src interface{}) error {
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

func (x *TypingRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *TypingRequest) Scan(src interface{}) error {
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
