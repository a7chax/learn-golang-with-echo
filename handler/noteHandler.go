package handler

import (
	DTO "echo-golang/dto"
	"echo-golang/model"
	"echo-golang/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type INoteHandler struct {
	Repo repository.INoteRepository
}

func NoteHandler(repo repository.INoteRepository) *INoteHandler {
	return &INoteHandler{repo}
}

func (h *INoteHandler) GetNote(context echo.Context) error {
	note, err := h.Repo.GetNote()

	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, note)
}

func (h *INoteHandler) InsertNote(context echo.Context) error {
	note := new(DTO.NoteDTO)
	// note := DTO.NoteDTO{}
	if err := context.Bind(note); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"errorBiing": "Invalid request"})
	}

	res, _ := h.Repo.InsertNote(model.Note{Title: note.Title, Content: note.Content})

	return context.JSON(http.StatusOK, res)
}
