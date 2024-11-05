package repository

import "echo-golang/model"

type INoteRepository interface {
	GetNote() ([]model.Note, error)
	InsertNote(note model.Note) (model.BaseResponse[model.Note], error)
}
