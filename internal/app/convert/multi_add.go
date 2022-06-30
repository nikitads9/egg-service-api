package convert

import (
	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToNotesInfo (req *desc.MultiAddRequest) []*model.NoteInfo {
	notes := req.GetNotes()
	res := make ([]*model.NoteInfo, 0, len(notes))

	for _, elem := range notes {
		res = append(res, &model.NoteInfo{
			Title: elem.GetTitle(),
			Content: elem.GetContent(),
		})
	}

	return res
}