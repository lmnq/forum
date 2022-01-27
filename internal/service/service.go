package service

import "forum/internal/store"

// Service ..
type Service struct {
	Store *store.ForumDB
}

// NewService ..
func NewService() (*Service, error) {
	db, err := store.InitDB()
	if err != nil {
		return nil, err
	}
	srv := &Service{
		Store: db,
	}
	return srv, nil
}
