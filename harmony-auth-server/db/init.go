package db

import (
	"database/sql"
	"fmt"
	"github.com/kataras/golog"
	_ "github.com/lib/pq"
	"os"
)

var DB sql.DB

var queries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		userid TEXT PRIMARY KEY NOT NULL, 
		email TEXT UNIQUE NOT NULL, 
		username TEXT NOT NULL, 
		avatar TEXT NOT NULL, 
		password TEXT NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS servers(
		userid TEXT NOT NULL REFERENCES users(userid), -- the userid of who owns this entry
		host TEXT PRIMARY KEY NOT NULL -- the host for the harmony instance
	)`,
	`CREATE TABLE IF NOT EXISTS sessions(
		sessionid TEXT PRIMARY KEY NOT NULL, -- session corresponding to a user's identity'
		expiration INTEGER NOT NULL, -- the time when the session expires
		userid TEXT NOT NULL REFERENCES users(userid) -- the owner of the session
	);
	`,
}

func OpenDB() *sql.DB {
	database, err := sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		os.Getenv("HARMONY_AUTH_USER"),
		os.Getenv("HARMONY_AUTH_PASSWORD"),
		"harmony-auth",
		os.Getenv("HARMONY_AUTH_HOST"),
		os.Getenv("HARMONY_AUTH_PORT"),
		"disable", ))
	if err != nil {
		golog.Fatalf("Harmony was unable to open the database! Reason : %v", err)
		return nil
	}
	initTransaction, err := database.Begin()
	if err != nil {
		golog.Fatalf("Error initializing transaction : %v", err)
		return nil
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