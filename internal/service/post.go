package service

import "forum/internal/app"

// GetCommentsToPost ..
func (s *Service) GetCommentsToPost(post *app.Post) ([]*app.Comment, error) {
	comments, err := s.Store.GetCommentsToPost(post)
	return comments, err
}

// GetPost ..
func (s *Service) GetPost(id int) (*app.Post, error) {
	post, err := s.Store.GetPost(id)
	return post, err
}
