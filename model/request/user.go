package model_request

type Login struct {
	Username string `form:"username" validate:"required,min=1,max=100"`
	Password string `form:"password" validate:"required,min=4,max=100"`
}

type Register struct {
	Username string `form:"username" validate:"required,min=1,max=100"`
	Password string `form:"password" validate:"required,min=4,max=100"`
	Email    string `form:"email" validate:"required,email"`
	Avatar   string `form:"avatar"`
	Role     int    `form:"role" validate:"required"`
}
