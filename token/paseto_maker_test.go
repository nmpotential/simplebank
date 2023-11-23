// Package token provides testing for Paseto tokens

package token

import (
	"testing"
	"time"

	"github.com/nmpotential/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	// Create a new Paseto token maker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Generate random username, set role as Depositor, and set a duration for the token
	username := util.RandomOwner()
	role := util.DepositorRole
	duration := time.Minute

	// Record the time of token issuance and expiration
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// Create a token and its payload
	token, payload, err := maker.CreateToken(username, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Verify the token and payload
	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Validate payload data
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	// Create a new Paseto token maker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Create an expired token
	token, payload, err := maker.CreateToken(util.RandomOwner(), util.DepositorRole, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Verify an expired token
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
