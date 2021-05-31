package data_service

type AdminInfo struct {
	Id      int64  `json:"id" gorm:"column:id"`
	Phone   string `json:"phone" gorm:"column:phone"`
	Name    string `json:"name" gorm:"column:name"`
	Pwd     string `json:"pwd" gorm:"column:pwd"`
	Salt    string `json:"salt" gorm:"column:salt"`
	Mtime   uint32 `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   uint32 `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`
}

func (ui *AdminInfo) TableName() string {
	return "tb_admin"
}
