package store

import "forum/internal/app"

// GetPost ..
func (db *ForumDB) GetPost(id int) (*app.Post, error) {
	row := db.DB.QueryRow(`
	SELECT
	*
	FROM posts where ID = ?;
	`, id)
	post := &app.Post{}
	err := row.Scan(&post.ID, &post.Author, &post.Title, &post.Content)
	return post, err
}
