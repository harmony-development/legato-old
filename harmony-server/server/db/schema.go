package db

var queries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		userid TEXT PRIMARY KEY NOT NULL, 
		email TEXT UNIQUE NOT NULL, 
		username TEXT NOT NULL, 
		avatar TEXT NOT NULL, 
		password TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS guilds(
		guildid TEXT PRIMARY KEY NOT NULL,
		guildname TEXT NOT NULL, 
		owner TEXT NOT NULL REFERENCES users(id),
		picture TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS guildmembers(
		userid TEXT NOT NULL REFERENCES users(id), 
		guildid TEXT REFERENCES guilds(guildid),
		UNIQUE(userid, guildid)
	);`,
	`CREATE TABLE IF NOT EXISTS invites(
		inviteid TEXT PRIMARY KEY UNIQUE,
		invitecount INTEGER NOT NULL DEFAULT 0,
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
		author TEXT REFERENCES users(id), 
		createdat INTEGER NOT NULL, 
		message TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS attachments(
		messageid TEXT NOT NULL REFERENCES messages(messageid),
		attachment TEXT NOT NULL
	);`,
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
	if err := db.AddUser(
		"82ee9c8dc9e165205548b7c3833e7372",
		"developer@harmonyapp.io",
		"Developer",
		"$2a$10$WHuq8sNHk0ks0JwlpkV36eNmpEvD7r9pqI/F7kB0q0yAUpENzmtne",
		"",
	); err != nil {
		return err
	}

	if err := db.AddUser(
		"dadcd6bf8c0338cbfc9aa9c369ea93cc",
		"developer2@harmonyapp.io",
		"Developer #2",
		"$2a$10$yTHVSHmbAAgcIysrJZg/cesPg7o9qpoTGxFgeM/7pQIgOLFjJZPLW",
		"",
	); err != nil {
		return err
	}

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

	if err := db.AddChannelToGuild("FNhj3bhbKFBYHUKG", "harmony-devs", "general"); err != nil {
		return err
	}

	if err := db.AddChannelToGuild("yf8e3bhbp0Bk8UKG", "harmony-devs", "media"); err != nil {
		return err
	}
}
