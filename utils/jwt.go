package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, idUser int) *JwtCustomClaims {
	claims := &JwtCustomClaims{
		username,
		idUser,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		},
	}

	return claims
}
