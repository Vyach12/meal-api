package meal_search_service

import (
	"context"

	meal_v1 "github.com/Vyach12/meal-api/pkg/proto/meal/v1"
)

type MealSeekerService struct {
	meal_v1.UnimplementedMealsServiceServer
}

func (s *MealSeekerService) SearchMeals(ctx context.Context, req *meal_v1.SearchMealsRequest) (*meal_v1.SearchMealsResponse, error) {

	resp := meal_v1.Meal {
		Name: "Еда",
	}

	return &meal_v1.SearchMealsResponse{
		Meals: []*meal_v1.Meal {
			&resp,
		},
	}, nil
}