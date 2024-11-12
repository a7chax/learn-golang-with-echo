package service

import (
	"echo-golang/model"
	repository "echo-golang/repository/note"
)

type INoteService interface {
	GetAllNote() ([]model.Note, error)
	InsertNote(note model.Note) (model.BaseResponse[model.Note], error)
	DeleteNoteById(idNote int) (model.BaseResponse[model.Note], error)
	UpdateNoteById(idNote int, note model.Note) (model.BaseResponse[model.Note], error)
}

type NoteService struct {
	repo repository.INoteRepository
}

func NewNoteService(repo repository.INoteRepository) *NoteService {
	return &NoteService{repo}
}

func (s *NoteService) GetAllNote() ([]model.Note, error) {
	return s.repo.GetNote()
}

func (s *NoteService) InsertNote(note model.Note) (model.BaseResponse[model.Note], error) {
	return s.repo.InsertNote(note)
}

func (s *NoteService) DeleteNoteById(idNote int) (model.BaseResponse[model.Note], error) {
	return s.repo.DeleteNoteById(idNote)
}

func (s *NoteService) UpdateNoteById(idNote int, note model.Note) (model.BaseResponse[model.Note], error) {
	return s.repo.UpdateNoteById(idNote, note)
}
