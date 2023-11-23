package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool // A connection pool for PostgreSQL
	*Queries               // A struct containing all the SQL queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,      // Initialize the connection pool
		Queries:  New(connPool), // Initialize the SQL queries
	}
}
