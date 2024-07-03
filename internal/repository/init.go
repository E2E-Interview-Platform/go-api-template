package repository

import (
	"context"
	"fmt"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func InitializeDatabase(ctx context.Context) (*Database, error) {
	mysqlDBURL := fmt.Sprintf(
		"%s:%s@(%s)/%s",
		environment.DB_USER,
		environment.DB_PASSWORD,
		environment.DB_URL,
		environment.DB_NAME,
	)

	db, err := sqlx.Connect("mysql", mysqlDBURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	ctxlogger.Info(
		ctx,
		"Successfully connected to database '%s' at URL '%s' with user '%s'",
		environment.DB_NAME,
		environment.DB_URL,
		environment.DB_USER,
	)
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}
