package model

type User struct {
	IdUser   int    `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
