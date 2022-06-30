package model

type NoteInfo struct {
	Id      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}
