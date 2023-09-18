package data

import (
	"github.com/jdstanhope/smalltown/internal/validator"
	"time"
)

type Photo struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Comment    string    `json:"comment"`
	StorageURL string    `json:"storageURL"`
	UserID     int64     `json:"userID"`
}

func ValidatePhoto(v *validator.Validator, photo *Photo) {
	v.Check(photo.Comment != "", "comment", "must be provided")
	v.Check(len(photo.Comment) <= 500, "comment", "must not be more than 500 bytes long")
}
