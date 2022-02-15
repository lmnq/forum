package service

import (
	"database/sql"
	"forum/internal/app"
)

// GetBookmarkedPosts ..
func (s *Service) GetBookmarkedPosts(userID int) ([]app.Post, error) {
	posts, err := s.Store.GetBookmarkedPosts(userID)
	if err != nil && err != sql.ErrNoRows {
		return posts, err
	}
	return posts, nil
}
