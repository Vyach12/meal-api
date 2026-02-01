package themealsdb_client_test

import (
	"context"
	_ "embed"
	"io"
	"net/http"
	"strings"
	"testing"

	themealsdb_client "github.com/Vyach12/meal-api/internal/clients/the_meals_db"
	"github.com/Vyach12/meal-api/internal/clients/the_meals_db/mocks"
	business "github.com/Vyach12/meal-api/internal/services/meal_fetcher/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:embed testdata/meal_response_ok.json
var mockRandomMealResponse string

func TestFetchRandomMeal(t *testing.T) {
	type testCase struct {
		name             string
		congigMock       func() themealsdb_client.TheMealsDbClientConfig
		httpClientMock   func() themealsdb_client.HTTPClient
		expectedResponse business.Meal
		expectedError    error
	}

	tests := []testCase{
		{
			name: "valid response",
			congigMock: func() themealsdb_client.TheMealsDbClientConfig {
				c := mocks.NewMockTheMealsDbClientConfig(t)
				c.EXPECT().Url().Return("pididy.com")

				return c
			},

			httpClientMock: func() themealsdb_client.HTTPClient {
				c := mocks.NewMockHTTPClient(t)
				c.EXPECT().Do(mock.Anything).
					Return(&http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(mockRandomMealResponse)),
					}, nil)

				return c
			},
			expectedResponse: business.Meal{
				ID:         0,
				ExternalID: 53367,
				Name:       "Chicken Fried Rice",
				Category: business.MealCategory{
					ID:   0,
					Name: "Chicken",
				},
				Cuisine: business.MealCuisine{
					ID:   0,
					Name: "Chinese",
				},
				Instructions: "Fried rice is best made with leftover..",
				ImageURL:     "https://www.themealdb.com/images/media/meals/wuyd2h1765655837.jpg",
				Tags:         []string{""},
				YouTubeURL:   "https://www.youtube.com/watch?v=mUr-7wjhHuU",
				Ingredients: []business.MealIngredient{
					{
						ID:      0,
						Name:    "Chicken Thighs",
						Measure: "1 lb",
					},
					{
						ID:      0,
						Name:    "Salt",
						Measure: "1 tsp",
					},
				},
				RecipeURL: "https://www.simplyrecipes.com/recipes/chicken_fried_rice/",
			}}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			theMealsDbClient := themealsdb_client.New(tc.congigMock(), tc.httpClientMock())
			meal, err := theMealsDbClient.FetchRandomMeals(ctx)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedResponse, meal)
		})
	}
}
