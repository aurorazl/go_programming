package model

import "gorm.io/datatypes"

type ActiveImei struct {
	Imei            string `gorm:"column:imei1_md5"`
	FirstActiveDate datatypes.Date `gorm:"column:active_date"`
}

func (ActiveImei) TableName() string {
	return "SALES_ODS.new_active_device_imei_all"
}
