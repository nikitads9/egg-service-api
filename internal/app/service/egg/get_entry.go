package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.NoteInfo, error) {
	return s.noteRepository.GetNote(ctx, id)
}
