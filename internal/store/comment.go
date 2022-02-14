package store

import "forum/internal/app"

// GetCommentsToPost ..
func (db *ForumDB) GetCommentsToPost(postID int) ([]app.Comment, error) {
	comments := []app.Comment{}
	rows, err := db.DB.Query(`
		SELECT
				comments.ID,
				posts.ID,
				users.username,
				comments.content,
				users.ID
		FROM comments INNER JOIN users ON comments.user_ID = users.ID
		INNER JOIN posts ON comments.post_ID = posts.ID WHERE posts.ID = ?;
		`, postID)
	if err != nil {
		return comments, err
	}
	for rows.Next() {
		comment := app.Comment{}
		rows.Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.Content, &comment.AuthorID)
		votes, rate, err := db.GetVotesToComment(comment.ID, comment.AuthorID)
		if err != nil {
			return comments, err
		}
		comment.Votes = votes
		comment.Rate = rate
		comments = append(comments, comment)
	}
	return comments, nil
}

// AddNewCommentToPost ..
func (db *ForumDB) AddNewCommentToPost(comment app.Comment) error {
	_, err := db.DB.Exec(`
		INSERT INTO
			comments(content, user_ID, post_ID)
		VALUES
			(?, ?, ?);
	`, comment.Content, comment.AuthorID, comment.PostID)
	return err
}
