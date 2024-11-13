package repository

import (
	"echo-golang/model"
	model_response "echo-golang/model/response"

	"github.com/stretchr/testify/mock"
)

type INoteRepositoryMock struct {
	Mock mock.Mock
}

func (repository *INoteRepositoryMock) GetNote() ([]model_response.Note, error) {
	args := repository.Mock.Called()
	return args.Get(0).([]model_response.Note), args.Error(1)
}

func (repository *INoteRepositoryMock) InsertNote(note model_response.Note) (model.BaseResponse[model_response.Note], error) {
	args := repository.Mock.Called(note)
	if args.Get(0) == nil {
		return model.BaseResponse[model_response.Note]{Message: "Failed to insert note", Data: nil}, args.Error(1)
	} else {
		return model.BaseResponse[model_response.Note]{Message: "Succesful insert note", Data: nil}, nil
	}
}

func (repository *INoteRepositoryMock) DeleteNoteById(idNote int) (model.BaseResponse[model_response.Note], error) {
	args := repository.Mock.Called(idNote)
	if args.Get(0) == nil {
		return model.BaseResponse[model_response.Note]{Message: "Failed to delete note", Data: nil}, args.Error(1)
	} else {
		return model.BaseResponse[model_response.Note]{Message: "Succesful delete note", Data: nil}, nil
	}
}

func (repository *INoteRepositoryMock) UpdateNoteById(idNote int, note model_response.Note) (model.BaseResponse[model_response.Note], error) {
	args := repository.Mock.Called(idNote, note)
	if args.Get(0) == nil {
		return model.BaseResponse[model_response.Note]{Message: "Failed to update note", Data: nil}, args.Error(1)
	} else {
		return model.BaseResponse[model_response.Note]{Message: "Succesful update note", Data: nil}, nil
	}
}
