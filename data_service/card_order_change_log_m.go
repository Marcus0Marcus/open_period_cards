package data_service

type CardOrderChangeLogInfo struct {

	Id           uint64  `json:"id" gorm:"column:id"`
	CardOrderId	uint64	`json:"card_order_id" gorm:"column:card_order_id"`
	ChangeLog	string	`json:"change_log" gorm:"column:change_log"`
	Describe	string	`json:"describe" gorm:"column:describe"`
	Mtime   uint32  `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   uint32  `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`

}

func (ui *CardOrderChangeLogInfo) TableName() string {
	return "tb_card_order_change_log"
}
