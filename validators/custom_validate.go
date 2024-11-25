package validators

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

// Validator instance
type CustomValidator struct {
	validator *validator.Validate
}

// New creates a new instance of CustomValidator.
func New() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Username":
				if err.Tag() == "required" {
					errors["username"] = "Username is required"
				} else if err.Tag() == "min" {
					errors["username"] = "Username must be at least 1 character"
				} else if err.Tag() == "max" {
					errors["username"] = "Username must be at most 100 characters"
				}
			case "Password":
				if err.Tag() == "required" {
					errors["password"] = "Password is required"
				} else if err.Tag() == "min" {
					errors["password"] = "Password must be at least 4 characters"
				} else if err.Tag() == "max" {
					errors["password"] = "Password must be at most 100 characters"
				}
			}
		}
		return echo.NewHTTPError(http.StatusBadRequest, errors)
	}
	return nil
}
