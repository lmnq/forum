package store

import (
	"database/sql"

	// sqlite3 ..
	_ "github.com/mattn/go-sqlite3"
)

// ForumDB ..
type ForumDB struct {
	DB *sql.DB
}

// InitDB ..
func InitDB() (*ForumDB, error) {
	db, err := sql.Open("sqlite3", "./forum.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = initTables(db); err != nil {
		return nil, err
	}
	if err = insertData(db); err != nil {
		return nil, err
	}
	// if err = deleteUser(db); err != nil {
	// 	return nil, err
	// }

	forumDB := &ForumDB{
		DB: db,
	}
	return forumDB, nil
}

func initTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS "users" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"username"	TEXT NOT NULL,
			"email"	TEXT NOT NULL UNIQUE,
			"password"	TEXT NOT NULL,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE  IF NOT EXISTS "posts" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"title"	TEXT NOT NULL DEFAULT 'untitled',
			"content"	TEXT NOT NULL,
			"authorID"	INTEGER NOT NULL,
			FOREIGN KEY("authorID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "comments" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"content"	TEXT NOT NULL,
			"user_ID"	INTEGER NOT NULL,
			"post_ID"	INTEGER NOT NULL,
			FOREIGN KEY("user_ID") REFERENCES "users"("ID") ON DELETE CASCADE,
			FOREIGN KEY("post_ID") REFERENCES "posts"("ID") ON DELETE CASCADE,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "categories" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT NOT NULL UNIQUE,
			PRIMARY KEY("ID" AUTOINCREMENT)
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "posts_categories" (
			"post_ID"	INTEGER NOT NULL,
			"category_ID"	INTEGER NOT NULL,
			PRIMARY KEY("post_ID","category_ID"),
			UNIQUE("post_ID", "category_ID")
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func insertData(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO users (username, email, password)
		VALUES
				("user1", "user1@gmail.com", "user1password"),
				("user2", "user2@gmail.com", "user2password"),
				("user3", "user3@gmail.com", "user3password");
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO posts (title, content, authorID)
		VALUES
				("title1", "content1", 1),
				("title2", "content2", 2),
				("title3", "content3", 3);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO comments (content, user_ID, post_ID)
		VALUES
				("commentary 1", 1, 1),
				("commentary 2", 2, 2),
				("commentary 3", 3, 3);
	`)
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(db *sql.DB) error {
	_, err := db.Exec(`
		DELETE FROM users
		WHERE ID = 1;
	`)
	if err != nil {
		return err
	}

	return nil
}
