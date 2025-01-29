package query

import (
	"database/sql"
	"rest_imageboard/internal/storage"
)

type User struct {
	Name     string	`json:"name"`
	Password string `json:"password"`
}

func UserIsInDB(db *sql.DB, user *User) (bool, error) {
	stmt, err := db.Prepare(`
		SELECT password FROM users 
		WHERE name = $1
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var password string
	err = stmt.QueryRow(user.Name).Scan(&password)
	if err == sql.ErrNoRows {
		return false, nil
	} 
	if err != nil {
		return false, err
	}
	if password != user.Password {
		return false, storage.ErrPasswordDoesNotMatch
	}
	return true, nil
}

func CreateUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare(`
		INSERT INTO users (name, password)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil 
}

func DeleteUser(db *sql.DB, user User) error {
	stmt, err := db.Prepare(`
		DELETE FROM users
		WHERE name = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name)
	if err != nil {
		return err
	}

	return nil
}