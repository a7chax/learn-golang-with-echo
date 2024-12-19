package middleware

import (
	"echo-golang/constant"
	"echo-golang/model"
	"echo-golang/utils"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func OnlyAdmin() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			claims := &utils.JwtCustomClaims{}

			tokenBearer := authHeader[len("Bearer "):]
			if tokenBearer == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing invalid token")
			}

			token, err := jwt.ParseWithClaims(tokenBearer, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token1")
			}

			claims, ok := token.Claims.(*utils.JwtCustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid invalid claims")
			}

			if claims.Role == constant.ROLE_ADMIN {
				return echo.NewHTTPError(http.StatusForbidden, model.BaseResponseNoData{
					IsSuccess: false,
					Message:   "Only Admin can access this endpoint",
				})
			}

			return next(c)
		}
	}
}
