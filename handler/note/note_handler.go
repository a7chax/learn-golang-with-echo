package handler

import (
	"echo-golang/model"
	model_request "echo-golang/model/request"
	service "echo-golang/service/note"
	"echo-golang/utils"
	"os"

	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type INoteHandler struct {
	service service.INoteService
}

func NoteHandler(service service.INoteService) *INoteHandler {
	return &INoteHandler{service}
}

func (h *INoteHandler) GetNote(context echo.Context) error {
	note, err := h.service.GetAllNote(model.Pagination{
		Page: 1,
		Size: 10,
	})

	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, note)
}

func (h *INoteHandler) InsertNote(context echo.Context) error {
	claims := &utils.JwtCustomClaims{}
	note := new(model_request.Note)
	token := context.Request().Header.Get("Authorization")
	secret := os.Getenv("JWT_SECRET")

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err := context.Bind(note); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"errorBiing": "Invalid request"})
	}

	res, _ := h.service.InsertNote(model_request.Note{
		Title:   note.Title,
		Content: note.Content,
		IdUser:  claims.Id,
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
