package service

import (
	"errors"
	"fmt"
	"forum/internal/app"
)

// GetPost ..
func (s *Service) GetPost(id int) (app.Post, error) {
	post, err := s.Store.GetPost(id)
	if err != nil {
		fmt.Println("1111")
		return post, err
	}
	categories, err := s.Store.GetCategoriesToPost(post.ID)
	if err != nil {
		fmt.Println("2222")
		return post, err
	}
	post.Categories = categories
	comments, err := s.Store.GetCommentsToPost(post.ID)
	if err != nil {
		fmt.Println("333333")
		return post, err
	}
	post.Comments = comments
	votes, rate, err := s.Store.GetVotesToPost(post.ID, post.AuthorID)
	if err != nil {
		fmt.Println("444444")
		return post, err
	}
	post.Votes = votes
	post.Rate = rate
	return post, nil
}

// ValidatePostInput ..
func (s *Service) ValidatePostInput(post app.Post, categories []app.Category) error {

	switch false {
	case isValidLen(post.Title, 2, 64):
		return errors.New("invalid title")
	case isValidLen(post.Content, 2, 512):
		return errors.New("invalid content")
	case len(post.Categories) > 0:
		return errors.New("no category")
	case len(post.Categories) <= 4:
		return errors.New("too many categories")
	}

	checked := make(map[int]bool)

	for _, v := range post.Categories {
		for _, k := range categories {
			if v.ID == k.ID {
				if checked[v.ID] {
					return errors.New("repeating category name")
				}
				checked[v.ID] = true
				break
			}
		}
		if !checked[v.ID] {
			return errors.New("invalid category name")
		}
	}

	return nil
}

// AddNewPost ..
func (s *Service) AddNewPost(post app.Post) (int, error) {
	postID, err := s.Store.AddNewPost(post)
	return postID, err
}
