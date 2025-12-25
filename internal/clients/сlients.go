package clients

import (
	"context"
	"net/http"

	"github.com/Vyach12/meal-api/internal/services/business"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MealClient interface {
	FetchRandomMeals(ctx context.Context) (business.Meal, error)
}
