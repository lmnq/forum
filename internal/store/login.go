package store

// GetHashPassword ..
func (db *ForumDB) GetHashPassword(email string) (string, error) {
	var hashPW string
	row := db.DB.QueryRow("SELECT password FROM users WHERE email = ?;", email)
	err := row.Scan(&hashPW)
	return hashPW, err
}