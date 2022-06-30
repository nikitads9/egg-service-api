package note_v1

import (
	serv "github.com/nikitads9/note-service-api/internal/app/service/note"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

type Implementation struct {
	desc.UnimplementedEggNutritionServer
	eggService *serv.Service
}

func NewEggNutrition(eggService *serv.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteV1Server{},
		eggService,
	}
}
