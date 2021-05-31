package data_service

type CardTypeInfoTplInfo struct {
	Id         int64  `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	CardTypeId int32  `json:"card_type_id" gorm:"column:card_type_id"`
	Tpl        string `json:"tpl" gorm:"column:tpl"`
	Describe   string `json:"describe" gorm:"column:describe"`
	Mtime      uint32 `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime      uint32 `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted    int32  `json:"deleted" gorm:"column:deleted"`
}

func (ui *CardTypeInfoTplInfo) TableName() string {
	return "tb_card_type_info_tpl"
}
