package egg

import (
	"context"

	"github.com/nikitads9/egg-service-api/internal/app/model"
)

func (s *Service) GetList(ctx context.Context, userId int64) ([]*model.MealInfo, error) {
	return s.mealRepository.GetList(ctx, userId)
}
