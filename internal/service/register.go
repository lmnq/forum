package service

import (
	"errors"
	"forum/internal/app"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// IsValidRegisterData ..
func (s *Service) IsValidRegisterData(user *app.User) error {
	switch false {
	case isValidPassword(user.Password):
		return errors.New("invalid password")
	case isValidPasswordUni(user.Password):
		return errors.New("invalid uni password")
	case isValidUsername(user.Username):
		return errors.New("invalid username")
	case isValidEmail(user.Email):
		return errors.New("invalid email")
	default:
		return nil
	}
}

// RegisterUser ..
func (s *Service) RegisterUser(user *app.User) error {
	// hash password
	hashPW, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashPassword = string(hashPW)
	err = s.Store.RegisterUser(user)
	return err
}

// SetCookie ..
func (s *Service) SetCookie(email string) (*http.Cookie, error) {
	uid := uuid.NewV4().String()
	expiration := time.Now().Add(4 * time.Hour)
	cookie := &http.Cookie{
		Name:    "session",
		Value:   uid,
		Expires: expiration,
	}
	err := s.Store.SetCookie(cookie, email)
	return cookie, err
}
