package data_service

type CardTypeInfo struct {

	Id           uint64  `json:"id" gorm:"column:id"`
	MerchantId	uint64	`json:"merchant_id" gorm:"column:merchant_id"`
	Type	uint32	`json:"type" gorm:"column:type"`
	PeriodTimes	uint32	`json:"period_times" gorm:"column:period_times"`
	TotalTimes	uint32	`json:"total_times" gorm:"column:total_times"`
	Describe	string	`json:"describe" gorm:"column:describe"`
	Mtime   uint32  `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   uint32  `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`

}

func (ui *CardTypeInfo) TableName() string {
	return "tb_card_type"
}
