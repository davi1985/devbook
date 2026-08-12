package database

import (
	"database/sql"
	"devbook-api/internal/config"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", config.DataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
