package handlers

import "forum/internal/service"

// Forum ..
type Forum struct {
	Service *service.Service
}

// NewForum ..
func NewForum() (*Forum, error) {
	srv, err := service.NewService()
	if err != nil {
		return nil, err
	}
	frm := &Forum{
		Service: srv,
	}
	return frm, nil
}