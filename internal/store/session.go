package store

import (
	"database/sql"
	"log"
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

// GetUserSession ..
func (db *ForumDB) GetUserSession(session string) (int, error) {
	var userID int
	row := db.DB.QueryRow("SELECT user_ID FROM sessions WHERE Value = ?;", session)
	err := row.Scan(&userID)
	return userID, err
}

// CleanSessions ..
func CleanSessions(db *sql.DB) {
	for {
		_, err := db.Exec("DELETE FROM sessions WHERE expires < ?", time.Now())
		if err != nil {
			log.Println(err)
		}
		time.Sleep(15 * time.Second)
	}
}
