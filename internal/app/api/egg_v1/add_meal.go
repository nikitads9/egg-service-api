package egg_v1

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nikitads9/egg-service-api/internal/app/convert"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func (i *Implementation) AddMeal(ctx context.Context, req *desc.AddMealRequest) (*desc.AddMealResponse, error) {
	id, err := i.eggService.AddMeal(ctx, convert.ToMealInfo(req))
	if err != nil {
		return nil, err
	}

	return &desc.AddMealResponse{
		Result: &desc.AddMealResponse_Result{
			Id: id,
		},
	}, nil
}
