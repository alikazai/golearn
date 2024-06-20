package postgres

import "database/sql"

type postgresStore struct {
	db *sql.DB
}

// NewPostgresStore create new isntance of postgresstore
func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
