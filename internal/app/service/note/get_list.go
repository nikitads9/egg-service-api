package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/model"
)

func (s *Service) GetList(ctx context.Context) ([]*model.NoteInfo, error) {
	return s.noteRepository.GetList(ctx)
}
