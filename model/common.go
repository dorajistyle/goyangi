package model

import (
	"time"
)

// omit is the bool type for omitting a field of struct.
type omit bool

// CommentList is list that contains comments and meta.
type CommentList struct {
	HasPrev     bool      `json:"hasPrev"`
	HasNext     bool      `json:"hasNext"`
	Count       int       `json:"count"`
	CurrentPage int       `json:"currentPage"`
	Comments    []Comment `json:"comments"`
}

// Comment is a comment model.
type Comment struct {
	Id          int64      `json:"id"`
	Content     string     `json:"content"`
	UserId      int64      `json:"userId"`
	LikingCount int        `json:"likingCount"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   time.Time  `json:"deletedAt"`
	User        PublicUser `json:"user"`
}

// LikingList is list that contains likings and meta.
type LikingList struct {
	HasPrev     bool          `json:"hasPrev"`
	HasNext     bool          `json:"hasNext"`
	Count       int           `json:"count"`
	CurrentPage int           `json:"currentPage"`
	IsLiked     bool          `json:"isLiked"`
	Likings     []*PublicUser `json:"likings"`
}

// LikedList is list that contains liked and meta.
type LikedList struct {
	HasPrev     bool          `json:"hasPrev"`
	HasNext     bool          `json:"hasNext"`
	Count       int           `json:"count"`
	CurrentPage int           `json:"currentPage"`
	Liked       []*PublicUser `json:"liked"`
}

// Image is a image model.
type Image struct {
	Id        int64     `json:"id"`
	Kind      int       `json:"kind"`
	Large     string    `json:"large"`
	Medium    string    `json:"medium"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"createdAt"`
}

// Tag is a tag model.
type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name",sql:"size:255"`
}

// Link is a link model.
type Link struct {
	Id        int64     `json:"id"`
	Kind      int       `json:"kind"`
	Name      string    `json:"title",sql:"size:255"`
	Url       string    `json:"url",sql:"size:512"`
	CreatedAt time.Time `json:"createdAt"`
	Icon      string    `json:"icon",sql:"size:255"`
}
