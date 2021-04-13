

package any

import "database/sql/driver"
import "google.golang.org/protobuf/proto"
import "fmt"






func (x *Any) Value() (driver.Value, error) {
	return proto.Marshal(x)
}

func (x *Any) Scan(src interface{}) error {
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


