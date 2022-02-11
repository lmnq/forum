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
