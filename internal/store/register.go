package store

import (
	"forum/internal/app"
	"time"
)

// RegisterUser ..
func (db *ForumDB) RegisterUser(user *app.User) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, email, password, created)
		VALUES (?, ?, ?, ?);`,
		user.Username, user.Email, user.HashPassword, time.Now())
	return err
}
