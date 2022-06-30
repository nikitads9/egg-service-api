package egg

import (
	"context"

	"github.com/nikitads9/egg-service-api/internal/app/model"
)

func (s *Service) GetMeal(ctx context.Context, id int64, userId int64) (*model.MealInfo, error) {
	return s.mealRepository.GetMeal(ctx, id, userId)
}
