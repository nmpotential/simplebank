// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context) (Account, error)
	DeleteAccount(ctx context.Context) error
	GetAccount(ctx context.Context) (Account, error)
	ListAccounts(ctx context.Context) ([]Account, error)
	// LIMIT 1;
	// OFFSET 2;
	UpdateAccount(ctx context.Context) (Account, error)
}

var _ Querier = (*Queries)(nil)
