package meal_repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Transactor interface {
	Transaction(ctx context.Context, f func(ctx context.Context) error) error
}

type repositoryImpl struct {
	db *pgxpool.Pool
	transactor Transactor
}

func New(pgx *pgxpool.Pool, transactor Transactor) *repositoryImpl {
	return &repositoryImpl{
		db: pgx,
		transactor: transactor,
	}
}
