package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTransferTx tests the transfer transaction functionality of a simple bank application.
func TestTransferTx(t *testing.T) {
	// Create two random accounts and print their balances.
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// Set the number of concurrent transfer transactions to run to 5 and the amount to transfer to 10.
	n := 5
	amount := int64(10)

	// Create channels for errors and results.
	errs := make(chan error)
	results := make(chan TransferTxResult)

	// Run n concurrent transfer transactions.
	for i := 0; i < n; i++ {
		go func() {
			// Initiate a transfer transaction and send errors and results to respective channels
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// Check the results of each transfer transaction.
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transfer.
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check entries.
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Check accounts.
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// Check balances.
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// Check the final updated balances of the accounts and verify that they are correct.
	updatedAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

// TestTransferTxDeadlock tests the scenario of potential deadlock in TransferTx function
func TestTransferTxDeadlock(t *testing.T) {
	// Creating two random accounts for the test
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// Number of concurrent transfers
	n := 10
	amount := int64(10)
	errs := make(chan error)

	// Initiating concurrent transfer transactions
	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		// For every odd iteration, the transfer direction is swapped to provoke concurrency issues
		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			// Executing the transfer transaction asynchronously
			_, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err
		}()
	}

	// Receiving and checking the errors from concurrent transactions
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// Checking the final updated balance after transfers
	updatedAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	// Ensuring the balances of the accounts remain unchanged after the transactions
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
