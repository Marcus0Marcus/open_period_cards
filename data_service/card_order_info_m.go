package data_service

type CardOrderInfoInfo struct {

	Id           uint64  `json:"id" gorm:"column:id"`
	CardOrderId	uint64	`json:"card_order_id" gorm:"column:card_order_id"`
	Content	string	`json:"content" gorm:"column:content"`
	Describe	string	`json:"describe" gorm:"column:describe"`
	Mtime   uint32  `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   uint32  `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`

}

func (ui *CardOrderInfoInfo) TableName() string {
	return "tb_card_order_info"
}
