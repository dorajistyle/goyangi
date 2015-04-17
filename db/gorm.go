package db

import (
	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var ORM, Errs = GormInit()

// GormInit init gorm ORM.
func GormInit() (gorm.DB, error) {
	db, err := gorm.Open("mysql", config.MysqlDSL())
	//db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	db.SingularTable(true)
	if config.Environment == "DEVELOPMENT" {
		// db.LogMode(true)
		db.AutoMigrate(&model.User{}, &model.Role{}, &model.Connection{}, &model.Language{}, &model.Article{}, &model.Location{}, &model.Comment{}, &model.File{})
		db.Model(&model.User{}).AddIndex("idx_user_token", "token")
	}
	log.CheckError(err)

	return db, err
}
