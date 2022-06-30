package convert

import (
	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToGetNoteResponse(note *model.NoteInfo) *desc.GetNoteResponse {
	return &desc.GetNoteResponse{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
	}
}
