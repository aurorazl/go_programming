package main

import (
	"ClickhouseSyncRedisCheck/database/clickhouse"
	"ClickhouseSyncRedisCheck/database/redis"
	"ClickhouseSyncRedisCheck/model"
	"ClickhouseSyncRedisCheck/utils"
	"fmt"
	"log"
	"strings"
)

/*
1.读取clickhouse数据
2。读取redis数据。
3。对比两者差异
*/

func queryClickhouseOneStringField(sql string) []string {
	rows := clickhouse.SelectList(sql)
	fieldList := []string{}
	defer rows.Close()
	for rows.Next() {
		var (
			field string
		)
		if err := rows.Scan(&field); err != nil {
			log.Fatal(err)
		}
		fieldList = append(fieldList, field)
	}
	fmt.Println(fmt.Sprintf("imei list in clickhouse length is %d", len(fieldList)))
	return fieldList
}

func queryClickhouseActiveImei() []string {
	var clickhouseImeis []model.ActiveImei
	clickhouse.SelectMany(&clickhouseImeis)
	imeis := []string{}
	for _, one := range clickhouseImeis {
		imeis = append(imeis, one.Imei)
	}
	return imeis
}

func ScanRedisActiveImeis() []string {
	prefix := "zg_1_"
	rets := redis.Scan(prefix + "*")
	imei_list := []string{}
	for _, imeiStr := range rets {
		split := strings.Split(imeiStr, prefix)
		imei_list = append(imei_list, split[1])
	}
	return imei_list
}

func SScanRedisImei(key string) []string {
	imeis := redis.SScan(key)
	return imeis
}

func CheckActiveImeiDiff() {
	//sql := "select distinct imei from SALES_DWD.all_channel_active_info_all"
	//clickhouseImeis := queryClickhouseOneStringField(sql)
	clickhouseImeis := queryClickhouseActiveImei()
	redisImeis := ScanRedisActiveImeis()
	onlyInSrc, onlyInDst := utils.GetTwoSliceDiff(clickhouseImeis, redisImeis)
	if onlyInSrc != nil && onlyInDst != nil {
		fmt.Println("only in redis imei list is ", len(onlyInDst))
		if len(onlyInDst) > 0 {
			fmt.Println(onlyInDst[:100])
		}
		fmt.Println("only in clickhouse imei list length is ", len(onlyInSrc))
		if len(onlyInSrc) > 0 {
			fmt.Println(onlyInSrc[:100])
		}
	}
}

func CheckLegalImeiDiff() {
	//sql := "select distinct imei from SALES_DWD.legal_imei_record_all"
	//clickhouseImeis := queryClickhouseOneStringField(sql)
	var clickhouseImeis []model.LegalImeiRecord
	clickhouse.SelectMany(&clickhouseImeis)
	redisImeis := SScanRedisImei("legal_imei_set")
	onlyInSrc, onlyInDst := utils.GetTwoSliceDiff(clickhouseImeis, redisImeis)
	if onlyInSrc != nil && onlyInDst != nil {
		fmt.Println("only in redis imei list is ", len(onlyInDst), onlyInDst)
		fmt.Println("only in clickhouse imei list length is ", len(onlyInSrc), onlyInSrc)
	}
}

type test struct {
	Name string
	Age  int
}

var a test

func (test) TableName() string {
	return "test.test"
}
func main() {
	CheckActiveImeiDiff()
	//CheckLegalImeiDiff()
	//rows, err := mysql.MysqlDb.Query("select * from test.test")
	//if err !=nil {
	//	log.Fatal(err)
	//}
	//for rows.Next() {
	//	rows.Scan(&a.Name,&a.Age)
	//}
	//fmt.Println(a)
	//time.Sleep(1000 * time.Second)
}
