package handler

import (
	DTO "echo-golang/dto"
	"echo-golang/model"
	repository "echo-golang/repository/note"
	"net/http"
	"strconv"

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

func (h *INoteHandler) DeleteNoteById(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	res, _ := h.Repo.DeleteNoteById(idInt)

	return context.JSON(http.StatusOK, res)
}

func (h *INoteHandler) UpdateNoteById(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	note := new(DTO.NoteDTO)
	if err := context.Bind(note); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"errorBiing": "Invalid request"})
	}
	res, _ := h.Repo.UpdateNoteById(idInt, model.Note{Title: note.Title, Content: note.Content})

	return context.JSON(http.StatusOK, res)
}
