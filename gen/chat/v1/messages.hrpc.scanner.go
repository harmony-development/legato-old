package v1

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (x *GetChannelMessagesRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetChannelMessagesRequest) Scan(src interface{}) error {
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

func (x *GetChannelMessagesResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetChannelMessagesResponse) Scan(src interface{}) error {
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

func (x *GetMessageRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetMessageRequest) Scan(src interface{}) error {
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

func (x *GetMessageResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetMessageResponse) Scan(src interface{}) error {
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

func (x *DeleteMessageRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *DeleteMessageRequest) Scan(src interface{}) error {
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

func (x *TriggerActionRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *TriggerActionRequest) Scan(src interface{}) error {
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

func (x *SendMessageRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *SendMessageRequest) Scan(src interface{}) error {
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

func (x *SendMessageResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *SendMessageResponse) Scan(src interface{}) error {
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

func (x *UpdateMessageTextRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *UpdateMessageTextRequest) Scan(src interface{}) error {
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
