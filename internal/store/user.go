package store

import "forum/internal/app"

// RegisterUser ..
func (db *ForumDB) RegisterUser(user *app.User) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?);`,
		user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec(`
		INSERT INTO posts (title, content, authorID)
		VALUES
				("title4", "content4", 4);
	`)
	if err != nil {
		return err
	}
	return nil
}
