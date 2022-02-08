package store

// GetVotesToEntity ..
func (db *ForumDB) GetVotesToEntity(entityname string, entityID, userID int) ([2]int, error) {
	row := db.DB.QueryRow(`
		SELECT ifnull(SUM(status), 0),
			CASE WHEN user_ID = ?
			THEN status ELSE 0
			END as vote
		FROM votes WHERE entity = ? AND entity_ID = ?;
	`, userID, entityname, entityID)
	var numVotes, vote int
	err := row.Scan(&numVotes, &vote)
	return [2]int{numVotes, vote}, err
}
