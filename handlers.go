package main

import (
	"echo-golang/model"
	"echo-golang/repository"
)

type NoteHandler struct {
	repo repository.NoteRepository
}

func NewNoteHandler(repo repository.NoteRepository) *NoteHandler {
	return &NoteHandler{repo}
}

func (h *NoteHandler) GetNote() ([]model.Note, error) {
	return h.repo.GetNote()
}
