package model

type LegalImeiRecord struct {
	Imei string
}

func (LegalImeiRecord) TableName() string {
	return "SALES_DWD.legal_imei_record_all"
}
