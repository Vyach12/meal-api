package meal_repo

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Vyach12/meal-api/internal/services/business"
)

func (r *repositoryImpl) CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error) {
	//TODO: по хорошему сделать создание n meals в одном запросе, а не в разных, если не в падлу будет
	createdMeals := make([]business.Meal, 0, len(meals))

	txErr := r.transactor.Transaction(ctx, func(ctx context.Context) error {
		for _, meal := range meals {
			m, err := r.createMeal(ctx, meal)
			if err != nil {
				return err
			}
			createdMeals = append(createdMeals, m)
		}
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return createdMeals, nil
}

func (r *repositoryImpl) createMeal(ctx context.Context, meal business.Meal) (business.Meal, error) {
	//TODO: Добавить обработку
	category, err := r.createMealCategory(ctx, meal.Category)
	if err != nil {
		return business.Meal{}, err
	}

	cuisine, err := r.createMealCuisine(ctx, meal.Cuisine)
	if err != nil {
		return business.Meal{}, err
	}
	
	ingrs, err := r.createMealsIngredient(ctx, meal.Ingredients)
	if err != nil {
		return business.Meal{}, err
	}
	r.saveMeal(ctx, meal, category.ID, cuisine.ID)

	for i, ingr := range ingrs {
		r.createMealIngredientLink(ctx, meal.ID, ingr.ID, ingr.Measure, int64(i+1))
	}

	return business.Meal{}, nil
}

func (r *repositoryImpl) createMealsIngredient(ctx context.Context, ingredients []business.MealIngredient) ([]business.MealIngredient, error) {

	insertBuilder := sq.Insert("ingredients").
		PlaceholderFormat(sq.Dollar).
		Columns("name")
	for _, meal := range ingredients {
		insertBuilder = insertBuilder.Values(meal.Name)
	}
	query, args, err := insertBuilder.
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", time.Now()).
		Suffix("RETURNING id, name").
		ToSql()

	if err != nil {
		log.Println("Ошибка формирования запроса")
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("Ошибка выполнения запроса:", err)
		return nil, err
	}
	defer rows.Close()

	createdIngredients := make([]business.MealIngredient, 0, len(ingredients))

	for rows.Next() {
		var ingredient business.MealIngredient
		if err := rows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
			return nil, err
		}
		createdIngredients = append(createdIngredients, ingredient)
	}

	return createdIngredients, nil
}

func (r *repositoryImpl) createMealCategory(ctx context.Context, category business.MealCategory) (business.MealCategory, error) {
	query, args, err := sq.Insert("categories").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(category.Name).
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", time.Now()).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println("Ошибка формирования запроса")
		return business.MealCategory{}, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		log.Println("Ошибка выполнения запроса: ", query)
		return business.MealCategory{}, err
	}

	category.ID = id

	return category, nil
}

func (r *repositoryImpl) createMealCuisine(ctx context.Context, cuisines business.MealCuisine) (business.MealCuisine, error) {
	query, args, err := sq.Insert("cuisines").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(cuisines.Name).
		Suffix("ON CONFLICT (name) DO UPDATE SET updated_at = ?", time.Now()).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println("Ошибка формирования запроса")
		return business.MealCuisine{}, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		log.Println("Ошибка выполнения запроса: ", query)
		return business.MealCuisine{}, err
	}

	cuisines.ID = id

	return cuisines, nil
}

func (r *repositoryImpl) createMealIngredientLink(ctx context.Context, mealId int64, ingredientId int64, meausre string, position int64) error {
	query, args, err := sq.Insert("meal_ingredients").
		PlaceholderFormat(sq.Dollar).
		Columns("meal_id", "ingredient_id", "measure", "position").
		Values(mealId, ingredientId, meausre, position).
		Suffix("ON CONFLICT (meal_id, ingredient_id) DO UPDATE SET updated_at = ?", time.Now()).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println("Ошибка формирования запроса")
		return err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		log.Println("Ошибка выполнения запроса: ", query)
		return err
	}

	return nil
}

func (r *repositoryImpl) saveMeal(ctx context.Context, meal business.Meal, categoryID, cuisineID int64) (int64, error) {

	query, args, err := sq.Insert("meals").
		PlaceholderFormat(sq.Dollar).
		Columns("external_id", "name", "category_id", "cuisine_id", "instructions", "image_url", "tags", "youtube_url", "recipe_url").
		Values(meal.ExternalID, meal.Name, categoryID, cuisineID, meal.Instructions, meal.ImageURL, meal.Tags, meal.YouTubeURL, meal.RecipeURL).
		Suffix("ON CONFLICT (external_id) DO UPDATE SET updated_at = ?").
		Suffix("RETURNING id, name").
		ToSql()

	if err != nil {
		return 0, err
	}

	var id int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
