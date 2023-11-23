# Target to start a PostgreSQL container with version 16 using Alpine image
postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

# Target to create a database named 'simple_bank' in the PostgreSQL container
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

# Target to drop the 'simple_bank' database from the PostgreSQL container
dropdb:
	docker exec -it postgres16 dropdb simple_bank

# Target to apply all pending migrations to the 'simple_bank' database
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# Target to apply a single pending migration to the 'simple_bank' database
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

# Target to revert all applied migrations from the 'simple_bank' database
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# Target to revert a single applied migration from the 'simple_bank' database
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

# Target to generate Go code using sqlc for interacting with the database
sqlc:
	sqlc generate

# Target to run tests for the Go code with verbose output and coverage information
test:
	go test -v -cover ./...

# Declare the targets as phony (non-file targets) to ensure they always run
.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test
