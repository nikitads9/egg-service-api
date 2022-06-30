package convert

import (
	"github.com/nikitads9/egg-service-api/internal/app/model"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func ToMealInfo(req *desc.AddMealRequest) *model.MealInfo {
	return &model.MealInfo{
		UserId: req.GetUserId(),
		Date:   req.GetDate(),
		Weight: req.GetWeight(),
	}
}
