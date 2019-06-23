package model

import (
	"time"
)

type App struct {
	Id uint `json:"id"`
	Name            string    `json:"name",sql:"size:255"`
	Key             string    `json:"name",sql:"size:255"`
	Token           string    `json:"token"`
	TokenExpiration time.Time `json:"tokenExperiation"`
}

type ExpiredTokenLog struct{
	UserId         		uint       `json:"userId"`
	AccessedAt        time.Time  `json:"activatedAt"`
}