package note_v1

import (
	"context"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) RemoveNote(ctx context.Context, req *desc.RemoveNoteRequest) (*desc.RemoveNoteResponse, error) {
	res, err := i.noteService.RemoveNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.RemoveNoteResponse{
		Removed: res,
	}, nil
}
