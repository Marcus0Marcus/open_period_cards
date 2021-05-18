package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-chassis/openlog"
)
type Cache struct{
	Conn redis.Conn
}
func NewCacheConn(dsn string) *Cache{
	
	conn, err := redis.Dial("tcp", dsn)
	if err != nil {
		openlog.Fatal("init Redis failed. " + err.Error())
	}
	return &Cache{
		Conn:conn,
	}
}
