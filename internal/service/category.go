package service

// GetAllCategories ..
func (s *Service) GetAllCategories() ([]string, error) {
	categories, err := s.Store.GetAllCategories()
	return categories, err
}
