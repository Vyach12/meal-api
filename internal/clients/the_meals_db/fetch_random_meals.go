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
	TheMealDbStatusNotOkError = errors.New("Status not 200")
)

func (c *Client) FetchRandomMeals(ctx context.Context) (models.Meal, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.config.Url(), nil)

	if err != nil {
		log.Fatal(err)
		return models.Meal{}, err
	}

	resp, err := c.cl.Do(req)
	if err != nil {
		log.Fatal(err)
		return models.Meal{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.Meal{}, err
	}

	var mealDto dto.RandomMealResponse
	if err := json.NewDecoder(resp.Body).Decode(&mealDto); err != nil {
		log.Fatal(err)
		return models.Meal{}, err
	}

	return mappers.TransportMealToBuisnessMeal(mealDto.Meals[0]), nil
}
