package store

import "database/sql"

// ForumDB ..
type ForumDB struct {
	DB *sql.DB
}

// NewDataBase ..
func NewDataBase() (*ForumDB, error) {
	db, err := initDB()
	return &ForumDB{DB: db}, err
}
