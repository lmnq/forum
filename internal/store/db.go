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
	SELECT * FROM posts
	`)
	if err != nil {
		return posts, err
	}
	// var id, authorid int
	// var title, content string
	for rows.Next() {
		post := &app.Post{}
		// rows.Scan(&id, &title, &content, &authorid)
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
		posts = append(posts, post)
	}
	return posts, nil
}
