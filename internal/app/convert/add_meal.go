package convert

import (
	"github.com/nikitads9/egg-service-api/internal/app/model"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func ToMealInfo(req *desc.AddMealRequest) *model.MealInfo {
	weight := req.GetWeight()
	return &model.MealInfo{
		UserId:    req.GetUserId(),
		Weight:    weight,
		Proteins:  0.127 * weight,
		Fat:       0.115 * weight,
		Carbo:     0.07 * weight,
		Nutrition: 157 * (weight / 100),
	}
}
