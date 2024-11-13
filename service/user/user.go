package service

import (
	"echo-golang/model"
	repository "echo-golang/repository/user"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IUserService interface {
	GetAllUser() ([]model.User, error)
	LoginUser(username string, password string) (model.BaseResponse[string], error)
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetAllUser() ([]model.User, error) {
	return s.repo.GetUser()
}

func (s *UserService) LoginUser(username string, password string) (model.BaseResponse[string], error) {

	// Validate input
	if len(password) < 4 {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "Password must be at least 4 characters long",
			Data:      nil,
		}, errors.New("password must be at least 4 characters long")
	}
	if username == "" {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "username cannot be empty",
			Data:      nil,
		}, errors.New("username cannot be empty")
	}

	user, err := s.repo.LoginUser(username, password)

	if user.Username == username && user.Password == password {
		claims := &JwtCustomClaims{
			user.Username,
			true,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, _ := token.SignedString([]byte("secret"))

		return model.BaseResponse[string]{
			IsSuccess: true,
			Message:   "Login success",
			Data:      &t,
		}, nil
	} else {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "Login failed",
			Data:      nil,
		}, err
	}
}
