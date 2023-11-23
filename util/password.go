package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash from the provided password with the default cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Return an error if there's a failure in hashing the password
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	// Return the hashed password as a string
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password matches the hashed password
func CheckPassword(password string, hashedPassword string) error {
	// Compare the provided password with the hashed password using bcrypt
	// This function returns an error if the passwords do not match
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
