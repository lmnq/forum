package store

import "forum/internal/app"

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
