package model

import (
	"time"
)

// Location is a location model.
type Location struct {
	Id             int64       `json:"id"`
	Name           string      `json:"name",sql:"size:255"`
	Url            string      `json:"url",sql:"size:512"`
	Content        string      `json:"content"`
	Address        string      `json:"address"`
	Latitude       float64     `json:"latitude"`
	Longitude      float64     `json:"longitude"`
	Type           string      `json:"type"`
	UserId         int64       `json:"userId"`
	Author         PublicUser  `json:"author"`
	ReferralId     int64       `json:"referralId"`
	ReferralUserId int64       `json:"referralUserId"`
	PrevId         int64       `json:"prevId"`
	NextId         int64       `json:"nextId"`
	LikingCount    int         `json:"likingCount"`
	CommentCount   int         `json:"commentCount"`
	SharingCount   int         `json:"sharingCount"`
	Activate       bool        `json:"active"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	DeletedAt      time.Time   `json:"deletedAt"`
	Comments       []Comment   `gorm:"many2many:locations_comments;"`
	CommentList    CommentList `json:"commentList"`
	Likings        []User      `gorm:"many2many:locations_users;"`
	LikingList     LikingList  `json:"likingList"`
}
