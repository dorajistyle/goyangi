package config

import (
	"github.com/dorajistyle/goyangi/config"
)

const (
	redisAddrDevelopment = ""
	redisPortDevelopment = "6379"
	redisAddrTest        = "goyangi-redis.ohpuh0.0001.usw1.cache.amazonaws.com"
	redisPortTest        = "6379"
	redisAddrProduction  = "goyangi-redis.ohpuh0.0001.usw1.cache.amazonaws.com"
	redisPortProduction  = "6379"
	Capacity             = 1
	MaxCap               = 2
)

// RedisAddr return redis address.
func RedisAddr() string {
	var redisAddr string
	switch config.Environment {
	case "DEVELOPMENT":
		redisAddr = redisAddrDevelopment + ":" + redisPortDevelopment
		// postgresDSL = PostgresDSLTest
		// postgresDSL = PostgresDSLProduction
	case "TEST":
		redisAddr = redisAddrTest + ":" + redisPortTest
	default:
		redisAddr = redisAddrProduction + ":" + redisPortProduction
	}
	return redisAddr
}
