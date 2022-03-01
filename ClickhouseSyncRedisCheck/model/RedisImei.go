package model

type ImeiInfo struct {
	UdidMd5 string `json:"udid_md5"`
	Imei string `json:"imei"`
	FirstActiveDate string `json:"first_active_date"`
}
