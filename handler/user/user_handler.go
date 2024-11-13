package handler

import (
	service "echo-golang/service/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserHandler struct {
	service service.IUserService
}

func UserHandler(service service.IUserService) *IUserHandler {
	return &IUserHandler{service}
}

func (h *IUserHandler) GetAllUser(context echo.Context) error {
	user, err := h.service.GetAllUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

	}
	return context.JSON(http.StatusOK, user)
}

func (h *IUserHandler) LoginUser(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")

	user, err := h.service.LoginUser(username, password)
	fmt.Println(user)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, user)
}
