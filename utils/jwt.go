package utils

import (
	model_response "echo-golang/model/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Role int    `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user model_response.User) *JwtCustomClaims {
	claims := &JwtCustomClaims{
		user.Username,
		user.IdUser,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
	}

	return claims
}
