package v1

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (x *CreateEmotePackRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *CreateEmotePackRequest) Scan(src interface{}) error {
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

func (x *CreateEmotePackResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *CreateEmotePackResponse) Scan(src interface{}) error {
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

func (x *GetEmotePacksRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetEmotePacksRequest) Scan(src interface{}) error {
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

func (x *GetEmotePacksResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetEmotePacksResponse) Scan(src interface{}) error {
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

func (x *GetEmotePackEmotesRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetEmotePackEmotesRequest) Scan(src interface{}) error {
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

func (x *GetEmotePackEmotesResponse) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *GetEmotePackEmotesResponse) Scan(src interface{}) error {
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

func (x *AddEmoteToPackRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *AddEmoteToPackRequest) Scan(src interface{}) error {
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

func (x *DeleteEmoteFromPackRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *DeleteEmoteFromPackRequest) Scan(src interface{}) error {
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

func (x *DeleteEmotePackRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *DeleteEmotePackRequest) Scan(src interface{}) error {
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

func (x *DequipEmotePackRequest) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *DequipEmotePackRequest) Scan(src interface{}) error {
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
