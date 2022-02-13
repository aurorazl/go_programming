package model

import "gorm.io/datatypes"

type ActiveImei struct {
	Imei            string
	FirstActiveDate datatypes.Date `gorm:"column:first_active_date"`
}

func (ActiveImei) TableName() string {
	return "SALES_DWD.all_channel_active_info_all"
}
