package redis

import (
	// "github.com/dorajistyle/goyangi/util/log"
	"time"

	"github.com/dorajistyle/goyangi/util/config"
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

var Pool, Resource, InitErr = RedisInit()

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

func RedisInit() (*pools.ResourcePool, ResourceConn, error) {
	var resourceConn ResourceConn
	pool := pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", config.RedisAddr())
		return ResourceConn{c}, err
	}, config.Capacity, config.MaxCap, time.Minute)
	// defer p.Close()
	ctx := context.TODO()
	r, err := pool.Get(ctx)
	if err != nil {
		// log.Fatal(err.Error())
		return pool, resourceConn, err
	}
	defer pool.Put(r)
	resourceConn = r.(ResourceConn)
	return pool, resourceConn, nil
}

func (*ResourceConn) Append(key string, value string) error {
	if InitErr != nil {
		return InitErr
	}
	_, err := Resource.Do("SET", key, value)
	return err
}

func (*ResourceConn) Get(key string) (string, error) {
	var value string
	var err error
	if InitErr != nil {
		return value, InitErr
	}
	value, err = redis.String(Resource.Do("GET", key))
	return value, err
}

func (*ResourceConn) Del(key string) error {
	if InitErr != nil {
		return InitErr
	}
	_, err := Resource.Do("DEL", key)
	return err
}
