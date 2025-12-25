package mappers

import (
	"fmt"
	"strconv"
	"strings"

	dto "github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	buisness "github.com/Vyach12/meal-api/internal/models"
	"github.com/samber/lo"
)

func TransportMealToBuisnessMeal(meal dto.Meal) (buisness.Meal, error) {
	id, err := strconv.Atoi(meal.IDMeal)
	if err != nil {
		return buisness.Meal{}, fmt.Errorf("can not convert id for meal from dto: %w", err)
	}

	var (
		ingridients = meal.GetIngredients()
		measures    = meal.GetMeasures()
	)

	return buisness.Meal{
		ExternalID: int64(id),
		Name:       meal.StrMeal,
		Category: buisness.MealCategory{
			Name: meal.StrCategory,
		},
		Cuisine: buisness.MealCuisine{
			Name: meal.StrArea,
		},
		Instructions: meal.StrInstructions,
		ImageURL:     meal.StrMealThumb,
		Tags:         strings.Split(meal.StrTags, ","),
		YoutubeUrl:   meal.StrYoutube,
		Ingredients: lo.Filter(lo.Map(ingridients, func(ingridientName string, idx int) buisness.MealIngridient {
			return buisness.MealIngridient{
				Name:    ingridientName,
				Measure: measures[idx],
			}
		}), func(ingridient buisness.MealIngridient, idx int) bool {
			return len(ingridient.Name) > 0
		}),
		RecipeUrl: meal.StrSource,
	}, nil
}
