package app

import (
	"context"
	"log/slog"

	meal_fetcher_service "github.com/Vyach12/meal-api/internal/services/meal_fetcher"
	business "github.com/Vyach12/meal-api/internal/services/meal_fetcher/model"
)

type Services struct {
	MealFetcherService MealFetcherService
}

type MealFetcherService interface {
	FetchRandomMeals(ctx context.Context) ([]business.Meal, error)
}

func InitServices(log *slog.Logger, ctx context.Context, clients Clients, repos Repositories) Services {
	mealFetcherService, err := meal_fetcher_service.New(log, clients.MealClient, repos.Meal)
	if err != nil {
		panic(err.Error())
	}

	return Services{
		MealFetcherService: mealFetcherService,
	}
}
