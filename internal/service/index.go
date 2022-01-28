package service

import "forum/internal/app"

// GetAllPosts ..
func (s *Service) GetAllPosts() ([]*app.Post, error) {
	posts, err := s.Store.GetAllPosts()
	return posts, err
}
