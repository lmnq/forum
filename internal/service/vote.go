package service

// VotePost ..
func (s *Service) VotePost(postID, userID int, rate int) error {
	err := s.Store.VotePost(postID, userID, rate)
	return err
}

// VoteComment ..
func (s *Service) VoteComment(commentID, userID int, rate int) error {
	err := s.Store.VoteComment(commentID, userID, rate)
	return err
}
