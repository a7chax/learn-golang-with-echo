package repository

import (
	model_response "echo-golang/model/response"
)

type IUserRepository interface {
	GetUser() ([]model_response.User, error)
	LoginUser(username string, password string) (model_response.User, error)
}
