package repository

import "echo-golang/model"

type INoteRepository interface {
	GetNote() ([]model.Note, error)
}
