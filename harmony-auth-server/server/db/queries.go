package db

import (
	"harmony-auth-server/server/comms"
)

// GetInstanceList gets the instances for a specific user
func (db DB) GetInstanceList(userID string) ([]comms.Instance, error) {
	res, err := db.Query("SELECT name, host FROM instances WHERE userid=$1", userID)
	if err != nil {
		return nil, err
	}
	var out []comms.Instance
	for res.Next() {
		var inst comms.Instance
		if err := res.Scan(&inst.Name, &inst.Host); err != nil {
			return nil, err
		}
		out = append(out, inst)
	}
	return out, nil
}

// AddInstance adds an instance to a user's list
func (db DB) AddInstance(name string, host string, userID string) error {
	_, err := db.Exec("INSERT INTO instances(name, host, userid) VALUES($1, $2, $3)", name, host, userID)
	return err
}

// RemoveInstance removes an instance from a user's list
func (db DB) RemoveInstance(host string, userID string) error {
	res, tx, err := db.Transact("DELETE FROM instances WHERE host=$1 AND userid=$2", host, userID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 1 {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}
	return nil
}

// GetUser returns the public information for a user
func (db DB) GetUser(userID string) (*User, error) {
	user := &User{}
	if err := db.QueryRow("SELECT username, avatar FROM users WHERE userid=$1", userID).Scan(user.Username, user.Avatar); err != nil {
		return nil, err
	}
	return user, nil
}

// RegisterUser inserts a new user into the DB
func (db DB) RegisterUser(userID string, email string, username string, passHash string) error {
	_, err := db.Exec("INSERT INTO users(userid, email, username, avatar, password) VALUES($1, $2, $3, $4, $5)", userID, email, username, "", passHash)
	return err
}
