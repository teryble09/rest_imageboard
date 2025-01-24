package query

import (
	"database/sql"
	"rest_imageboard/internal/storage"
)

type Thread struct {
	Name string `json:"name"`
}

func SaveThread(db *sql.DB, thr Thread) error {
	stmt, err := db.Prepare(`
		INSERT INTO threads (name)
		VALUES ($1)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(thr.Name)
	if err != nil {
		return err
	}

	return nil
}

func DeleteThread(db *sql.DB, thr Thread) error {
	stmt, err := db.Prepare(`
		DELETE FROM threads
		WHERE name = $1
	`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(thr.Name)
	if err != nil {
		return err
	}

	aff_row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff_row == 0 {
		return storage.ErrThreadNotFound
	}

	return nil
}

func GetThreads(db *sql.DB) ([]Thread, error) {
	rows, err := db.Query(`SELECT * FROM threads`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Thread{}
	thr := Thread{}
	for rows.Next() {
		err := rows.Scan(&thr.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, thr)
	}
	return result, nil
}
