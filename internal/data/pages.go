package data

import "time"

type Page struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	UserID     int64     `json:"userID"`
	StorageURL string    `json:"storageURL"`
	PhotoID    int64     `json:"photoID"`
}
