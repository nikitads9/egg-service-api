package egg_v1

import (
	"context"

	"github.com/nikitads9/egg-service-api/internal/app/convert"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

func (i *Implementation) GetList(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	mealInfo, err := i.eggService.GetList(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return convert.ToGetListResponse(mealInfo), nil
}
