package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/jmoiron/sqlx"
)

type BaseRepository struct {
	DB *sqlx.DB
}

type BaseTransaction struct {
	tx *sqlx.Tx
}

func (repo *BaseRepository) TimeNow() time.Time {
	return time.Now()
}

func (repo *BaseRepository) BeginTx(ctx context.Context) (repository.Transaction, error) {
	txObj, err := repo.DB.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		ctxlogger.Error(ctx, "error occurred while initiating database transaction: %v", err.Error())
		return nil, err
	}

	return &BaseTransaction{tx: txObj}, nil
}

func (repo *BaseRepository) HandleTransaction(ctx context.Context, tx repository.Transaction, isSuccess bool) error {
	var err error
	if !isSuccess {
		err = tx.Rollback()
		if err != nil {
			log.Printf("error occurred while rollback database transaction: %v", err.Error())
			return err
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error occurred while commit database transaction: %v", err.Error())
		return err
	}
	return err
}

func (repo *BaseTransaction) Commit() error {
	return repo.tx.Commit()
}

func (repo *BaseTransaction) Rollback() error {
	return repo.tx.Rollback()
}

func (repo *BaseRepository) initiateQueryExecutor(tx repository.Transaction) sqlx.Ext {
	// Executor to execute query based on transaction or db
	var executor sqlx.Ext = repo.DB
	if tx != nil {
		txObj := tx.(*BaseTransaction)
		executor = txObj.tx
	}

	return executor
}
