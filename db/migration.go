package db

import (
	"time"
	"github.com/dorajistyle/goyangi/model"
)

func Migrate() {
	MigrateUserAndRole()
	ORM.AutoMigrate(&model.Article{}, &model.CommentList{}, &model.Comment{}, &model.LikingList{}, &model.LikedList{}, &model.Image{}, &model.Tag{}, &model.Link{}, &model.App{}, &model.ExpiredTokenLog{})
	ORM.AutoMigrate(&model.File{}, &model.Language{})
	ORM.AutoMigrate(&model.UsersFollowers{}, &model.Connection{})
}

func MigrateUserAndRole() {
	var adminRole model.Role
	var userRole model.Role

	hasRole := ORM.HasTable(&model.Role{})
	hasUser := ORM.HasTable(&model.User{})

	if !hasRole {
		ORM.CreateTable(&model.Role{})
		adminRole.Name = "admin"
		adminRole.Description = "administrator"
		ORM.Create(&adminRole)
		userRole.Name = "user"
		userRole.Description = "ordinary user"
		ORM.Create(&userRole)
	}

	if !hasUser {
		ORM.CreateTable(&model.User{})
		var user model.User
		user.Username = "admin"
		user.Email = "admin@goyangi.github.io"
		user.Password = "$2a$10$voqxhv08H2eWHbLJo2rEeO1GwGlg8ZLW3Y8348aqe0XBqVgEZxGOu" // password
		user.Name = "Goyangi"
		user.Birthday = time.Now()
		user.Gender = 2
		user.Md5 = "10d17498672e2dd040e8c0cf5a337a61"
		user.Activation = true
		user.Token = "168355cf5b6d31827c694260ab24e3bc3e990290ca94c7c30c6489ae1c1f212c"
		user.TokenExpiration = time.Now().Add(time.Hour * 24 * 365 * 100)
		ORM.Create(&user)
		ORM.Model(&user).Association("Roles").Append(adminRole)
		ORM.Model(&user).Association("Roles").Append(userRole)
	}
}