package meal_repo

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Vyach12/meal-api/internal/services/business"
)

func (r *repositoryImpl) CreateMealsIngredient(ctx context.Context, meals []business.MealIngredient) ([]business.MealIngredient, error) {

	createdMeals := make([]business.MealIngredient, 0, len(meals))

	txErr := r.transactor.Transaction(ctx, func(ctx context.Context) error {
		for _, meal := range meals {
			createdMeal, err := r.createMealIngredient(ctx, meal)
			if err != nil {
				return fmt.Errorf("failed to insert ingredient %s: %w", meal.Name, err)
			}
			createdMeals = append(createdMeals, createdMeal)
		}
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return createdMeals, nil
}

func (r *repositoryImpl) createMealIngredient(ctx context.Context, meal business.MealIngredient) (business.MealIngredient, error) {
	query, args, err := sq.Insert("ingredients").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(meal.Name).
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", time.Now()).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println("Ошибка формирования запроса")
		return business.MealIngredient{}, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		log.Println("Ошибка выполнения запроса: ", query)
		return business.MealIngredient{}, err
	}

	meal.ID = id

	return meal, nil
}
