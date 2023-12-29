package viper

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("../../.")
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
