package data_service

type TBName struct {
##########
FieldName|string|`json:"field_name" gorm:"column:field_name"`
##########
}

func (ui *TBName) TableName() string {
	return "tb_name"
}
