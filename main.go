package main

import (
	"log"
	"os"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/script"
	"github.com/dorajistyle/goyangi/server"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// switch os.Getenv("ENV") {
	// case "DEVELOPMENT":
	// 	viper.SetConfigName(".env.dev")
	// case "TEST":
	// 	viper.SetConfigName(".env.test")
	// case "PRODUCTION":
	// 	viper.SetConfigName(".env.prod")
	// }

	viper.SetConfigName("config.yml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Fatal error config file", err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Goyangi script tool"
	app.Usage = "run scripts!"
	app.Version = "0.1.0"

	app.Author = "https://github.com/dorajistyle(JoongSeob Vito Kim)"
	app.Commands = script.Commands()
	app.Action = func(c *cli.Context) {
		println("Run Server.")
		println(viper.GetString("app.name"))

		server.Run()
	}
	println(viper.GetString("database.type"))
	db.GormInit()

	app.Run(os.Args)
}
