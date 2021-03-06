package service

import (
	"database/sql"
	"errors"
	"forum/internal/app"
	"strconv"
)

// GetAllCategories ..
func (s *Service) GetAllCategories() ([]app.Category, error) {
	categories, err := s.Store.GetAllCategories()
	return categories, err
}

// GetCategoriesFromInput ..
func (s *Service) GetCategoriesFromInput(input []string) ([]app.Category, error) {
	unique := make(map[string]bool)
	categories := []app.Category{}
	for _, v := range input {
		id, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		if unique[v] {
			return nil, errors.New("repeating category")
		}
		categories = append(categories, app.Category{ID: id})
		unique[v] = true
	}
	return categories, nil
}

// GetPostsByCategory ..
func (s *Service) GetPostsByCategory(categoryID, userID int) ([]app.Post, error) {
	posts, err := s.Store.GetPostsByCategory(categoryID, userID)
	if err != nil && err != sql.ErrNoRows {
		return posts, err
	}
	return posts, nil
}

// GetCategoryByID ..
func (s *Service) GetCategoryByID(categoryID int) (app.Category, error) {
	category, err := s.Store.GetCategoryByID(categoryID)
	return category, err
}