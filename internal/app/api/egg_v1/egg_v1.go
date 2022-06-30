package egg_v1

import (
	serv "github.com/nikitads9/egg-service-api/internal/app/service/egg"
	desc "github.com/nikitads9/egg-service-api/pkg/egg_api"
)

type Implementation struct {
	desc.UnimplementedEggNutritionServer
	eggService *serv.Service
}

func NewEggNutrition(eggService *serv.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedEggNutritionServer{},
		eggService,
	}
}
