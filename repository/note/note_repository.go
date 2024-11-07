package repository

import "echo-golang/model"

type INoteRepository interface {
	GetNote() ([]model.Note, error)
	InsertNote(note model.Note) (model.BaseResponse[model.Note], error)
	DeleteNoteById(id int) (model.BaseResponse[model.Note], error)
	UpdateNoteById(id int, note model.Note) (model.BaseResponse[model.Note], error)
}
