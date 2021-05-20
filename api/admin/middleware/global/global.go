package global

import (
	"admin/middleware/cache"
	"admin/middleware/conf"
	"admin/middleware/db"
)

type Global struct {
	Conf      *conf.Conf
	CacheConn *cache.Cache
	DbConn    *db.DbConn
}

var globalInfo *Global

func GetDbConn() *db.DbConn {
	if globalInfo == nil {
		return nil
	}
	return globalInfo.DbConn
}
func GetCacheConn() *cache.Cache {
	if globalInfo == nil {
		return nil
	}
	return globalInfo.CacheConn
}
func GetConfig() *conf.Conf {
	if globalInfo == nil {
		return nil
	}
	return globalInfo.Conf
}
func NewGlobal() *Global {
	// load json config
	jsonConf := conf.LoadConfig()

	// init db
	dbConn := db.NewDBConn(jsonConf.Config.Mysql.DSN, jsonConf.Config.Mysql.Debug)

	// init cache
	cacheConn := cache.NewCacheConn(jsonConf.Config.Redis.DSN)
	globalInfo = new(Global)
	globalInfo.Conf = jsonConf
	globalInfo.CacheConn = cacheConn
	globalInfo.DbConn = dbConn
	return globalInfo
}

func ClearGlobal(global *Global) {
	defer global.DbConn.Conn.Close()
	defer global.CacheConn.Conn.Close()
}
