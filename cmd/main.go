package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Vyach12/meal-api/cmd/app"
	"github.com/Vyach12/meal-api/internal/config"
)

func main() {
	ctx := context.Background()
	config := config.NewConfig(".cfg/values.yaml")

	var (
		clients = app.InitClients(ctx, *config)
	)

	var v, err = clients.MealClient.FetchRandomMeals(ctx)

	if(err != nil) {
		log.Fatal(err)
	}

	fmt.Println(v)
}
