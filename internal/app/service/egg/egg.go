package egg

import "github.com/nikitads9/egg-service-api/internal/app/repository"

type Service struct {
	mealRepository repository.IEggNutritionRepository
}

func NewEggNutritionService(mealRepository repository.IEggNutritionRepository) *Service {
	return &Service{
		mealRepository: mealRepository,
	}
}
