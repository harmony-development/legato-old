package db

// User is the schema for the users table
type User struct {
	UserID   string
	Email    string
	Username string
	Avatar   string
	Password string
}

var queries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		userid TEXT PRIMARY KEY NOT NULL, 
		email TEXT UNIQUE NOT NULL, 
		username TEXT NOT NULL,
		avatar TEXT NOT NULL,
		password TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS instances(
		userid TEXT NOT NULL REFERENCES users(userid), -- the userid of who owns this entry
		host TEXT PRIMARY KEY NOT NULL, -- the host for the harmony instance
		name TEXT NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS sessions(
		session TEXT NOT NULL,
		userid TEXT NOT NULL REFERENCES users(userid)
	)`,
}

// Migrate automatically generates the DB schema
func (db DB) Migrate() error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, q := range queries {
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return nil
}
