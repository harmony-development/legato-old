package cassandra

import (
	"testing"

	"github.com/harmony-development/legato/server/test"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/require"
)

func TestUserGet(t *testing.T) {
	a := require.New(t)
	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{})
	db, err := New(test.DefaultConf(), test.MockLogger{}, sonyflake)
	a.NoError(err)
	user, err := db.GetLocalUserByID(0)
	a.NoError(err)
	a.NotNil(user)
}
