package database

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPgx(connectionString string) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("[pgx] can't parse connection: %v", err)
	}
	if connConfig == nil {
		return nil, errors.New("[pgx] connection config is nil")
	}

	db := sqlx.NewDb(stdlib.OpenDB(*connConfig), "pgx")
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[pgx] can't ping: %s", err)
	}

	return db, nil
}
