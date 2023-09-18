package data

import (
	"github.com/jdstanhope/smalltown/internal/validator"
	"time"
)

type Page struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	UserID     int64     `json:"userID"`
	StorageURL string    `json:"storageURL"`
	PhotoID    int64     `json:"photoID"`
}

func ValidatePage(v *validator.Validator, page *Page) {
	v.Check(page.Name != "", "name", "must be provided")
	v.Check(len(page.Name) <= 500, "name", "must not be more than 500 bytes long")
}
