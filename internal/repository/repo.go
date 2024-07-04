package repository

import (
	"context"
)

type RepositoryTransaction interface {
	BeginTx(ctx context.Context) (Transaction, error)
	HandleTransaction(ctx context.Context, tx Transaction, isSuccess bool) error
}

type Transaction interface {
	Commit() error
	Rollback() error
}
