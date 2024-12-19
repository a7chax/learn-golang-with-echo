package repository

import (
	"database/sql"
	"echo-golang/model"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]model_response.User, model.Metadata, error) {
	var result []model_response.User
	var resultMetaData model.Metadata
	query := "SELECT user_id, username, email, password FROM note_user"
	rows, err := r.db.Query(query)

	queryMetadata := "SELECT COUNT(*) AS total_size, CEIL(COUNT(*)::NUMERIC / 10) AS total_pages FROM note_user"
	errMetadata := r.db.QueryRow(queryMetadata).Scan(&resultMetaData.TotalSize, &resultMetaData.TotalPages)

	if errMetadata != nil {
		fmt.Println("Error get metadata")
		return nil, model.Metadata{}, errMetadata
	}

	if err != nil {
		return nil, model.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		each := model_response.User{}
		if err := rows.Scan(&each.IdUser, &each.Username, &each.Email, &each.Password); err != nil {
			return nil, model.Metadata{}, err
		}
		result = append(result, each)
	}
	return result, model.Metadata{
			TotalSize:  resultMetaData.TotalSize,
			TotalPages: resultMetaData.TotalPages,
		},
		nil
}

func (r *userRepository) GetUser(id int) (model_response.User, error) {
	var result model_response.User
	query := "SELECT user_id, username, password, email FROM note_user WHERE user_id = $1"
	err := r.db.QueryRow(query, id).Scan(&result.IdUser, &result.Username, &result.Password, &result.Email)

	if err != nil {
		return model_response.User{}, err
	}
	return result, nil
}

func (r *userRepository) LoginUser(login model_request.Login) (model_response.User, error) {
	var result model_response.User
	query := "SELECT user_id, username, password, email, role FROM note_user WHERE username = $1"
	err := r.db.QueryRow(query, login.Username).Scan(&result.IdUser, &result.Username, &result.Password, &result.Email, &result.Role)

	if err != nil {
		return model_response.User{}, err
	}
	return result, nil
}

func (r *userRepository) RegisterUser(register model_request.Register) (sql.Result, error) {
	query := `INSERT INTO note_user (username, password, email, role) VALUES ($1, $2, $3, $4) RETURNING user_id`
	execResult, err := r.db.Exec(query, register.Username, register.Password, register.Email, register.Role)

	return execResult, err
}
