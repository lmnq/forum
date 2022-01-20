package store

import (
	"database/sql"
	"forum/internal/app"
)

// ForumDB ..
type ForumDB struct {
	DB *sql.DB
}

// DataBase ..
type DataBase interface {
	// fill with db methods (crud)
	GetAllPosts() ([]*app.Post, error)
}

// GetAllPosts ..
func (db *ForumDB) GetAllPosts() ([]*app.Post, error) {
	posts := []*app.Post{}
	rows, err := db.DB.Query(`
	SELECT
			posts.ID,
			title,
			content,
			username
	FROM posts INNER JOIN users ON posts.authorID = users.ID;
	`)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := &app.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		posts = append(posts, post)
	}
	return posts, nil
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
	FROM comments INNER JOIN users ON comments.user_ID = users.ID INNER JOIN posts ON comments.post_ID = posts.ID WHERE posts.ID = ?;
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
