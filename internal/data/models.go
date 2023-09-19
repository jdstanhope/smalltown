package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Model struct {
	Photos PhotoModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		Photos: PhotoModel{DB: db},
	}
}
