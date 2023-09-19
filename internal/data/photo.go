package data

import (
	"database/sql"
	"errors"
	"github.com/jdstanhope/smalltown/internal/validator"
	"time"
)

type Photo struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	StorageURL string    `json:"storageURL"`
	UserID     int64     `json:"userID"`
	Version    int32     `json:"version"`
}

func ValidatePhoto(v *validator.Validator, photo *Photo) {
	v.Check(photo.Name != "", "comment", "must be provided")
	v.Check(len(photo.Name) <= 500, "comment", "must not be more than 500 bytes long")
}

type PhotoModel struct {
	DB *sql.DB
}

func (model PhotoModel) Insert(photo *Photo) error {
	query := `
		INSERT INTO photos (name)
		VALUES ($1)
		RETURNING id, created_at, version`

	args := []interface{}{photo.Name}

	return model.DB.QueryRow(query, args...).Scan(&photo.ID, &photo.CreatedAt, &photo.Version)
}

func (model PhotoModel) Get(id int64) (*Photo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, created_at, name, version FROM photos WHERE id = $1`
	var photo Photo

	err := model.DB.QueryRow(query, id).Scan(&photo.ID, &photo.CreatedAt, &photo.Name, &photo.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &photo, nil
}

func (model PhotoModel) Update(userID int64) error {
	return nil
}

func (model PhotoModel) Delete(userID int64) error {
	return nil
}
