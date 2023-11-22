postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=498040 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root bank_service

dropdb:
	docker exec -it postgres16 dropdb bank_service

migrateup:
	migrate -path db/migration -database "postgresql://root:498040@localhost:5432/bank_service?sslmode=disable" -verbose up

migratedown:

	migrate -path db/migration -database "postgresql://root:498040@localhost:5432/bank_service?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
