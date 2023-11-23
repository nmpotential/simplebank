package db

import "context"

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"` // ID of the account to transfer money from
	ToAccountID   int64 `json:"to_account_id"`   // ID of the account to transfer money to
	Amount        int64 `json:"amount"`          // Amount of money to transfer
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`     // Details of the completed transfer
	FromAccount Account  `json:"from_account"` // Information about the account money was transferred from
	ToAccount   Account  `json:"to_account"`   // Information about the account money was transferred to
	FromEntry   Entry    `json:"from_entry"`   // Entry for the transaction in the 'from' account
	ToEntry     Entry    `json:"to_entry"`     // Entry for the transaction in the 'to' account
}

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, adds account entries, and updates accounts' balance within a database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// Execute a database transaction for the transfer operation
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create a transfer record for the transaction
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create an entry in the 'from' account reflecting the deduction of money
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create an entry in the 'to' account reflecting the addition of money
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// Perform the actual balance transfer between accounts within the transaction
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err
	})

	return result, err // Return the transaction result and any encountered error
}

// addMoney performs the adjustment of account balances within a transaction
func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	// Update the balance of the first account
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	// Update the balance of the second account
	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}
