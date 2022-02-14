package store

import "database/sql"

// GetVotesToPost ..
func (db *ForumDB) GetVotesToPost(postID, userID int) (int, int, error) {
	row := db.DB.QueryRow(`
		SELECT ifnull(SUM(rate), 0),
			CASE WHEN user_ID = ?
			THEN rate ELSE 0
			END as vote
		FROM post_votes WHERE post_ID = ?;
	`, userID, postID)
	var votes, rate int
	err := row.Scan(&votes, &rate)
	return votes, rate, err
}

// GetVotesToComment ..
func (db *ForumDB) GetVotesToComment(commentID, userID int) (int, int, error) {
	row := db.DB.QueryRow(`
		SELECT ifnull(SUM(rate), 0),
			CASE WHEN user_ID = ?
			THEN rate ELSE 0
			END as vote
		FROM comment_votes WHERE comment_ID = ?;
	`, userID, commentID)
	var votes, rate int
	err := row.Scan(&votes, &rate)
	return votes, rate, err
}

// VotePost ..
func (db *ForumDB) VotePost(postID, userID int, rate int) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(`
			SELECT ID, rate
			FROM post_votes
			WHERE user_ID = ? AND post_ID = ?;
	`, userID, postID)

	var currVoteID, currVoteRate int
	err = row.Scan(&currVoteID, &currVoteRate)

	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return err
	}

	if err == nil {
		_, errin := tx.Exec(`DELETE FROM post_votes WHERE ID = ?`, currVoteID)
		if errin != nil {
			tx.Rollback()
			return errin
		}
	}

	if err == sql.ErrNoRows || currVoteRate != rate {
		_, errin := tx.Exec(`
					INSERT INTO
							post_votes(rate, user_ID, post_ID)
					VALUES
							(?, ?, ?);
		`, rate, userID, postID)
		if errin != nil {
			tx.Rollback()
			return errin
		}
	}

	return tx.Commit()
}

// VoteComment ..
func (db *ForumDB) VoteComment(commentID, userID int, rate int) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(`
			SELECT ID, rate
			FROM comment_votes
			WHERE user_ID = ? AND comment_ID = ?;
	`, userID, commentID)

	var currVoteID, currVoteRate int
	err = row.Scan(&currVoteID, &currVoteRate)

	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return err
	}

	if err == nil {
		_, errin := tx.Exec(`DELETE FROM comment_votes WHERE ID = ?`, currVoteID)
		if errin != nil {
			tx.Rollback()
			return errin
		}
	}

	if err == sql.ErrNoRows || currVoteRate != rate {
		_, errin := tx.Exec(`
					INSERT INTO
							comment_votes(rate, user_ID, comment_ID)
					VALUES
							(?, ?, ?);
		`, rate, userID, commentID)
		if errin != nil {
			tx.Rollback()
			return errin
		}
	}

	return tx.Commit()
}
