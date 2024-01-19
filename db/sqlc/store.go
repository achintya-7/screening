package db

import "database/sql"

type Store interface {
	Querier
}

type SqlStore struct {
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		Queries: New(db),
	}
}
