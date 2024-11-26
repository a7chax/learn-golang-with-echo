package repository

import (
	"database/sql"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
)

type IUserRepository interface {
	GetUser() ([]model_response.User, error)
	LoginUser(login model_request.Login) (model_response.User, error)
	RegisterUser(register model_request.Register) (sql.Result, error)
}
