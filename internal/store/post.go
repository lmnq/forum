package store

import (
	"forum/internal/app"
	"log"
)

// GetPost ..
func (db *ForumDB) GetPost(id int) (app.Post, error) {
	row := db.DB.QueryRow("SELECT * FROM posts where ID = ?;", id)
	post := app.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
	return post, err
}

// AddNewPost ..
func (db *ForumDB) AddNewPost(post app.Post) (int, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("11111")
		log.Println(err)
		return 0, err
	}

	res, err := tx.Exec(`
		INSERT INTO
				posts (title, content, author_ID)
		VALUES
				(?, ?, ?);
	`, post.Title, post.Content, post.AuthorID)

	if err != nil {
		log.Println("2222")
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		log.Println("333")
		log.Println(err)
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
			log.Println("444")
			log.Println(err)
			tx.Rollback()
			return 0, err
		}
	}

	return int(postID), tx.Commit()
}
