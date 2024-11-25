package repository

import (
	"database/sql"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
)

type userRepository struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser() ([]model_response.User, error) {
	var result []model_response.User
	query := "SELECT * FROM note_user"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		each := model_response.User{}
		if err := rows.Scan(&each.IdUser, &each.Username, &each.Email, &each.Password); err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}

func (r *userRepository) LoginUser(login model_request.Login) (model_response.User, error) {
	var result model_response.User
	query := "SELECT * FROM note_user WHERE username = $1 AND password = $2"
	err := r.db.QueryRow(query, login.Username, login.Password).Scan(&result.IdUser, &result.Username, &result.Password, &result.Email)

	if err != nil {
		return model_response.User{}, err
	}
	return result, nil
}
