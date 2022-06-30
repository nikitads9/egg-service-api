package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/model"
)

func (s *Service) UpdateNote(ctx context.Context, note *model.NoteInfo) (error) {
	return s.noteRepository.UpdateNote(ctx, note)
}
