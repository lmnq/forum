package store

import "forum/internal/app"

// RegisterUser ..
func (db *ForumDB) RegisterUser(user *app.User) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?);`,
		user.Username, user.Email, user.HashPassword)
	return err
}
