package meal_repo

import (
	"context"
	"fmt"

	meal_domain_mappers "github.com/Vyach12/meal-api/internal/repositories/mappers"
	"github.com/Vyach12/meal-api/internal/services/business"
)

func (r *repositoryImpl) CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error) {
	mealIngr := meal_domain_mappers.BusinessMealsToDomainIngredients(meals)

	createdMeals := make([]business.Meal, 0, len(meals))

	tx, err := r.pgx.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	const query = `INSERT INTO ingredients(name) VALUES($1)`

	for _, ingr := range mealIngr {
		_, err := tx.Exec(ctx, query, ingr.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to insert ingredient %s: %w", ingr.Name, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return createdMeals, nil
}
