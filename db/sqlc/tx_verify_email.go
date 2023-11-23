package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

// VerifyEmailTxParams contains parameters required for verifying an email.
type VerifyEmailTxParams struct {
	EmailId    int64  // ID of the email to be verified
	SecretCode string // Secret code associated with the email verification process
}

// VerifyEmailTxResult contains the result of the email verification transaction.
type VerifyEmailTxResult struct {
	User        User        // Updated user information after email verification
	VerifyEmail VerifyEmail // Updated email verification details
}

// VerifyEmailTx performs a transaction to verify an email address.
func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	// Execute a database transaction for verifying the email
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Update the email verification details
		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		// Update the user information after successful email verification
		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		})
		return err
	})

	return result, err // Return the transaction result and any encountered error
}
