package convert

import (
	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToNoteInfo(req *desc.AddNoteRequest) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}
}
