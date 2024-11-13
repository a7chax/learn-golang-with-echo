package middleware

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, d echo.Context) (bool, error) {

		loginUrl := d.Request().URL.Path == "/login"
		basicAuthUsername := subtle.ConstantTimeCompare([]byte(username), []byte("arung")) == 1
		basicAuthPassword := subtle.ConstantTimeCompare([]byte(password), []byte("12345")) == 1
		if loginUrl {
			if basicAuthUsername && basicAuthPassword {
				return true, nil
			} else {
				return false, nil
			}
		}
		return true, nil
	})
}
