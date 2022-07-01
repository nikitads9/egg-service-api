package convert

import (
	"github.com/nikitads9/egg-service-api/internal/app/model"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func ToMealUpdateInfo(req *desc.UpdateMealRequest) *model.MealInfo {
	return &model.MealInfo{
		Id:       req.GetId(),
		UserId:   req.GetUserId(),
		MealDate: req.GetMealDate(),
		Weight:   req.GetWeight(),
	}
}
