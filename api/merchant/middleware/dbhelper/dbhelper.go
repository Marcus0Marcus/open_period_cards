package dbhelper

import (
	"github.com/jinzhu/gorm"
	"merchant/middleware/constant"
	"merchant/middleware/global"
	"merchant/middleware/response"
)

func GetDataById(id int32, data interface{}) *response.FWError {
	dbConn := global.GetDbConn()
	db := dbConn.Conn.Where("id=?", id).Find(data)
	if db.Error != nil && !gorm.IsRecordNotFoundError(db.Error) {
		return constant.ErrDb
	}
	if gorm.IsRecordNotFoundError(db.Error) {
		return constant.ErrDBNoRecord
	}
	return nil
}

func GetDataByCond(cond interface{}, data interface{}) *response.FWError {
	dbConn := global.GetDbConn()
	db := dbConn.Conn.Where(cond).Find(data)
	if db.Error != nil && !gorm.IsRecordNotFoundError(db.Error) {
		return constant.ErrDb
	}
	if gorm.IsRecordNotFoundError(db.Error) {
		return constant.ErrDBNoRecord
	}
	return nil
}

func CreateData(data interface{}) (*response.FWError, interface{}) {
	dbConn := global.GetDbConn()
	db := dbConn.Conn.Create(data)
	if db.Error != nil {
		return constant.ErrDb, nil
	}
	return nil, data
}

// id must be filled update other all fields
func UpdateData(data interface{}) (*response.FWError, int64) {
	dbConn := global.GetDbConn()
	db := dbConn.Conn.Model(data).Update(data)
	if db.Error != nil {
		return constant.ErrDb, 0
	}
	return nil, dbConn.Conn.RowsAffected
}
