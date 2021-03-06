package store

import (
	"forum/internal/app"
	"log"
)

// GetAllCategories ..
func (db *ForumDB) GetAllCategories() ([]app.Category, error) {
	rows, err := db.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		return nil, err
	}
	var categories []app.Category
	category := app.Category{}
	for rows.Next() {
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
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

// GetPostsByCategory ..
func (db *ForumDB) GetPostsByCategory(categoryID, userID int) ([]app.Post, error) {
	posts := []app.Post{}
	rows, err := db.DB.Query(`
		SELECT
				posts.ID,
				title,
				content,
				username
		FROM posts INNER JOIN users ON posts.author_ID = users.ID
		WHERE posts.ID IN (
							SELECT post_ID FROM posts_categories
							WHERE category_ID = ?
						)
		ORDER BY posts.ID DESC;
		`, categoryID)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := app.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		categories, err := db.GetCategoriesToPost(post.ID)
		if err != nil {
			log.Println("2")
			return posts, err
		}
		post.Categories = categories
		votes, rate, err := db.GetVotesToPost(post.ID, userID)
		if err != nil {
			log.Println("3")
			return posts, err
		}
		post.Votes = votes
		post.Rate = rate
		posts = append(posts, post)
	}

	return posts, nil
}

// GetCategoryByID ..
func (db *ForumDB) GetCategoryByID(categoryID int) (app.Category, error) {
	category := app.Category{}
	row := db.DB.QueryRow("SELECT ID, name FROM categories WHERE ID = ?", categoryID)
	err := row.Scan(&category.ID, &category.Name)
	return category, err
}
