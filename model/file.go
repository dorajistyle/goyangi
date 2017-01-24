package model

import (
	"time"
)

// File is a file model that contains meta.
type File struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"userId"`
	Name      string    `json:"name",sql:"size:255"`
	Size      int       `json:"size",sql:"size:255"`
	CreatedAt time.Time `json:"createdAt"`
}
