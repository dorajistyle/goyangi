package model

import (
	"time"
)

// Article is a article model.
type Article struct {
	Id             int64       `json:"id"`
	Title          string      `json:"title",sql:"size:255"`
	Url            string      `json:"url",sql:"size:512"`
	Content        string      `json:"content"`
	UserId         int64       `json:"userId"`
	Author         PublicUser  `json:"author"`
	ReferralId     int64       `json:"referralId"`
	ReferralUserId int64       `json:"referralUserId"`
	CategoryId     int         `json:"categoryId"`
	PrevId         int64       `json:"prevId"`
	NextId         int64       `json:"nextId"`
	LikingCount    int         `json:"likingCount"`
	CommentCount   int         `json:"commentCount"`
	SharingCount   int         `json:"sharingCount"`
	ImageName      string      `json:"imageName",sql:"size:512"`
	ThumbnailName  string      `json:"thumbnailName",sql:"size:512"`
	Activate       bool        `json:"active"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	DeletedAt      time.Time   `json:"deletedAt"`
	Comments       []Comment   `gorm:"many2many:articles_comments;"`
	CommentList    CommentList `json:"commentList"`
	Likings        []User      `gorm:"many2many:articles_users;"`
	LikingList     LikingList  `json:"likingList"`
}
