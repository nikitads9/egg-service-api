package note

import "context"

func (s *Service) RemoveNote(ctx context.Context, id int64) (int64, error) {
	return s.noteRepository.RemoveNote(ctx, id)
}