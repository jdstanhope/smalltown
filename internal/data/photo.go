package data

import "time"

type Photo struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Comment    string    `json:"comment"`
	StorageURL string    `json:"storageURL"`
	UserID     int64     `json:"userID"`
}
