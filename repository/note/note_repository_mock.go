package repository

import (
	"echo-golang/model"

	"github.com/stretchr/testify/mock"
)

type INoteRepositoryMock struct {
	Mock mock.Mock
}

func (repository *INoteRepositoryMock) GetNote() ([]model.Note, error) {
	args := repository.Mock.Called()
	return args.Get(0).([]model.Note), args.Error(1)
}

func (repository *INoteRepositoryMock) InsertNote(note model.Note) (model.BaseResponse[model.Note], error) {
	args := repository.Mock.Called(note)
	if args.Get(0) == nil {
		return model.BaseResponse[model.Note]{Message: "Failed to insert note", Data: nil}, args.Error(1)
	} else {
		return model.BaseResponse[model.Note]{Message: "Succesful insert note", Data: nil}, nil
	}
}

func (repository *INoteRepositoryMock) DeleteNoteById(idNote int) (model.BaseResponse[model.Note], error) {
	args := repository.Mock.Called(idNote)
	if args.Get(0) == nil {
		return model.BaseResponse[model.Note]{Message: "Failed to delete note", Data: nil}, args.Error(1)
	} else {
		return model.BaseResponse[model.Note]{Message: "Succesful delete note", Data: nil}, nil
	}
}

func (repository *INoteRepositoryMock) UpdateNoteById(idNote int, note model.Note) (model.BaseResponse[model.Note], error) {
	args := repository.Mock.Called(idNote, note)
	if args.Get(0) == nil {
		return model.BaseResponse[model.Note]{Message: "Failed to update note", Data: nil}, args.Error(1)
	} else {
		return model.BaseResponse[model.Note]{Message: "Succesful update note", Data: nil}, nil
	}
}
