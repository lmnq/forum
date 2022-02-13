package store

// GetVotesToEntity ..
func (db *ForumDB) GetVotesToEntity(entityname string, entityID, userID int) (int, int, error) {
	row := db.DB.QueryRow(`
		SELECT ifnull(SUM(status), 0),
			CASE WHEN user_ID = ?
			THEN status ELSE 0
			END as vote
		FROM votes WHERE entity = ? AND entity_ID = ?;
	`, userID, entityname, entityID)
	var votes, rate int
	err := row.Scan(&votes, &rate)
	return votes, rate, err
}
