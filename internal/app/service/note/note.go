package note

import "github.com/nikitads9/note-service-api/internal/app/repository"

type Service struct {
	noteRepository repository.INoteRepository
}

func NewNoteService(noteRepository repository.INoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
