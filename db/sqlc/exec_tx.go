package db

import (
	"context"
	"fmt"
)

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Begin a transaction
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	// Create a new Queries instance with the transaction
	q := New(tx)

	// Execute the provided function with the Queries instance
	err = fn(q)
	if err != nil {
		// Roll back the transaction if the function returns an error
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// Commit the transaction if the function returns no error
	return tx.Commit(ctx)
}
