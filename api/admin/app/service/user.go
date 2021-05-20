package service

import "admin/middleware/global"

type UserService struct {

}
type UserInfo struct {
	Phone string        `json:"phone";gorm="phone"`
	Name string         `json:"name"`
	Pwd string          `json:"pwd"`
	Salt string         `json:"salt"`
	Mtime int32         `json:"mtime"`
	Ctime int32         `json:"ctime"`
	Deleted int32       `json:"deleted"`
	
}
func (r *UserService) GetUserById(id int32) *UserInfo{
	dbConn := global.GetDbConn()
	userInfo := &UserInfo{}
	db := dbConn.Conn.First(userInfo,id)
	if db.RowsAffected == 0 {
		return nil
	}
	return userInfo
	
}