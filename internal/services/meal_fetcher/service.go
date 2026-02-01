package meal_fetcher_service

import (
	"context"
	"log/slog"

	business "github.com/Vyach12/meal-api/internal/services/meal_fetcher/model"
)

type MealClient interface {
	FetchRandomMeals(ctx context.Context) (business.Meal, error)
}

type MealRepository interface {
	CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error)
}

type serviceImpl struct {
	log *slog.Logger
	client MealClient
	repo MealRepository
}

func New(log *slog.Logger, mealClient MealClient, mealRepo MealRepository) (*serviceImpl, error) {
	return &serviceImpl{
		log: log,
		client: mealClient,
		repo: mealRepo,
	}, nil
}