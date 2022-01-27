package service

import "forum/internal/store"

// Service ..
type Service struct {
	Store *store.ForumDB
}

// NewService ..
func NewService(db *store.ForumDB) *Service {
	return &Service{
		Store: db,
	}
}
