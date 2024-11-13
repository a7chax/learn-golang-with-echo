package repository

import (
	"database/sql"
	"echo-golang/model"
)

type userRepository struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser() ([]model.User, error) {
	var result []model.User
	query := "SELECT * FROM note_user"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		each := model.User{}
		if err := rows.Scan(&each.IdUser, &each.Username, &each.Email, &each.Password); err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}

func (r *userRepository) LoginUser(username string, password string) (model.User, error) {
	var result model.User
	query := "SELECT * FROM note_user WHERE username = $1 AND password = $2"
	err := r.db.QueryRow(query, username, password).Scan(&result.IdUser, &result.Username, &result.Password, &result.Email)

	if err != nil {
		return model.User{}, err
	}
	return result, nil
}
