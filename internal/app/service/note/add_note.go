package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/model"
)

func (s *Service) AddNote(ctx context.Context, note *model.NoteInfo) (int64, error) {
	return s.noteRepository.AddNote(ctx, note)
}
