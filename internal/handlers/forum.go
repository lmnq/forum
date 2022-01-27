package handlers

import "forum/internal/service"

// Forum ..
type Forum struct {
	Service *service.Service
}

// NewForum ..
func NewForum(srv *service.Service) *Forum {
	return &Forum{
		Service: srv,
	}
}