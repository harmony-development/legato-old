package db

// AddUser inserts a new user to the DB
func (db *DB) AddUser(userID string, email string, username string, password string, avatar string) error {
	_, err := db.Exec(`INSERT INTO 
    	users(userid, email, username, password, avatar) 
    	VALUES($1, $2, $3, $4, $5)
    	ON CONFLICT DO NOTHING
	`, userID, email, username, password, avatar)
	return err
}

// AddGuild inserts a new guild to the DB
func (db *DB) AddGuild(guildID string, owner string, guildName string, picture string) error {
	_, err := db.Exec(`INSERT INTO 
    	guilds(guildid, owner, guildname, picture) 
    	VALUES($1, $2, $3, $4)
    	ON CONFLICT DO NOTHING;
	`, guildID, owner, guildName, picture)
	return err
}

// AddInvite inserts a new invite to the DB
func (db *DB) AddInvite(inviteID string, guildID string) error {
	_, err := db.Exec(`INSERT INTO 
    invites(inviteid, guildid) 
    VALUES($1, $2)
	ON CONFLICT DO NOTHING;
	`, guildID, inviteID, guildID)
	return err
}

// AddMemberToGuild adds a new member to a guild
func (db *DB) AddMemberToGuild(userID string, guildID string) error {
	_, err := db.Exec(`
	INSERT INTO guildmembers(userid, guildid) 
	VALUES($1, $2) 
	ON CONFLICT DO NOTHING;`, userID, guildID)
	return err
}

// AddChannelToGuild adds a new channel to a guild
func (db *DB) AddChannelToGuild(channelID string, guildID string, channelName string) error {
	_, err := db.Exec(`
	INSERT INTO channels(channelid, guildid, channelname)
	VALUES($1, $2, $3) 
	ON CONFLICT DO NOTHING;`, channelID, guildID, channelName)
	return err
}