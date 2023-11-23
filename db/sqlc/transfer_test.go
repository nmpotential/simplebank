// Package db provides the database operations for the simple bank application.
package db

import (
	"context"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// createRandomTransfer creates a random transfer between two accounts and returns the transfer object.
func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	// Create a CreateTransferParams object with the FromAccountID, ToAccountID, and Amount fields set to the corresponding account IDs and a random amount.
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	// Call the CreateTransfer method of the testStore object with the context.Background() and arg as parameters.
	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	// Assert that the FromAccountID, ToAccountID, and Amount fields of the returned transfer object are equal to the corresponding fields of the CreateTransferParams object.
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	// Assert that the ID and CreatedAt fields of the returned transfer object are not zero.
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

// TestCreateTransfer tests the CreateTransfer method.
func TestCreateTransfer(t *testing.T) {
	// Create two random accounts using the createRandomAccount function.
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Call the createRandomTransfer function with these two accounts as parameters.
	createRandomTransfer(t, account1, account2)
}

// TestGetTransfer tests the GetTransfer method.
func TestGetTransfer(t *testing.T) {
	// Create two random accounts using the createRandomAccount function.
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Create a random transfer between these two accounts using the createRandomTransfer function.
	transfer1 := createRandomTransfer(t, account1, account2)

	// Call the GetTransfer method of the testStore object with the context.Background() and the ID of the created transfer as parameters.
	transfer2, err := testStore.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	// Assert that the fields of the returned transfer object are equal to the corresponding fields of the created transfer object.
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	// Assert that the CreatedAt field of the returned transfer object is within one second of the CreatedAt field of the created transfer object.
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

// TestListTransfer tests the ListTransfer method.
func TestListTransfer(t *testing.T) {
	// Create two random accounts using the createRandomAccount function.
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Create five random transfers between these two accounts and five random transfers between the second account and the first account using the createRandomTransfer function.
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	// Call the ListTransfer method of the testStore object with a ListTransfersParams object with the FromAccountID and ToAccountID fields set to the ID of the first account, Limit set to 5, and Offset set to 5.
	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	// Assert that each transfer object in the returned transfers slice has a non-empty value and either the FromAccountID or ToAccountID field is equal to the ID of the first account.
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
