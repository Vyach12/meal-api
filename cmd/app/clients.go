package app

import (
	"context"
	"log"
	"net/http"

	themealsdb_client "github.com/Vyach12/meal-api/internal/clients/the_meals_db"
	"github.com/Vyach12/meal-api/internal/services/model"
	"github.com/gomeal/config/pkg/config"
)

type Clients struct {
	MealClient MealClient
}

type MealClient interface {
	FetchRandomMeals(ctx context.Context) (business.Meal, error)
}

func InitClients(ctx context.Context, provider config.Provider) Clients {
	theMealDbConfig, err := themealsdb_client.NewConfig(ctx, provider)
	if err != nil {
		log.Fatal(err)
	}

	return Clients{
		MealClient: themealsdb_client.New(theMealDbConfig, &http.Client{
			Timeout: theMealDbConfig.Timeout(),
		}),
	}
}
