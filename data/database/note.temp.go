package database

import (
	"echo-golang/model"
	"fmt"
)

func QueryGetNote() ([]model.Note, error) {
	// Query
	db, _ := ConnectDatabaseNote()

	defer db.Close()

	rows, err := db.Query("SELECT * FROM note")

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer rows.Close()

	var result []model.Note

	for rows.Next() {
		var each = model.Note{}
		err := rows.Scan(&each.IdNote, &each.Title, &each.Content, &each.Date_created, &each.Date_updated)

		if err != nil {
			fmt.Printf(err.Error())
			return nil, fmt.Errorf(err.Error())
		}

		result = append(result, each)
	}
	return result, nil
}
