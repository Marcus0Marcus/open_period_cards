package db

import (
	"github.com/go-chassis/openlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// sudo docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
type DbConn struct {
	Conn *gorm.DB
}

func NewDBConn(dsn string, debug bool) *DbConn {
	// dsn example : root:XXXX@tcp(127.0.0.1:3306)/test
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		openlog.Fatal("init MysqlDB failed. " + err.Error())
	}
	conn.LogMode(debug)
	return &DbConn{
		Conn: conn,
	}
}
