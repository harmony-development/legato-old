package db

var queries = []string{
	`CREATE TABLE IF NOT EXISTS sessions(
		userid TEXT NOT NULL, 
		session TEXT PRIMARY KEY NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS guilds(
		guildid TEXT PRIMARY KEY NOT NULL,
		guildname TEXT NOT NULL, 
		userid TEXT NOT NULL REFERENCES users(userid),
		picture TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS guildmembers(
		userid TEXT NOT NULL REFERENCES users(userid), 
		guildid TEXT REFERENCES guilds(guildid),
		UNIQUE(userid, guildid)
	);`,
	`CREATE TABLE IF NOT EXISTS invites(
		inviteid TEXT PRIMARY KEY UNIQUE,
		uses INTEGER NOT NULL DEFAULT 0,
		guildid TEXT REFERENCES guilds(guildid)
	);`,
	`CREATE TABLE IF NOT EXISTS channels(
		channelid TEXT PRIMARY KEY UNIQUE, 
		guildid TEXT REFERENCES guilds(guildid), 
		channelname TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS messages(
		messageid TEXT PRIMARY KEY, 
		guildid TEXT REFERENCES guilds(guildid), 
		channelid TEXT REFERENCES channels(channelid), 
		userid TEXT REFERENCES users(userid), 
		createdat INTEGER NOT NULL, 
		message TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS attachments(
		messageid TEXT NOT NULL REFERENCES messages(messageid),
		attachment TEXT NOT NULL
	);`,
}

type Message struct {
	UserID      string
	MessageID   string
	Message     string
	Attachments []string
	CreatedAt   int
}

type Invite struct {
	ID   string
	Uses int
}

// Migrate applies the DB layout to the connected DB
func (db *DB) Migrate() error {
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

// AddSampleData adds sample values to the DB for testing
func (db *DB) AddSampleData() error {
	if err := db.AddGuild(
		"harmony-devs",
		"82ee9c8dc9e165205548b7c3833e7372",
		"Harmony Development",
		"",
	); err != nil {
		return err
	}

	if err := db.AddInvite("join-harmony-dev", "harmony-devs"); err != nil {
		return err
	}

	if err := db.AddMemberToGuild("82ee9c8dc9e165205548b7c3833e7372", "harmony-devs"); err != nil {
		return err
	}

	if err := db.AddMemberToGuild("dadcd6bf8c0338cbfc9aa9c369ea93cc", "harmony-devs"); err != nil {
		return err
	}

	if err := db.AddChannelToGuild("yf8e3bhbp0Bk8UKG", "harmony-devs", "media"); err != nil {
		return err
	}

	return nil
}
