package sqlc

import (
	"database/sql"
)

// Provides all functions to execute database queries and transactions
type Store interface {
	Querier
}

// Provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// Creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
