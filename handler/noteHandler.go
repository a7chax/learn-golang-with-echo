package handler

import (
	DTO "echo-golang/dto"
	"echo-golang/model"
	"echo-golang/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type noteHandler struct {
	repo repository.INoteRepository
}

func NoteHandler(repo repository.INoteRepository) *noteHandler {
	return &noteHandler{repo}
}

func (h *noteHandler) GetNote(context echo.Context) error {
	notes, err := h.repo.GetNote()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, notes)
}

func (h *noteHandler) InsertNote(context echo.Context) error {
	note := new(DTO.NoteDTO)
	if err := context.Bind(note); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"errorBiing": "Invalid request"})
	}

	// if err := context.Validate(note); err != nil {
	// 	return context.JSON(http.StatusBadRequest, map[string]string{"errorValidate": err.Error()})
	// }

	res, _ := h.repo.InsertNote(model.Note{Title: note.Title, Content: note.Content})

	return context.JSON(http.StatusOK, res)
}
