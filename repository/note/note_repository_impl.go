package repository

import (
	"database/sql"
	"echo-golang/model"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
)

type noteRepository struct {
	db *sql.DB
}

func NoteRepository(db *sql.DB) INoteRepository {
	return &noteRepository{db}
}

func (r *noteRepository) GetNote(pagination model.Pagination) ([]model_response.Note, error) {
	var result []model_response.Note
	query := `SELECT id_notes, title, content, date_created, date_updated FROM note ORDER BY date_created ASC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, pagination.Size, pagination.Page)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		each := model_response.Note{}
		if err := rows.Scan(&each.IdNote, &each.Title, &each.Content, &each.Date_created, &each.Date_updated); err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}

func (r *noteRepository) InsertNote(note model_request.Note) (sql.Result, error) {
	query := `INSERT INTO note (title, content) VALUES ($1, $2) RETURNING id_notes`
	execResult, err := r.db.Exec(query, note.Title, note.Content)

	return execResult, err
}

func (r *noteRepository) DeleteNoteById(id int) (sql.Result, error) {
	query := `DELETE FROM note WHERE id_notes=$1`
	execResult, err := r.db.Exec(query, id)

	return execResult, err
}

func (r *noteRepository) UpdateNoteById(id int, note model_request.Note) (sql.Result, error) {
	query := `UPDATE note SET title=$1, content=$2 WHERE id_notes=$3`
	execResult, err := r.db.Exec(query, note.Title, note.Content, id)
	return execResult, err
}
