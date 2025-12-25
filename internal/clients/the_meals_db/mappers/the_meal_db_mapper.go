package mappers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	"github.com/Vyach12/meal-api/internal/services/business"
)

func TransportMealToBusinessMeal(meal dto.Meal) (business.Meal, error) {

	id, err := strconv.Atoi(meal.IDMeal)
	if err != nil {
		return business.Meal{}, fmt.Errorf("can not convert id for meal from dto: %w", err)
	}

	return business.Meal{
		Id:           int64(id),
		MealName:     meal.StrMeal,
		Category:     meal.StrCategory,
		Area:         meal.StrArea,
		Instructions: meal.StrInstructions,
		MealThumb:    meal.StrMealThumb,
		Tags:         strings.Split(meal.StrTags, ","),
		YoutubeUrl:   meal.StrYoutube,
		Ingredients:  meal.GetIngredients(),
		Measure:      meal.GetMeasures(),
		Source:       meal.StrSource,
		ImageSource:  meal.StrImageSource,
	}, nil
}
