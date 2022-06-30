package egg

import (
	"context"

	"github.com/nikitads9/egg-service-api/internal/app/model"
)

func (s *Service) UpdateMeal(ctx context.Context, note *model.MealInfo) (error) {
	return s.mealRepository.UpdateMeal(ctx, note)
}
