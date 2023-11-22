package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:   "Ben",
		Balance: 100,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

}
