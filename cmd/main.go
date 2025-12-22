package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Vyach12/meal-api/cmd/app"
)

func main() {
	ctx := context.Background()

	var (
		clients = app.InitClients(ctx)
	)

	var v, err = clients.MealClient.FetchRandomMeals(ctx)

	if(err != nil) {
		log.Fatal(err)
	}

	fmt.Println(v)
}
