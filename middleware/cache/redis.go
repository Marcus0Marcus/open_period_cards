package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type Cache struct {
	Conn *redis.Pool
}

func NewCacheConn(dsn string) *Cache {
	return &Cache{
		Conn: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", dsn) },
		},
	}
}
