package meal_domain_mappers

import (
	"database/sql"

	domain "github.com/Vyach12/meal-api/internal/repositories/meal/models"
	"github.com/Vyach12/meal-api/internal/services/business"
	"github.com/samber/lo"
)

func BusinessMealToDomainMeals(businessMeals []business.Meal) []domain.Meal {
	return lo.Map(businessMeals, func(businessMeal business.Meal, i int) domain.Meal {
		return domain.Meal{
			ID:           businessMeal.ID,
			ExternalID:   businessMeal.ExternalID,
			Name:         businessMeal.Name,
			CategoryID:   businessMeal.Category.ID,
			CuisineID:    businessMeal.Cuisine.ID,
			Instructions: businessMeal.Instructions,
			ImageURL: sql.NullString{
				Valid:  businessMeal.ImageURL != "",
				String: businessMeal.ImageURL,
			},
			Tags: businessMeal.Tags,
			YouTubeURL: sql.NullString{
				Valid:  businessMeal.YouTubeURL != "",
				String: businessMeal.YouTubeURL,
			},
			RecipeURL: sql.NullString{
				Valid:  businessMeal.RecipeURL != "",
				String: businessMeal.RecipeURL,
			},
		}
	})
}

func BusinessMealsToDomainCategories(businessMeals []business.Meal) []domain.MealCategory {
	return lo.UniqBy(lo.Map(businessMeals, func(businessMeal business.Meal, idx int) domain.MealCategory {
		return domain.MealCategory{
			ID:   businessMeal.Category.ID,
			Name: businessMeal.Category.Name,
		}
	}), func(domainCategory domain.MealCategory) string {
		return domainCategory.Name
	})
}

func BusinessMealsToDomainCuisines(businessMeals []business.Meal) []domain.MealCuisine {
	return lo.UniqBy(lo.Map(businessMeals, func(businessMeal business.Meal, idx int) domain.MealCuisine {
		return domain.MealCuisine{
			ID:   businessMeal.Cuisine.ID,
			Name: businessMeal.Cuisine.Name,
		}
	}), func(domainCuisine domain.MealCuisine) string {
		return domainCuisine.Name
	})
}

func BusinessMealsToDomainIngredients(businessMeals []business.Meal) []domain.MealIngredient {
	businessIngredients := make([]business.MealIngredient, 0, len(businessMeals))
	for _, meal := range businessMeals {
		for _, ingredient := range meal.Ingredients {
			businessIngredients = append(businessIngredients, ingredient)
		}
	}

	return lo.UniqBy(lo.Map(businessIngredients, func(businessIngredient business.MealIngredient, idx int) domain.MealIngredient {
		return domain.MealIngredient{
			ID:   businessIngredient.ID,
			Name: businessIngredient.Name,
		}
	}), func(domainIngredient domain.MealIngredient) string {
		return domainIngredient.Name
	})
}

func BusinessMealsToDomainMealIngredientLinks(businessMeals []business.Meal, mealsByExternalID map[int64]domain.Meal, ingredientsByName map[string]domain.MealIngredient) []domain.MealIngridientsLink {
	var links []domain.MealIngridientsLink

	for _, businessMeal := range businessMeals {
		meal, exists := mealsByExternalID[businessMeal.ExternalID]
		if !exists {
			continue
		}

		for position, ingredient := range businessMeal.Ingredients {
			domainIngredient, exists := ingredientsByName[ingredient.Name]
			if !exists {
				continue
			}

			links = append(links, domain.MealIngridientsLink{
				MealID:       meal.ID,
				IngridientID: domainIngredient.ID,
				Measure:      ingredient.Measure,
				Position:     int64(position + 1),
			})
		}
	}

	return links
}
