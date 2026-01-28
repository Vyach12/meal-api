package app

import (
	"context"
	"log"

	app_config "github.com/Vyach12/meal-api/internal/config"
	transaction "github.com/Vyach12/meal-api/internal/repositories"
	meal_repo "github.com/Vyach12/meal-api/internal/repositories/meal"
	"github.com/Vyach12/meal-api/internal/services/business"
	"github.com/gomeal/config/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealRepository interface {
	CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error)
}

type Repositories struct {
	Meal MealRepository
}

func InitRepo(ctx context.Context, provider config.Provider) Repositories {
	pool, err := pgxpool.New(ctx, app_config.PostgresURI(provider))
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	return Repositories{
		Meal: meal_repo.New(pool, transaction.New(pool)),
	}
}
