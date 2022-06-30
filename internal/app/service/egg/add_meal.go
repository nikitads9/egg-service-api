package egg

import (
	"context"

	"github.com/nikitads9/egg-service-api/internal/app/model"
)

func (s *Service) AddMeal(ctx context.Context, meal *model.MealInfo) (int64, error) {
	return s.mealRepository.AddMeal(ctx, meal)
}
