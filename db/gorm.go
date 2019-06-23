package db

import (
	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/util/log"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var ORM, Errs = GormInit()

// GormInit init gorm ORM.
func GormInit() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", config.PostgresDSL())
	// db, err := gorm.Open("mysql", config.MysqlDSL())
	//db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	// db.SingularTable(true)
	if config.Environment == "DEVELOPMENT" {
		db.LogMode(true)
	}
	log.CheckError(err)

	return db, err
}
