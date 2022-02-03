package store

import (
	"database/sql"
	"forum/internal/app"
	"net/http"
	"time"
)

// SetCookie ..
func (db *ForumDB) SetCookie(cookie *http.Cookie, email string) error {
	var userID int
	row := db.DB.QueryRow("SELECT ID FROM users WHERE email=?;", email)
	if err := row.Scan(&userID); err != nil {
		return err
	}
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("DELETE FROM sessions WHERE user_ID=?;", userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = db.DB.Exec("INSERT INTO sessions VALUES (?, ?, ?);", cookie.Value, cookie.Expires, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

// CheckForSession ..
func (db *ForumDB) CheckForSession(user *app.User) error {
	// check
	var value string
	var expires time.Time
	row := db.DB.QueryRow("SELECT Value, Expires FROM sessions WHERE user_ID = ?;", user.ID)
	err := row.Scan(&value, &expires)
	switch err {
	case sql.ErrNoRows:

	}
	return nil
}
