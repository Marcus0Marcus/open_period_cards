package data_service

type MerchantInfo struct {

	Id           uint64  `json:"id" gorm:"column:id"`
	Phone	string	`json:"phone" gorm:"column:phone"`
	ShopName	string	`json:"shop_name" gorm:"column:shop_name"`
	IndustryName	string	`json:"industry_name" gorm:"column:industry_name"`
	Pwd	string	`json:"pwd" gorm:"column:pwd"`
	Salt	string	`json:"salt" gorm:"column:salt"`
	Status	int32	`json:"status" gorm:"column:status"`
	Mtime   uint32  `json:"mtime" gorm:"autoUpdateTime <-:update"`
	Ctime   uint32  `json:"ctime" gorm:"autoCreateTime <-:create"`
	Deleted int32  `json:"deleted" gorm:"column:deleted"`

}

func (ui *MerchantInfo) TableName() string {
	return "tb_merchant"
}
