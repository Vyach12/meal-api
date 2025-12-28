package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Vyach12/meal-api/cmd/app"
	"github.com/Vyach12/meal-api/internal/services/business"
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

	var mealIngredients = make([]business.MealIngredient, 0)
	for _, ingredient := range r.Ingredients {
		mealIngredients = append(mealIngredients, ingredient)
	}

	v, err := repositories.Meal.CreateMealsIngredient(ctx, mealIngredients)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

}
