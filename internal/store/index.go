package store

import (
	"forum/internal/app"
)

// GetAllPosts ..
func (db *ForumDB) GetAllPosts(userID int) ([]app.Post, error) {
	posts := []app.Post{}
	rows, err := db.DB.Query(`
	SELECT
			posts.ID,
			title,
			posts.created,
			content,
			username,
			author_ID
	FROM posts INNER JOIN users ON posts.author_ID = users.ID;
	`)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := app.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Created, &post.Content, &post.Author, &post.AuthorID)
		categories, err := db.GetCategoriesToPost(post.ID)
		if err != nil {
			return posts, err
		}
		post.Categories = categories
		votes, rate, err := db.GetVotesToPost(post.ID, userID)
		if err != nil {
			return posts, err
		}
		post.Votes = votes
		post.Rate = rate
		posts = append(posts, post)
	}
	return posts, nil
}
