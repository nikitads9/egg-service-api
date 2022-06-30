package convert

import (
	"github.com/nikitads9/egg-service-api/internal/app/model"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func ToGetListResponse(meals []*model.MealInfo) *desc.GetListResponse {
	res := make([]*desc.GetListResponse_Result, 0, len(meals))
	for _, elem := range meals {
		res = append(res, &desc.GetListResponse_Result{
			Id:        elem.Id,
			UserId:    elem.UserId,
			Date:      elem.Date,
			Weight:    elem.Weight,
			Proteins:  elem.Proteins,
			Fat:       elem.Fat,
			Carbo:     elem.Carbo,
			Nutrition: elem.Nutrition,
		})
	}
	return &desc.GetListResponse{
		Results: res,
	}
}
