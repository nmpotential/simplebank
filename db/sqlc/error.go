package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Constants for Postgres database constraint violations
const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

// ErrRecordNotFound is returned when a query returns no rows
var ErrRecordNotFound = pgx.ErrNoRows

// ErrUniqueViolation is returned when a unique constraint is violated
var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

// ErrorCode returns the error code as a string
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
