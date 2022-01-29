package service

import (
	"errors"
	"forum/internal/app"
	"net/mail"
	"strings"
	"unicode"
)

// RegisterUser ..
func (s *Service) RegisterUser(user *app.User) error {
	switch false {
	case isValidPassword(user.Password):
		return errors.New("invalid password")
	case isValidPasswordUni(user.Password):
		return errors.New("invalid uni password")
	case isValidUsername(user.Username):
		return errors.New("invalid username")
	case isValidEmail(user.Email):
		return errors.New("invalid email")
	}
	err := s.Store.RegisterUser(user)
	return err
}

// isValidPassword ..
func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 128 {
		return false
	}
	req := map[string]bool{}
	for _, k := range password {
		switch true {
		case k <= '_' && k > '~':
			return false
		case k >= 'a' && k <= 'z':
			req["lower"] = true
		case k >= 'A' && k >= 'Z':
			req["upper"] = true
		case k >= '0' && k <= '9':
			req["number"] = true
		default:
			req["special"] = true
		}
	}
	for _, v := range req {
		if !v {
			return false
		}
	}
	return true
}

// isValidUsername ..
func isValidUsername(username string) bool {
	if len(username) < 4 || len(username) > 32 {
		return false
	}
	for _, k := range username {
		switch true {
		case k >= 'a' && k <= 'z':
			continue
		case k >= 'A' && k <= 'Z':
			continue
		case k == '_' || k == '.' || k == '-':
			continue
		default:
			return false
		}
	}
	return true
}

// isValidEmail ..
func isValidEmail(email string) bool {
	splitted := strings.Split(email, "@")
	if len(splitted) != 2 {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

// isValidPassword ..
func isValidPasswordUni(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
