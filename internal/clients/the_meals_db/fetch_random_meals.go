package themealsdb_client

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	"github.com/Vyach12/meal-api/internal/clients/the_meals_db/mappers"
	"github.com/Vyach12/meal-api/internal/models"
)

var (
	MealError = errors.New("Ашибка")
)

func (c *Client) FetchRandomMeals(ctx context.Context) ([]models.Meal, error) {
	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, c.config.Url(), nil)

	if err != nil {
		log.Fatal(err)
		return []models.Meal{}, err
	}
	defer resp.Body.Close()

	var mealDto dto.RandomMealResponse
	if err := json.NewDecoder(resp.Body).Decode(&mealDto); err != nil {
		log.Fatal(err)
		return []models.Meal{}, err
	}

	var buisnessMeals = make([]models.Meal, len(mealDto.Meals))

	for _, m := range mealDto.Meals {
		buisnessMeals = append(buisnessMeals, mappers.TransportMealToBuisnessMeal(m))
	}

	return buisnessMeals, nil
}
