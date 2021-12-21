package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"login"	NVARCHAR(64) NOT NULL UNIQUE,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if err != nil {
		log.Println("ff")
		log.Fatal(err.Error())
	}
	statement.Exec()

	stmt, err := db.Prepare(`
		INSERT INTO users (login) VALUES (?)
	`)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec("test@gmail.com")
	if err != nil {
		log.Println(err.Error())
	}
	stmt, err = db.Prepare(`
		INSERT INTO users (login) VALUES (?)
	`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec("test2@gmail.com")
}
