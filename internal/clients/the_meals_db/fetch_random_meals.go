package themealsdb_client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	"github.com/Vyach12/meal-api/internal/clients/the_meals_db/mappers"
	business "github.com/Vyach12/meal-api/internal/services/model"
)

var (
	TheMealDbStatusIsNotOkError = errors.New("Status is not 200(((")
	TheMealDbEmptyRespone       = errors.New("Empty response from the meal db")
)

func (c *clientImpl) FetchRandomMeals(ctx context.Context) (business.Meal, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.config.Url(), nil)

	if err != nil {
		return business.Meal{}, err
	}

	resp, err := c.cl.Do(req)
	if err != nil {
		return business.Meal{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return business.Meal{}, TheMealDbStatusIsNotOkError
	}

	var mealDto dto.RandomMealResponse
	if err := json.NewDecoder(resp.Body).Decode(&mealDto); err != nil {
		return business.Meal{}, err
	}

	if len(mealDto.Meals) == 0 {
		return business.Meal{}, TheMealDbEmptyRespone
	}

	meal, err := mappers.TransportMealToBusinessMeal(mealDto.Meals[0])
	if err != nil {
		return business.Meal{}, err
	}

	return meal, nil
}
