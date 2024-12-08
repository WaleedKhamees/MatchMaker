package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email address")
	}
	return nil
}

func ValidateUsername(username string) error {
	if len(username) < 4 || len(username) > 20 {
		return errors.New("Username must be between 4 and 20 characters")
	}
	userRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !userRegex.MatchString(username) {
		return errors.New("username must contain only letters, numbers and underscores")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}
	return nil
}
