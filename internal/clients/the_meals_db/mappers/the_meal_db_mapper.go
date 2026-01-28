package mappers

import (
	"fmt"
	"strconv"
	"strings"

	dto "github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	"github.com/Vyach12/meal-api/internal/services/business"
	"github.com/samber/lo"
)

func TransportMealToBusinessMeal(meal dto.Meal) (business.Meal, error) {
	id, err := strconv.Atoi(meal.IDMeal)
	if err != nil {
		return business.Meal{}, fmt.Errorf("can not convert id for meal from dto: %w", err)
	}

	var (
		ingridients = meal.GetIngredients()
		measures    = meal.GetMeasures()
	)

	return business.Meal{
		ExternalID: int64(id),
		Name:       meal.StrMeal,
		Category: business.MealCategory{
			Name: meal.StrCategory,
		},
		Cuisine: business.MealCuisine{
			Name: meal.StrArea,
		},
		Instructions: meal.StrInstructions,
		ImageURL:     meal.StrMealThumb,
		Tags:         strings.Split(meal.StrTags, ","),
		YouTubeURL:   meal.StrYoutube,
		Ingredients: lo.Filter(lo.Map(ingridients, func(ingridientName string, idx int) business.MealIngredient {
			return business.MealIngredient{
				Name:    ingridientName,
				Measure: measures[idx],
			}
		}), func(ingridient business.MealIngredient, idx int) bool {
			return len(ingridient.Name) > 0
		}),
		RecipeURL: meal.StrSource,
	}, nil
}
