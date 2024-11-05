package repository

import (
	"database/sql"
	"echo-golang/model"
)

type noteRepository struct {
	db *sql.DB
}

func NoteRepository(db *sql.DB) INoteRepository {
	return &noteRepository{db}
}

func (r *noteRepository) GetNote() ([]model.Note, error) {
	var result []model.Note
	query := "SELECT * FROM note"
	rows, _ := r.db.Query(query)
	defer rows.Close()

	for rows.Next() {
		each := model.Note{}
		if err := rows.Scan(&each.IdNote, &each.Title, &each.Content, &each.Date_created, &each.Date_updated); err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}
