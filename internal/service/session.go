package service

// GetUserSession ..
func (s *Service) GetUserSession(session string) (int, error) {
	userID, err := s.Store.GetUserSession(session)
	return userID, err
}
