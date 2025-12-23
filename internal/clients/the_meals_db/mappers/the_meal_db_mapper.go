package mappers

import (
	"fmt"
	"strconv"
	"strings"

	dto "github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	buisness "github.com/Vyach12/meal-api/internal/models"
)

func TransportMealToBuisnessMeal(meal dto.Meal) (buisness.Meal, error) {

	id, err := strconv.Atoi(meal.IDMeal)
	if err != nil {
		return buisness.Meal{}, fmt.Errorf("can not convert id for meal from dto: %w", err)
	}

	return buisness.Meal{
		Id:            int64(id),
		MealName:      meal.StrMeal,
		Category:      meal.StrCategory,
		Area:          meal.StrArea,
		Instructions:  meal.StrInstructions,
		MealThumb:     meal.StrMealThumb,
		Tags:          strings.Split(meal.StrTags, ","),
		YoutubeUrl:    meal.StrYoutube,
		Ingredients:   meal.GetIngredients(),
		Measure:       meal.GetMeasures(),
		Source:        meal.StrSource,
		ImageSource:   meal.StrImageSource,
	}, nil
}
