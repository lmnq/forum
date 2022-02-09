package service

import (
	"errors"
	"forum/internal/app"
)

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

// ValidatePostInput ..
func (s *Service) ValidatePostInput(post *app.Post, categories []string) error {

	switch false {
	case isValidLen(post.Title, 2, 64):
		return errors.New("invalid title")
	case isValidLen(post.Content, 2, 512):
		return errors.New("invalid content")
	case len(post.Categories) > 0:
		return errors.New("no category")
	}

	checked := make(map[string]bool)

	for _, v := range post.Categories {
		for _, k := range categories {
			if v == k {
				if checked[v] {
					return errors.New("repeating category name")
				}
				checked[v] = true
				break
			}
		}
		if !checked[v] {
			return errors.New("invalid category name")
		}
	}

	return nil
}

// AddNewPost ..
func (s *Service) AddNewPost(post *app.Post) (int, error) {
	
}