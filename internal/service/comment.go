package service

import (
	"errors"
	"forum/internal/app"
)

// ValidateComment ..
func (s *Service) ValidateComment(comment app.Comment) error {
	if !isValidLen(comment.Content, 1, 128) {
		return errors.New("invalid comment")
	}
	return nil
}

// AddNewCommentToPost ..
func (s *Service) AddNewCommentToPost(comment app.Comment) error {
	err := s.Store.AddNewCommentToPost(comment)
	return err
}
