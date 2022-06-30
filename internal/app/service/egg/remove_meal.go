package egg

import "context"

func (s *Service) RemoveMeal(ctx context.Context, id int64, userId int64) (int64, error) {
	return s.mealRepository.RemoveMeal(ctx, id, userId)
}