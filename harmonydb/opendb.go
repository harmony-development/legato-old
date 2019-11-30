package harmonydb

import (
	"bufio"
	"database/sql"
	"github.com/kataras/golog"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func OpenDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "./harmony.db")
	file, err := os.Open("./initialize-harmony.sql")
	if err != nil {
		golog.Fatalf("Harmony was unable to read the initialization file! The server cannot continue! Reason : %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err = database.Exec(scanner.Text())
		if err != nil {
			golog.Fatalf("Harmony was not able to initialize the database! The server cannot continue! Reason : %v", err)
		}
	}
	return database
}

var DBInst *sql.DB
