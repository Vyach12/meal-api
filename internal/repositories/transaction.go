package transaction

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *transactorImpl) Transaction(ctx context.Context, operation func(ctx context.Context) error) error {

	fmt.Println("Start transaction")
	defer fmt.Println("End transaction")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := operation(ctx); err != nil {
		fmt.Println("попытка отката")
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			fmt.Println("ошибка отката")
			return rbErr
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}	

	return nil
}
