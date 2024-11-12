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
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
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

func (r *noteRepository) InsertNote(note model.Note) (model.BaseResponse[model.Note], error) {
	query := `INSERT INTO note (title, content) VALUES ($1, $2)`
	_, err := r.db.Exec(query, note.Title, note.Content)
	defer r.db.Close()

	if err != nil {
		return model.BaseResponse[model.Note]{Message: "Failed to insert note", Data: nil}, err
	}

	return model.BaseResponse[model.Note]{Message: "Succesful insert note", Data: nil}, nil
}

func (r *noteRepository) DeleteNoteById(id int) (model.BaseResponse[model.Note], error) {
	query := `DELETE FROM note WHERE id_notes=$1`
	_, err := r.db.Exec(query, id)
	defer r.db.Close()

	if err != nil {
		return model.BaseResponse[model.Note]{Message: "Failed to delete note", Data: nil}, err
	}

	return model.BaseResponse[model.Note]{Message: "Succesful delete note", Data: nil}, nil
}

func (r *noteRepository) UpdateNoteById(id int, note model.Note) (model.BaseResponse[model.Note], error) {
	query := `UPDATE note SET title=$1, content=$2 WHERE id_notes=$3`
	_, err := r.db.Exec(query, note.Title, note.Content, id)
	defer r.db.Close()
	if err != nil {
		return model.BaseResponse[model.Note]{Message: "Failed to update note", Data: nil}, err
	}

	return model.BaseResponse[model.Note]{Message: "Succesful update note", Data: nil}, nil
}
