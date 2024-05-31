package database

import (
	"database/sql"
	"log/slog"

	"github.com/guicastro13/go-store/config/env"
	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	postgresURI := env.Env.DatabaseURL
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	slog.Info("database connected", slog.String("package", "database"))
	return db, nil
}
