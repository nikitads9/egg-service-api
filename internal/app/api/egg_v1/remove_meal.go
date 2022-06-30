package egg_v1

import (
	"context"

	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func (i *Implementation) RemoveMeal(ctx context.Context, req *desc.RemoveMealRequest) (*desc.RemoveMealResponse, error) {
	res, err := i.eggService.RemoveMeal(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &desc.RemoveMealResponse{
		Removed: res,
	}, nil
}
