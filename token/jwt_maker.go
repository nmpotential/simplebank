package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	// Check if the secret key meets the minimum size requirement
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	// Generate a new token payload
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	// Create a JWT token with the specified signing method and payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// Sign the JWT token with the secret key
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// Define a function to retrieve the key used to verify the token signature
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Ensure the token signing method is HMAC-based
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		// Return the secret key as the verification key
		return []byte(maker.secretKey), nil
	}

	// Parse and validate the JWT token with the keyFunc and expected claims
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		// Handle token validation errors
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	// Extract and validate the claims from the token
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
