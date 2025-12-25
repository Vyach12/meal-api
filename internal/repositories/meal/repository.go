package meal_repo

import "github.com/jackc/pgx/v5/pgxpool"

type repositoryImpl struct {
	pgx *pgxpool.Pool
}

func New(pgx *pgxpool.Pool) *repositoryImpl {
	return &repositoryImpl{
		pgx: pgx,
	}
}
