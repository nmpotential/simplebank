package db

import (
	"context"
	"log"
	"os"
	"simplebank/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

// testStore is a global variable that holds a Store instance for testing
var testStore Store

// TestMain initializes a testStore variable with a connection to the database
func TestMain(m *testing.M) {
	// Load the configuration file
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Create a connection pool with the configuration's DBSource
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Create a new store with the connection pool
	testStore = NewStore(connPool)

	// Run the tests using m.Run()
	os.Exit(m.Run())
}
