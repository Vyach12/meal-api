package app

import (
	"context"
	"log"
	"net/http"

	clients "github.com/Vyach12/meal-api/internal/clients"
	themealsdb_client "github.com/Vyach12/meal-api/internal/clients/the_meals_db"
	config "github.com/Vyach12/meal-api/internal/config"
)

type Clients struct {
	MealClient clients.MealClient
}

func InitClients(ctx context.Context, cfg config.Config) Clients {
	theMealDbConfig, err := themealsdb_client.NewConfig(ctx, cfg)
	
	if err != nil {
		log.Fatal(err)
	}

	return Clients{
		MealClient: themealsdb_client.New(theMealDbConfig, &http.Client{
			Timeout: theMealDbConfig.Timeout(),
		}),
	}
}
