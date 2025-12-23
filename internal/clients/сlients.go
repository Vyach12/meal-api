package clients

import (
	"context"
	"net/http"

	models "github.com/Vyach12/meal-api/internal/models"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MealClient interface {
	FetchRandomMeals(ctx context.Context) (models.Meal, error)
}
