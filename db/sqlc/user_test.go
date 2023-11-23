package db

import (
	"context"
	"testing"
	"time"

	"simplebank/util"

	"github.com/stretchr/testify/require"
)

// createRandomUser generates a random user for testing purposes.
func createRandomUser(t *testing.T) User {
	// Generate a hashed password using a random string.
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	// Create a random user with specified parameters and store it in the database.
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// Verify if the user attributes match the generated random user attributes.
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user // Return the generated random user for further testing
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t) // Create a random user for testing user creation
}

// The following Test functions check various scenarios for updating user information.
// Each test creates a random user, updates specific user information, and verifies the changes.

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t) // Create a random user for testing user retrieval
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	// Verify if the retrieved user's attributes match the original random user's attributes.
	// Also, ensure that time fields are within a second duration accuracy.
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
