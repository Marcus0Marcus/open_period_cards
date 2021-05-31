package data_service

type CardOrderInfo struct {
	Id               int64  `json:"id" gorm:"column:id"`
	MerchantId       int32  `json:"merchant_id" gorm:"column:merchant_id"`
	UserId           int32  `json:"user_id" gorm:"column:user_id"`
	CardId           int32  `json:"card_id" gorm:"column:card_id"`
	CardTypeId       int32  `json:"card_type_id" gorm:"column:card_type_id"`
	SendType         int32  `json:"send_type" gorm:"column:send_type"`
	SendDayList      string `json:"send_day_list" gorm:"column:send_day_list"`
	PeriodSendTimes  uint32 `json:"period_send_times" gorm:"column:period_send_times"`
	TotalSendTimes   uint32 `json:"total_send_times" gorm:"column:total_send_times"`
	IsTotalFinished  int32  `json:"is_total_finished" gorm:"column:is_total_finished"`
	IsPeriodFinished int32  `json:"is_period_finished" gorm:"column:is_period_finished"`
	LastSendTime     uint32 `json:"last_send_time" gorm:"column:last_send_time"`
	Describe         string `json:"describe" gorm:"column:describe"`
	Mtime            uint32 `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime            uint32 `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted          int32  `json:"deleted" gorm:"column:deleted"`
}

func (ui *CardOrderInfo) TableName() string {
	return "tb_card_order"
}
