package storage

import "errors"

var (
	ErrThreadNotFound  = errors.New("thread not found")
	ErrMessageNotFound = errors.New("message not found")
)
