package egg_v1

import (
	"context"

	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func (i *Implementation) GetMeal(ctx context.Context, req *desc.GetMealRequest) (*desc.GetMealResponse, error) {
	meal, err := i.eggService.GetMeal(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &desc.GetMealResponse{
		Id:        meal.Id,
		UserId:    meal.UserId,
		Date:      meal.Date,
		Weight:    meal.Weight,
		Proteins:  meal.Proteins,
		Fat:       meal.Fat,
		Carbo:     meal.Carbo,
		Nutrition: meal.Nutrition,
	}, nil
}
