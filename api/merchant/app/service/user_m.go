package service

type UserInfo struct {
	Phone   string `json:"phone" gorm:"column:phone"`
	Name    string `json:"name" gorm:"column:name"`
	Pwd     string `json:"pwd" gorm:"column:pwd"`
	Salt    string `json:"salt" gorm:"column:salt"`
	Mtime   int32  `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   int32  `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`
}

func (ui *UserInfo) TableName() string {
	return "tb_user"
}
