package egg_v1

import (
	"context"

	"github.com/nikitads9/egg-service-api/internal/app/convert"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateMeal(ctx context.Context, req *desc.UpdateMealRequest) (*emptypb.Empty, error) {
	err := i.eggService.UpdateMeal(ctx, convert.ToMealUpdateInfo(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
