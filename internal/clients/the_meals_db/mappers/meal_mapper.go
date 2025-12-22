package mappers

import (
	"log"
	"strconv"
	"strings"

	dto "github.com/Vyach12/meal-api/internal/clients/the_meals_db/dto"
	buisness "github.com/Vyach12/meal-api/internal/models"
)

func TransportMealToBuisnessMeal(meal dto.Meal) buisness.Meal {

	id, err := strconv.Atoi(meal.IDMeal)
	if err != nil {
		log.Fatal("Сын шлюхи")
	}

	return buisness.Meal{
		Id:            int64(id),
		MealName:      meal.StrMeal,
		MealAlternate: meal.StrMealAlternate,
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
	}
}
