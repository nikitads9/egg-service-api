package convert

import (
	"github.com/nikitads9/egg-service-api/internal/app/model"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func ToGetNoteResponse(meal *model.MealInfo) *desc.GetMealResponse {
	return &desc.GetMealResponse{
		Id:        meal.Id,
		UserId:    meal.UserId,
		MealDate:  meal.MealDate,
		Weight:    meal.Weight,
		Proteins:  meal.Proteins,
		Fat:       meal.Fat,
		Carbo:     meal.Carbo,
		Nutrition: meal.Nutrition,
	}
}
