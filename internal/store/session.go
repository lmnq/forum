package store

import (
	"database/sql"
	"forum/internal/app"
	"net/http"
	"time"
)

// SetCookie ..
func (db *ForumDB) SetCookie(cookie *http.Cookie, email string) error {
	userID := 42 //get user ID
	_, err := db.DB.Exec(`
		INSERT INTO sessions
			VALUES (?, ?, ?);
	`, cookie.Value, cookie.Expires, userID)
	if err != nil {
		return err
	}
	// insert
	return nil
}

// CheckForSession ..
func (db *ForumDB) CheckForSession(user *app.User) error {
	// check
	var value string
	var expires time.Time
	row := db.DB.QueryRow("SELECT Value, Expires FROM sessions WHERE user_ID = ?", user.ID)
	err := row.Scan(&value, &expires)
	switch err {
	case sql.ErrNoRows:

	}
	return nil
}
