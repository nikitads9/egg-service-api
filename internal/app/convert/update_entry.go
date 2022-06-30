package convert

import (
	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToNoteUpdateInfo(req *desc.UpdateNoteRequest) *model.NoteInfo {
	return &model.NoteInfo{
		Id:      req.GetId(),
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}
}
