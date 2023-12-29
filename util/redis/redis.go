package redis

import (
	// "github.com/dorajistyle/goyangi/util/log"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/rueian/rueidis"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func GetClient() rueidis.Client {
	redisAddr := viper.GetString("redis.addr") + ":" + viper.GetString("redis.port")
	options := rueidis.ClientOption{
		InitAddress: []string{redisAddr},
	}
	client, err := rueidis.NewClient(options)
	if err != nil {
		log.Error("Cannot get ruedis client.", err)
	}
	return client
}

func Append(key string, value string) error {
	client := GetClient()
	defer client.Close()
	cmd := client.B().Set().Key(key).Value(value).Build()
	res := client.Do(context.Background(), cmd)
	return res.Error()
}

func Get(key string) (value string, err error) {
	client := GetClient()
	defer client.Close()
	cmd := client.B().Get().Key(key).Build()
	res := client.Do(context.Background(), cmd)
	value, _ = res.ToString()
	return value, res.Error()
}

func Del(key string) error {
	client := GetClient()
	defer client.Close()
	cmd := client.B().Del().Key(key).Build()
	res := client.Do(context.Background(), cmd)
	return res.Error()
}
