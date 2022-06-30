package convert

import (
	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToGetListResponse(notes []*model.NoteInfo) *desc.GetListResponse {
	res := make([]*desc.GetListResponse_Result, 0, len(notes))
	for _, elem := range notes {
		res = append(res, &desc.GetListResponse_Result{
			Title:   elem.Title,
			Content: elem.Content,
		})
	}
	return &desc.GetListResponse{
		Results: res,
	}
}
