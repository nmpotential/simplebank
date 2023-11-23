package db

import "context"

// CreateUserTxParams contains parameters required for creating a user transaction.
type CreateUserTxParams struct {
	CreateUserParams CreateUserParams      // Parameters for creating a user
	AfterCreate      func(user User) error // Function to be executed after user creation
}

// CreateUserTxResult contains the result of a create user transaction.
type CreateUserTxResult struct {
	User User // User created as a result of the transaction
}

// CreateUserTx performs a transaction to create a user with specified parameters.
func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	// Execute a transaction within the store's context.
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create a user using the provided CreateUserParams.
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		// Execute the function AfterCreate with the created user as a parameter.
		return arg.AfterCreate(result.User)
	})

	return result, err // Return the transaction result and any error encountered
}
