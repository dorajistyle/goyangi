package db

import (
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/spf13/viper"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var ORM *gorm.DB

// GormInit init gorm ORM.
func GormInit() {
	appEnvironmnet := viper.GetString("app.environment")
	dbType := viper.GetString("database.type")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbDatabase := viper.GetString("database.database")
	dbOptions := viper.GetString("database.options")

	var databaseDSL string

	switch dbType {
	case "postgres":
		databaseDSL = "host=" + dbHost + " port=" + dbPort + " dbname=" + dbDatabase + " " + dbOptions + " user=" + dbUser + " password=" + dbPassword
	case "mysql":
		databaseDSL = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?" + dbOptions

	}
	println(databaseDSL)
	db, err := gorm.Open(dbType, databaseDSL)

	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	// db.SingularTable(true)
	if appEnvironmnet == "DEVELOPMENT" {
		db.LogMode(true)
	}
	log.CheckError(err)
	ORM = db
}
