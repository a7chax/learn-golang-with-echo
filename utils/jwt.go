package utils

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Admin bool   `json:"admin"`
	Jwt   jwt.RegisteredClaims
}
