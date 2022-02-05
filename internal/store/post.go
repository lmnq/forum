package store

import "forum/internal/app"

// GetPost ..
func (db *ForumDB) GetPost(id int) (*app.Post, error) {
	row := db.DB.QueryRow("SELECT * FROM posts where ID = ?;", id)
	post := &app.Post{}
	err := row.Scan(&post.ID, &post.Author, &post.Title, &post.Content)
	return post, err
}

// GetCommentsToPost ..
func (db *ForumDB) GetCommentsToPost(post *app.Post) ([]*app.Comment, error) {
	comments := []*app.Comment{}
	rows, err := db.DB.Query(`
		SELECT
				comments.ID,
				posts.ID,
				users.username,
				comments.content
		FROM comments INNER JOIN users ON comments.user_ID = users.ID
		INNER JOIN posts ON comments.post_ID = posts.ID WHERE posts.ID = ?;
		`, post.ID)
	if err != nil {
		return comments, err
	}
	for rows.Next() {
		comment := &app.Comment{}
		rows.Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.Content)
		comments = append(comments, comment)
	}
	return comments, nil
}

// GetCategoriesToPost ..
func (db *ForumDB) GetCategoriesToPost(postID int) ([]string, error) {
	categories := []string{}
	rows, err := db.DB.Query(`
		SELECT categories.name
		FROM posts_categories AS pc INNER JOIN posts ON pc.post_ID = posts.ID
		INNER JOIN categories AS c ON pc.category_ID = c.ID WHERE posts.ID = ?;
	`, postID)
	if err != nil {
		return categories, err
	}
	for rows.Next() {
		var category string
		rows.Scan(&category)
		categories = append(categories, category)
	}
	return categories, nil
}
