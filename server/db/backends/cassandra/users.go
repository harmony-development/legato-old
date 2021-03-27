package cassandra

import (
	"github.com/scylladb/gocqlx/v2/table"
	"github.com/ztrue/tracerr"
)

type User struct {
	UserID      uint64
	LocalUserID uint64
	HomeServer  string
	Email       string
	Username    string
	Avatar      string
}

func (db *database) setupUserTable() {
	db.userTable = table.New(table.Metadata{
		Name:    "user",
		Columns: []string{"user_id", "email", "username", "avatar"},
		PartKey: []string{"user_id"},
		SortKey: []string{"username"},
	})
}

func (db *database) GetLocalUserByID(userID uint64) (*User, error) {
	selectQuery := db.userTable.SelectQuery(db.Session)
	selectQuery.BindStruct(User{
		UserID: userID,
	})
	var result []*User
	if err := selectQuery.Select(&result); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if len(result) == 0 {
		return nil, tracerr.New("no results found")
	}
	return result[0], nil
}
