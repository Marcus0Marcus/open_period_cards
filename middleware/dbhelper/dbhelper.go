package dbhelper

import (
	"github.com/jinzhu/gorm"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/global"
	"open_period_cards/middleware/response"
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
func GetPagedDataByCond(cond interface{}, data interface{}, pageNo int64, pageSize int64) (*response.FWError, int64) {
	dbConn := global.GetDbConn()
	// get total
	var total int64
	db := dbConn.Conn.Model(cond).Where(cond).Count(&total)
	if db.Error != nil {
		return constant.ErrDb, 0
	}

	// get paged data
	limit := pageSize
	offset := (pageNo - 1) * pageSize
	db = dbConn.Conn.Offset(offset).Limit(limit).Where(cond).Find(data)
	if db.Error != nil {
		return constant.ErrDb, 0
	}
	return nil, total
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
