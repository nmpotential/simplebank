// Package token manages PASETO tokens

package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305" // Import for cryptographic functionalities
	"github.com/o1egl/paseto"          // Import for PASETO tokens
)

// PasetoMaker is a struct for creating PASETO tokens
type PasetoMaker struct {
	paseto       *paseto.V2 // Instance of Paseto V2
	symmetricKey []byte     // Key used for encryption and decryption
}

// NewPasetoMaker creates a new PasetoMaker instance
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	// Validate key size
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	// Initialize a PasetoMaker instance
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),       // Create a new instance of Paseto V2
		symmetricKey: []byte(symmetricKey), // Store the symmetric key for encryption/decryption
	}

	return maker, nil
}

// CreateToken generates a PASETO token for a given username, role, and duration
func (maker *PasetoMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	// Generate a new payload for the token
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	// Encrypt the payload to create a token
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken checks the validity of a provided PASETO token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	// Decrypt the token to extract the payload
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken // Return error for an invalid token
	}

	// Check if the token payload is valid
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
