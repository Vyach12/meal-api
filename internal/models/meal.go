package models

type (
	Meal struct {
		Id           int64
		ExternalID   int64
		Name         string
		Category     MealCategory
		Cuisine      MealCuisine
		Instructions string
		ImageURL     string
		Tags         []string
		YoutubeUrl   string
		Ingredients  []MealIngridient
		RecipeUrl    string
	}

	MealCategory struct {
		Id   int64
		Name string
	}

	MealCuisine struct {
		Id   int64
		Name string
	}

	MealIngridient struct {
		Id      int64
		Name    string
		Measure string
	}
)
