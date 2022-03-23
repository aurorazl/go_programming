package hbase

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"hbaseDemo/config"
	"log"
)


var (
	client gohbase.Client

)
func init() {
	client = gohbase.NewClient(config.Config.DBConf.HbaseConf.ZkQuorum)
}

func GetRowKey(table, key string) (ret []*hrpc.Cell) {
	getRequest, err := hrpc.NewGetStr(context.Background(), table, key)
	if err != nil {
		log.Fatal(err)
	}
	getRsp, err := client.Get(getRequest)
	if getRsp != nil {
		ret = getRsp.Cells
	}
	return
}

func ScanPrefix(prefix string) {
	pFilter := filter.NewPrefixFilter([]byte(prefix))
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "table", hrpc.Filters(pFilter))
	scanRsp := client.Scan(scanRequest)
	fmt.Println(scanRsp)
}

func ScanAll(table string) (ret hrpc.Scanner) {
	// 只能设置fetch多少，下次需要重新offset然后fetch
	scanRequest, _ := hrpc.NewScanStr(context.Background(), table, hrpc.NumberOfRows(20000))
	ret = client.Scan(scanRequest)
	return
}