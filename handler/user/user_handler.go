package handler

import (
	"echo-golang/model"
	model_request "echo-golang/model/request"
	service "echo-golang/service/user"
	"echo-golang/validators"
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
	var login model_request.Login

	err := context.Bind(&login)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	validator := validators.New()

	if err = validator.Validate(login); err != nil {
		return context.JSON(http.StatusBadRequest, model.BaseResponseNoData{
			IsSuccess: false,
			Message:   err.Error(),
		})
	}

	response, err := h.service.LoginUser(login)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, response)

}

func (h *IUserHandler) RefreshToken(context echo.Context) error {
	token := context.Request().Header.Get("Authorization")
	response, err := h.service.RefreshToken(token)
	fmt.Println(response, "tokennya", err)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.BaseResponseNoData{
			IsSuccess: false,
			Message:   err.Error(),
		})
	}
	return context.JSON(http.StatusOK, response)
}
