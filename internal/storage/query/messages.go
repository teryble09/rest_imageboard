package query

import (
	"database/sql"
	//"rest_imageboard/internal/storage"
	"time"
)

type Message struct {
	id          int
	authorName  string
	threadName  string
	createdAt   time.Time
	referenceId int
	text        string
	imageName   string
}

func SaveMessage(db *sql.DB) error {
	return nil
}

func DeleteMessage(db *sql.DB) error {
	return nil
}

func GetMessages(db *sql.DB, thr Thread) error {
	return nil
}
