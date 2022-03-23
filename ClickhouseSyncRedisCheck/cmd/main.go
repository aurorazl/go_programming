package main

import (
	"ClickhouseSyncRedisCheck/database/clickhouse"
	"ClickhouseSyncRedisCheck/database/redis"
	"ClickhouseSyncRedisCheck/model"
	"ClickhouseSyncRedisCheck/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
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
	prefix := "zg_3_"
	rets := redis.Scan(prefix + "*")
	fmt.Println("redis scan done")
	imei_list := []string{}
	for _, imeiStr := range rets {
		split := strings.Split(imeiStr, prefix)
		imei_list = append(imei_list, split[1])
	}
	return imei_list
}

func DeleteRedisImeis(imeis []interface{}) {
	prefix := "zg_1_"
	for _, imei := range imeis {
		v,ok := imei.(string)
		if ! ok {
			log.Fatalf("imei: %s convert fail", imei)
		}
		firstActiveDate := GetImeiFirstActiveDate(v)
		if time.Now().Format("2006-01-02") == firstActiveDate {
			fmt.Printf("delete imei: %s \n", v)
			redis.DeleteKey(prefix + v)
		}
	}
}

func GetImeiFirstActiveDate(imei string) string {
	prefix := "zg_3_"
	data := redis.HGetAll(prefix + imei)
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	var imeiInfo model.ImeiInfo
	if err:= json.Unmarshal(jsonStr, &imeiInfo);err!=nil {
		fmt.Println(err)
	}
	return imeiInfo.FirstActiveDate
}

func SScanRedisImei(key string) []string {
	imeis := redis.SScan(key)
	return imeis
}

func CheckActiveImeiDiff() {
	//sql := "select distinct imei from SALES_DWD.all_channel_active_info_all"
	//clickhouseImeis := queryClickhouseOneStringField(sql)
	// 先查redis，再查imei，同步会有延迟。
	redisImeis := ScanRedisActiveImeis()
	clickhouseImeis := queryClickhouseActiveImei()
	fmt.Printf("redis count is %d, clickhouse count is %d \n", len(redisImeis), len(clickhouseImeis))
	onlyInSrc, onlyInDst := utils.GetTwoSliceDiff(clickhouseImeis, redisImeis)
	if onlyInSrc != nil && onlyInDst != nil {
		fmt.Println("only in redis imei list is ", len(onlyInDst))
		if len(onlyInDst) > 0 {
			fmt.Println(onlyInDst[:len(onlyInDst)%10])
		}
		fmt.Println("only in clickhouse imei list length is ", len(onlyInSrc))
		if len(onlyInSrc) > 0 {
			fmt.Println(onlyInSrc[:len(onlyInSrc)%10])
		}
	}
	// bug： 两个查询之间有间隔延迟imei，需要排除掉。
	//DeleteRedisImeis(onlyInDst)
}

func CheckLegalImeiDiff() {
	//sql := "select distinct imei from SALES_DWD.legal_imei_record_all"
	//clickhouseImeis := queryClickhouseOneStringField(sql)
	var clickhouseImeis []model.LegalImeiRecord
	clickhouse.SelectMany(&clickhouseImeis)
	clickhouseImeiList := []string{}
	for _, model := range clickhouseImeis {
		clickhouseImeiList = append(clickhouseImeiList, model.Imei)
	}
	redisImeis := SScanRedisImei("legal_imei_set")
	onlyInSrc, onlyInDst := utils.GetTwoSliceDiff(clickhouseImeiList, redisImeis)
	if onlyInSrc != nil && onlyInDst != nil {
		fmt.Println("only in redis imei list is ", len(onlyInDst))
		if len(onlyInDst) > 0 {
			fmt.Println(onlyInDst[:10])
		}
		fmt.Println("only in clickhouse imei list length is ", len(onlyInSrc))
		if len(onlyInSrc) > 0 {
			fmt.Println(onlyInSrc[:10])
		}
	}
}

func main() {
	CheckActiveImeiDiff()
	//CheckLegalImeiDiff()
}

