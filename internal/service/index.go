package service

import "forum/internal/app"

// GetAllPosts ..
func (s *Service) GetAllPosts(userID int) ([]*app.Post, error) {
	posts, err := s.Store.GetAllPosts(userID)
	return posts, err
}
