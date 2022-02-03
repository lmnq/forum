package service

import (
	"errors"
	"forum/internal/app"

	"golang.org/x/crypto/bcrypt"
)

// LoginUser ..
func (s *Service) LoginUser(user *app.User) error {
	hashPW, err := s.Store.GetHashPassword(user.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPW), []byte(user.Password))
	return err
}

// IsValidLoginData ..
func (s *Service) IsValidLoginData(user *app.User) error {
	switch false {
	case isValidPassword(user.Password):
		return errors.New("invalid password")
	case isValidPasswordUni(user.Password):
		return errors.New("invalid uni password")
	case isValidEmail(user.Email):
		return errors.New("invalid email")
	default:
		return nil
	}
}
