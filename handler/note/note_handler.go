package handler

import (
	model_request "echo-golang/model/request"
	service "echo-golang/service/note"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type INoteHandler struct {
	service service.INoteService
}

func NoteHandler(service service.INoteService) *INoteHandler {
	return &INoteHandler{service}
}

func (h *INoteHandler) GetNote(context echo.Context) error {
	note, err := h.service.GetAllNote()

	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, note)
}

func (h *INoteHandler) InsertNote(context echo.Context) error {
	note := new(model_request.Note)
	if err := context.Bind(note); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"errorBiing": "Invalid request"})
	}

	res, _ := h.service.InsertNote(model_request.Note{
		Title:   note.Title,
		Content: note.Content,
	})

	return context.JSON(http.StatusOK, res)
}

func (h *INoteHandler) DeleteNoteById(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	res, _ := h.service.DeleteNoteById(idInt)

	return context.JSON(http.StatusOK, res)
}

func (h *INoteHandler) UpdateNoteById(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	note := new(model_request.Note)
	if err := context.Bind(note); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"errorBiing": "Invalid request"})
	}
	res, _ := h.service.UpdateNoteById(idInt, model_request.Note{Title: note.Title, Content: note.Content})

	return context.JSON(http.StatusOK, res)
}
