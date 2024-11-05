package handler

import (
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
