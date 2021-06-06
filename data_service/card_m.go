package data_service

type CardInfo struct {

	Id           uint64  `json:"id" gorm:"column:id"`
	MerchantId	uint64	`json:"merchant_id" gorm:"column:merchant_id"`
	UserId	uint64	`json:"user_id" gorm:"column:user_id"`
	Name	string	`json:"name" gorm:"column:name"`
	CardTypeId	uint64	`json:"card_type_id" gorm:"column:card_type_id"`
	SerialCode	string	`json:"serial_code" gorm:"column:serial_code"`
	Used	int32	`json:"used" gorm:"column:used"`
	Describe	string	`json:"describe" gorm:"column:describe"`
	Mtime   uint32  `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   uint32  `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`

}

func (ui *CardInfo) TableName() string {
	return "tb_card"
}
