package clients

import (
	"context"
	"net/http"
	"time"

	models "github.com/Vyach12/meal-api/internal/models"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MealClient interface {
	FetchRandomMeals(ctx context.Context) (models.Meal, error)
}

type TheMealsDbClientConfig interface {
	Url() string
	Timeout() time.Duration
}
