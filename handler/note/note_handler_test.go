package handler

import (
	"echo-golang/model"
	repository "echo-golang/repository/note"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var noteRepositoryMock = &repository.INoteRepositoryMock{Mock: mock.Mock{}}
var noteHandlerMock = INoteHandler{noteRepositoryMock}

func TestNoteHandler_GetAllNote(t *testing.T) {
	note := []model.Note{
		{
			IdNote:       1,
			Title:        "Title 1",
			Content:      "Content 1",
			Date_created: "2021-08-01",
			Date_updated: "2021-08-01",
		},
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")
	noteRepositoryMock.Mock.On("GetNote").Return(note, nil)

	jsonnya, _ := json.Marshal(note)
	if assert.NoError(t, noteHandlerMock.GetNote(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(jsonnya)+"\n", rec.Body.String())
	}
}

func TestNoteHandler_InsertNote(t *testing.T) {
	noteRepositoryMock.Mock.On("InsertNote").Return(nil, nil)
	noteHandlerMock.InsertNote(nil)
	noteRepositoryMock.Mock.AssertExpectations(t)
}

func TestNoteHandler_DeleteNoteById(t *testing.T) {
	noteRepositoryMock.Mock.On("DeleteNoteById").Return(nil, nil)
	noteHandlerMock.DeleteNoteById(nil)
	noteRepositoryMock.Mock.AssertExpectations(t)
}

func TestNoteHandler_UpdateNoteById(t *testing.T) {
	noteRepositoryMock.Mock.On("UpdateNoteById").Return(nil, nil)
	noteHandlerMock.UpdateNoteById(nil)
	noteRepositoryMock.Mock.AssertExpectations(t)
}
