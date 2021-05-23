package service

type MerchantInfo struct {
	Id           int32  `json:"id" gorm:"column:id"`
	Phone        string `json:"phone" gorm:"column:phone"`
	ShopName     string `json:"shop_name" gorm:"column:shop_name"`
	IndustryName string `json:"industry_name" gorm:"column:industry_name"`
	Pwd          string `json:"pwd" gorm:"column:pwd"`
	Status       int32  `json:"status" gorm:"column:status"`
	Salt         string `json:"salt" gorm:"column:salt"`
	Mtime        int64  `json:"mtime" gorm:"autoUpdateTime"`
	Ctime        int64  `json:"ctime" gorm:"autoCreateTime"`
	Deleted      int32  `json:"deleted" gorm:"column:deleted"`
}

func (ui *MerchantInfo) TableName() string {
	return "tb_merchant"
}
