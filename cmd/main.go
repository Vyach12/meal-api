package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Vyach12/meal-api/cmd/app"
	business "github.com/Vyach12/meal-api/internal/services/model"
	"github.com/gomeal/config/pkg/config"
)

func main() {
	ctx := context.Background()

	provider := config.NewProvider(".cfg/values.yaml")

	var (
		repositories = app.InitRepo(ctx, provider)
		clients      = app.InitClients(ctx, provider)
	)

	r, err := clients.MealClient.FetchRandomMeals(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(r)

	v, err := repositories.Meal.CreateMeals(ctx, []business.Meal{r})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

}
