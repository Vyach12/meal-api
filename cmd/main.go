package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Vyach12/meal-api/cmd/app"
	"github.com/gomeal/config/pkg/config"
)

func main() {
	ctx := context.Background()

	provider := config.NewProvider(".cfg/values.yaml")

	var (
		clients = app.InitClients(ctx, provider)
	)

	v, err := clients.MealClient.FetchRandomMeals(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
