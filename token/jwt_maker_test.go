package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nmpotential/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	// Create a new JWT maker instance with a random secret key
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	// Generate random user data for testing
	username := util.RandomOwner()
	role := util.DepositorRole
	duration := time.Minute

	// Define issuedAt and expiredAt timestamps based on duration
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// Create a new JWT token with the provided data
	token, payload, err := maker.CreateToken(username, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Verify the created token and check the payload
	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	// Create a new JWT maker instance with a random secret key
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	// Generate an expired token for testing
	token, payload, err := maker.CreateToken(util.RandomOwner(), util.DepositorRole, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Verify the expired token and ensure it returns an error
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// Create a token payload with the provided user data
	payload, err := NewPayload(util.RandomOwner(), util.DepositorRole, time.Minute)
	require.NoError(t, err)

	// Create a new JWT token with the "None" signing method (unsafe)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	// Create a new JWT maker instance with a random secret key
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	// Verify the "None" algorithm token and ensure it returns an error
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
