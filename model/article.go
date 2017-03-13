package model

import (
	"time"
)

// Article is a article model.
type Article struct {
	Id             uint        `json:"id"`
	Title          string      `json:"title",sql:"size:255"`
	Url            string      `json:"url",sql:"size:512"`
	Content        string      `json:"content"`
	UserId         uint        `json:"userId"`
	Author         User        `json:"author"`
	ReferralId     uint        `json:"referralId"`
	ReferralUserId uint        `json:"referralUserId"`
	CategoryId     int         `json:"categoryId"`
	PrevId         uint        `json:"prevId"`
	NextId         uint        `json:"nextId"`
	LikingCount    int         `json:"likingCount"`
	CommentCount   int         `json:"commentCount"`
	SharingCount   int         `json:"sharingCount"`
	ImageName      string      `json:"imageName",sql:"size:512"`
	ThumbnailName  string      `json:"thumbnailName",sql:"size:512"`
	Activate       bool        `json:"active"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	DeletedAt      *time.Time  `json:"deletedAt"`
	Comments       []Comment   `gorm:"many2many:articles_comments;"`
	CommentList    CommentList `json:"commentList"`
	Likings        []User      `gorm:"many2many:articles_users;"`
	LikingList     LikingList  `json:"likingList"`
}
