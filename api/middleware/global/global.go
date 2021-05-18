package global

import (
	"middleware/cache"
	"middleware/conf"
	"middleware/db"
)

type Global struct {
	Conf *conf.Conf
	CacheConn *cache.Cache
	DbConn *db.DbConn
}
func NewGlobal() *Global{
	// load json config
	jsonConf := conf.LoadConfig()
	
	// init db
	dbConn := db.NewDBConn(jsonConf.Config.Mysql.DSN)
	
	// init cache
	cacheConn := cache.NewCacheConn(jsonConf.Config.Redis.DSN)
	return &Global{
		Conf : jsonConf,
		CacheConn : cacheConn,
		DbConn : dbConn,
	}
}

func ClearGlobal(global *Global)  {
	defer global.DbConn.Conn.Close()
	defer global.CacheConn.Conn.Close()
}
