package service

import (
	entity_note "echo-golang/entity/note"
	"echo-golang/model"
	model_request "echo-golang/model/request"
	repository "echo-golang/repository/note"
	user_repository "echo-golang/repository/user"
	"fmt"
)

type INoteService interface {
	GetAllNote() ([]entity_note.Note, error)
	InsertNote(note model_request.Note) (model.BaseResponseNoData, error)
	DeleteNoteById(idNote int) (model.BaseResponseNoData, error)
	UpdateNoteById(idNote int, note model_request.Note) (model.BaseResponseNoData, error)
}

type NoteService struct {
	repo repository.INoteRepository
	user user_repository.IUserRepository
}

func NewNoteService(
	repo repository.INoteRepository,
	user user_repository.IUserRepository) *NoteService {
	return &NoteService{
		repo,
		user,
	}
}

func (s *NoteService) GetAllNote() ([]entity_note.Note, error) {
	notes, err := s.repo.GetNote()
	if err != nil {
		return nil, err
	}

	var entityNotes []entity_note.Note
	for _, note := range notes {
		entityNotes = append(entityNotes, entity_note.Note{
			IdNote:  note.IdNote,
			Title:   note.Title,
			Content: note.Content,
			// Add other fields as necessary
		})
	}

	return entityNotes, err
}

func (s *NoteService) InsertNote(note model_request.Note) (model.BaseResponseNoData, error) {
	_, err := s.repo.InsertNote(note)

	if err != nil {
		return model.BaseResponseNoData{
			Message:   "Failed to insert note",
			IsSuccess: false,
		}, err
	}
	fmt.Println(note.IdUser, "idser")

	response, err := s.user.GetUser(note.IdUser)
	if err != nil {
		return model.BaseResponseNoData{
			Message:   "Failed to insert note1",
			IsSuccess: false,
		}, err
	}

	return model.BaseResponseNoData{
		Message:   "Succesful insert note" + response.Username,
		IsSuccess: true,
	}, nil
}

func (s *NoteService) DeleteNoteById(idNote int) (model.BaseResponseNoData, error) {
	_, err := s.repo.DeleteNoteById(idNote)

	if err != nil {
		return model.BaseResponseNoData{
			Message:   "Failed to delete note",
			IsSuccess: false,
		}, err
	}
	return model.BaseResponseNoData{
		Message:   "Succesful delete note",
		IsSuccess: true,
	}, err
}

func (s *NoteService) UpdateNoteById(idNote int, note model_request.Note) (model.BaseResponseNoData, error) {
	// return s.repo.UpdateNoteById(idNote, note)
	_, err := s.repo.UpdateNoteById(idNote, note)

	if err != nil {
		return model.BaseResponseNoData{
			Message:   "Failed to update note",
			IsSuccess: false,
		}, err
	}
	return model.BaseResponseNoData{
		Message:   "Succesful update note",
		IsSuccess: true,
	}, err
}
