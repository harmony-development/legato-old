package harmonydb

import (
	"database/sql"
	"fmt"
	"github.com/kataras/golog"
	_ "github.com/lib/pq"
	"os"
)

var queries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY NOT NULL, 
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
		channelname TEXT
	);`,
	`CREATE TABLE IF NOT EXISTS messages(
		messageid TEXT PRIMARY KEY, 
		guildid TEXT REFERENCES guilds(guildid), 
		channelid TEXT REFERENCES channels(channelid), 
		author TEXT REFERENCES users(id), 
		createdat INTEGER NOT NULL, 
		message TEXT NOT NULL
	);`,
	`INSERT INTO users(id, email, username, avatar, password) VALUES(
		'82ee9c8dc9e165205548b7c3833e7372', 
		'developer@harmonyapp.io', 
		'developer', 
		'', 
		'$2a$10$WHuq8sNHk0ks0JwlpkV36eNmpEvD7r9pqI/F7kB0q0yAUpENzmtne'
	) ON CONFLICT DO NOTHING;`,
	`INSERT INTO guilds(guildid, owner, guildname, picture) VALUES(
		'harmony-devs', 
		'82ee9c8dc9e165205548b7c3833e7372',
		'Harmony Development', 
		'') ON CONFLICT DO NOTHING;`,
	`INSERT INTO invites(inviteid, guildid) VALUES(
		'join-harmony-dev', 
		'harmony-devs')
		ON CONFLICT DO NOTHING;`,
	`INSERT INTO users(id, email, username, avatar, password) VALUES(
		'dadcd6bf8c0338cbfc9aa9c369ea93cc', 
		'developer2@harmonyapp.io', 
		'developer2', 
		'', 
		'$2a$10$yTHVSHmbAAgcIysrJZg/cesPg7o9qpoTGxFgeM/7pQIgOLFjJZPLW') ON CONFLICT DO NOTHING;`,
	`INSERT INTO guildmembers(userid, guildid) VALUES(
		'82ee9c8dc9e165205548b7c3833e7372', 
		'harmony-devs'
	) ON CONFLICT DO NOTHING;`,
	`INSERT INTO guildmembers(userid, guildid) VALUES(
		'dadcd6bf8c0338cbfc9aa9c369ea93cc', 
		'harmony-devs'
	) ON CONFLICT DO NOTHING;`,
	`INSERT INTO channels(channelid, guildid, channelname) VALUES(
		'FNhj3bhbKFBYHUKG', 
		'harmony-devs', 
		'general') ON CONFLICT DO NOTHING;`,
	`INSERT INTO channels(channelid, guildid, channelname) VALUES(
		'MDjSJMfNDFs', 
		'harmony-devs', 
		'bruh') ON CONFLICT DO NOTHING;`,
}

func OpenDB() *sql.DB {
	database, err := sql.Open("postgres", fmt.Sprintf("user=%v password=%v host=%v port=%v sslmode=%v",
		os.Getenv("HARMONY_DB_USER"),
		os.Getenv("HARMONY_DB_PASSWORD"),
		os.Getenv("HARMONY_DB_HOST"),
		os.Getenv("HARMONY_DB_PORT"),
		"disable", ))
	if err != nil {
		golog.Fatalf("Harmony was unable to open the database! Reason : %v", err)
	}
	initTransaction, err := database.Begin()
	if err != nil {
		golog.Fatalf("Error initializing transaction : %v", err)
	}
	for i := range queries {
		_, err := initTransaction.Exec(queries[i])
		if err != nil {
			golog.Fatalf("Harmony was not able to initialize the database! The server cannot continue! Query: \n%v\nReason : %v", queries[i], err)
		}
	}
	err = initTransaction.Commit()
	if err != nil {
		golog.Fatalf("Error running initialization transaction! %v", err)
	}
	return database
}

var DBInst *sql.DB
