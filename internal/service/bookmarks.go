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

// GetProfilePosts ..
func (s *Service) GetProfilePosts(profileID int) ([]app.Post, error) {
	posts, err := s.Store.GetProfilePosts(profileID)
	if err != nil && err != sql.ErrNoRows {
		return posts, err
	}
	return posts, nil
}