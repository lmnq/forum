package store

import "forum/internal/app"

// GetPost ..
func (db *ForumDB) GetPost(id int) (app.Post, error) {
	row := db.DB.QueryRow("SELECT * FROM posts where ID = ?;", id)
	post := app.Post{}
	err := row.Scan(&post.ID, &post.Author, &post.Title, &post.Content)
	if err != nil {
		return post, err
	}
	categories, err := db.GetCategoriesToPost(post.ID)
	if err != nil {
		return post, err
	}
	post.Categories = categories
	return post, nil
}

// GetCommentsToPost ..
func (db *ForumDB) GetCommentsToPost(post app.Post) ([]app.Comment, error) {
	comments := []app.Comment{}
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
		comment := app.Comment{}
		rows.Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.Content)
		comments = append(comments, comment)
	}
	return comments, nil
}

// GetCategoriesToPost ..
func (db *ForumDB) GetCategoriesToPost(postID int) ([]app.Category, error) {
	categories := []app.Category{}
	rows, err := db.DB.Query(`
		SELECT c.ID, c.name
		FROM posts_categories AS pc INNER JOIN posts ON pc.post_ID = posts.ID
		INNER JOIN categories AS c ON pc.category_ID = c.ID WHERE posts.ID = ?;
	`, postID)
	if err != nil {
		return categories, err
	}
	for rows.Next() {
		category := app.Category{}
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}
	return categories, nil
}

// AddNewPost ..
func (db *ForumDB) AddNewPost(post app.Post) (int, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(`
		INSERT INTO
				posts (title, content, author_ID)
		VALUES
				(?, ?, ?);
	`, post.Title, post.Content, post.Author)

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
		`, postID, category.ID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return int(postID), tx.Commit()
}
