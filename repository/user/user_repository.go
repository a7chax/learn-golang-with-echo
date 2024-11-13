package repository

import "echo-golang/model"

type IUserRepository interface {
	GetUser() ([]model.User, error)
	LoginUser(username string, password string) (model.User, error)
}
