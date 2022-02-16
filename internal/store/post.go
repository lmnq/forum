package store

import (
	"forum/internal/app"
	"time"
)

// GetPost ..
func (db *ForumDB) GetPost(id int) (app.Post, error) {
	row := db.DB.QueryRow("SELECT * FROM posts where ID = ?;", id)
	post := app.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Created, &post.AuthorID)
	if err != nil {
		return post, err
	}
	categories, err := db.GetCategoriesToPost(post.ID)
	if err != nil {
		return post, err
	}
	post.Categories = categories
	comments, err := db.GetCommentsToPost(post.ID)
	if err != nil {
		return post, err
	}
	post.Comments = comments
	votes, rate, err := db.GetVotesToPost(post.ID, post.AuthorID)
	if err != nil {
		return post, err
	}
	post.Votes = votes
	post.Rate = rate
	return post, nil
}

// AddNewPost ..
func (db *ForumDB) AddNewPost(post app.Post) (int, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(`
		INSERT INTO
				posts (title, created, content, author_ID)
		VALUES
				(?, ?, ?, ?);
	`, post.Title, time.Now(), post.Content, post.AuthorID)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// add categories by id
	for _, category := range post.Categories {
		_, err := tx.Exec(`
				INSERT INTO
						posts_categories (post_ID, category_ID)
				VALUES
						(?, ?);
		`, int(postID), category.ID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return int(postID), tx.Commit()
}
