package meal_fetcher_service

import (
	"context"

	business "github.com/Vyach12/meal-api/internal/services/meal_fetcher/model"
)

func (s *serviceImpl) FetchRandomMeals(ctx context.Context) ([]business.Meal, error) {
	s.log.Debug("try to fetch random meal")
	defer s.log.Debug("end fetching random meal")

	meal, err := s.client.FetchRandomMeals(ctx)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	createdMeals, err := s.repo.CreateMeals(ctx, []business.Meal{meal})
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return createdMeals, nil
}
