package dbhelper

import (
	"admin/middleware/global"
)

func GetDataById(id int32, val interface{}) error {
	dbConn := global.GetDbConn()
	db := dbConn.Conn.Where("id=?", id).Find(val)
	return db.Error
}

func GetDataByCond(cond interface{}, val interface{}) error {
	dbConn := global.GetDbConn()
	db := dbConn.Conn.Where(cond).Find(val)
	return db.Error
}
