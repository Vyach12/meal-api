package app

import (
	"context"
	"log"
	"net/http"

	clients "github.com/Vyach12/meal-api/internal/clients"
	themealsdb_client "github.com/Vyach12/meal-api/internal/clients/the_meals_db"
)

type Clients struct {
	MealClient clients.MealClient
}

func InitClients(ctx context.Context) Clients {
	config, err := themealsdb_client.NewConfig(ctx)

	if err != nil {
		log.Fatal("Сын бляди")
	}

	return Clients{
		MealClient: themealsdb_client.New(config, &http.Client{
			Timeout: config.Timeout(),
		}),
	}
}
