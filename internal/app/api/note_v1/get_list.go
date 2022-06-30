package note_v1

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/convert"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetList(ctx context.Context, in *emptypb.Empty) (*desc.GetListResponse, error) {
	noteInfo, err := i.noteService.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return convert.ToGetListResponse(noteInfo), nil
}
