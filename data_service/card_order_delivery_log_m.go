package data_service

type CardOrderDeliveryLogInfo struct {
	Id          int64  `json:"id" gorm:"column:id"`
	CardOrderId int32  `json:"card_order_id" gorm:"column:card_order_id"`
	Content     string `json:"content" gorm:"column:content"`
	Describe    string `json:"describe" gorm:"column:describe"`
	Mtime       uint32 `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime       uint32 `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted     int32  `json:"deleted" gorm:"column:deleted"`
}

func (ui *CardOrderDeliveryLogInfo) TableName() string {
	return "tb_card_order_delivery_log"
}
