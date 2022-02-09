package store

import "database/sql"

// GetAllCategories ..
func (db *ForumDB) GetAllCategories() ([]string, error) {
	rows, err := db.DB.Query(`SELECT name FROM categories`)
	if err != nil {
		return []string{}, err
	}
	var categories []string
	var category string
	for rows.Next() {
		rows.Scan(&category)
		categories = append(categories, category)
	}
	return categories, nil
}
