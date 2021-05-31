package cachehelper

import (
	"github.com/garyburd/redigo/redis"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/global"
	"open_period_cards/middleware/response"
)

func KeySet(name string, val string) *response.FWError {
	cache := global.GetCacheConn().Conn.Get()
	_, err := cache.Do("SET", name, val)
	if err != nil {
		return constant.ErrCacheSet
	}
	return nil
}
func KeyGet(name string) (*response.FWError, string) {
	cache := global.GetCacheConn().Conn.Get()
	val, err := redis.String(cache.Do("GET", name))
	if err == redis.ErrNil {
		return constant.ErrCacheNotExist, ""
	}
	if err != nil {
		return constant.ErrCacheGet, ""
	}
	return nil, val
}

func KeyDel(name string) *response.FWError {
	cache := global.GetCacheConn().Conn.Get()
	_, err := cache.Do("DEL", name)
	if err == redis.ErrNil {
		return constant.ErrCacheNotExist
	}
	if err != nil {
		return constant.ErrCacheDel
	}
	return nil
}
