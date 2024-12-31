package utils

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt algorithm.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword compares the given password with the hashed password.
// Returns true if the passwords match, false otherwise.
func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Print(err)
	return err == nil
}

func ValidateEmail(email string) error {
	// Basic email validation regex
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword checks if the provided password meets the requirements:
// length between 8 and 30, contains uppercase, lowercase, number, and special character.
func ValidatePassword(password string) error {
	// Check the length of the password
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password must be between 8 and 30 characters long")
	}

	// Regular expressions to check for different character types
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`).MatchString // Adjust the special characters as per your requirements
	)

	if !hasUpper(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber(password) {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
