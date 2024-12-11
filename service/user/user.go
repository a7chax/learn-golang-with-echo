package service

import (
	"context"
	"echo-golang/model"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
	repository "echo-golang/repository/user"
	"echo-golang/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type IUserService interface {
	GetAllUser() ([]model_response.User, error)
	GetUser(id int) (model.BaseResponse[model_response.User], error)
	LoginUser(login model_request.Login) (model.BaseResponse[string], error)
	RefreshToken(token string) (model.BaseResponse[string], error)
	RegisterUser(register model_request.Register) (model.BaseResponse[string], error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetAllUser() ([]model_response.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetUser(id int) (model.BaseResponse[model_response.User], error) {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis address
	})
	ctx := context.Background()

	idUser := strconv.Itoa(id)
	cachedUser, err := rdb.Get(ctx, idUser).Result()
	if err == nil {
		return model.BaseResponse[model_response.User]{
			IsSuccess: true,
			Message:   "Get user success",
			Data:      &model_response.User{Username: cachedUser}, // Assuming cachedUser is a JSON string and you need to unmarshal it to User struct
		}, nil
	}

	user, err := s.repo.GetUser(id)
	if err != nil {
		return model.BaseResponse[model_response.User]{
			IsSuccess: false,
			Message:   "Failed to get user",
			Data:      nil,
		}, err
	}

	err = rdb.Set(ctx, idUser, "userJSON", 0).Err()
	if err != nil {
		return model.BaseResponse[model_response.User]{
			IsSuccess: false,
			Message:   "Failed to store user in Redis",
			Data:      nil,
		}, err
	}

	return model.BaseResponse[model_response.User]{
		IsSuccess: true,
		Message:   "Get user success",
		Data:      &user,
	}, nil
}

func (s *UserService) LoginUser(login model_request.Login) (model.BaseResponse[string], error) {

	user, err := s.repo.LoginUser(login)

	decrypted := utils.DecryptPassword(login.Password, user.Password)

	if user.Username == login.Username && decrypted == nil {
		claims := utils.GenerateJWT(user.Username, user.IdUser)

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

func (s *UserService) RefreshToken(token string) (model.BaseResponse[string], error) {
	claims := &utils.JwtCustomClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "Token invalid",
			Data:      nil,
		}, err
	}

	if !tkn.Valid {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "Token invalid",
			Data:      nil,
		}, err
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := newToken.SignedString([]byte("secret"))

	return model.BaseResponse[string]{
		IsSuccess: true,
		Message:   "Refresh token success",
		Data:      &t,
	}, nil
}

func (s *UserService) RegisterUser(register model_request.Register) (model.BaseResponse[string], error) {
	encryptedPassword, err := utils.EncryptPassword(register.Password)

	if err != nil {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "Failed to encrypt password",
			Data:      nil,
		}, err
	}

	register.Password = encryptedPassword

	_, err = s.repo.RegisterUser(register)

	if err != nil {
		return model.BaseResponse[string]{
			IsSuccess: false,
			Message:   "Failed to register user",
			Data:      nil,
		}, err
	}

	return model.BaseResponse[string]{
		IsSuccess: true,
		Message:   "Register user success",
		Data:      nil,
	}, nil
}
