package main

import (
	"fmt"
	"github.com/tsuna/gohbase/hrpc"
	"hbaseDemo/database/hbase"
	"hbaseDemo/utils"
	"io"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	cnt = 0
	zeroSlice = make([]string,0)
	lock = sync.Mutex{}
	wg = sync.WaitGroup{}
	channel = make(chan *hrpc.Result, 1000)
	workerNum = 8
)

type gpsGeo struct {
	lat string
	lon string
	pid int
	townId int
	rowKey string
}

func consumerOneRecord(row *hrpc.Result) (err error) {
	var gps gpsGeo
	for _, cell := range row.Cells {
		if string(cell.Qualifier) == "pid" {
			gps.pid, _ = strconv.Atoi(string(cell.Value))
			gps.rowKey = string(cell.Row)
		}else if string(cell.Qualifier) == "town_id" {
			gps.townId,_ = strconv.Atoi(string(cell.Value))
		}
	}
	if gps.townId != 0 && gps.pid == 0 {
		lock.Lock()
		zeroSlice = append(zeroSlice, gps.rowKey)
		lock.Unlock()
	}
	return
}

func finishScan(startTime time.Time) {
	close(channel)
	wg.Wait()
	fmt.Printf("cost  is %f second\n", time.Now().Sub(startTime).Seconds())
	fmt.Println("all goroutine finished!")
	err := utils.SaveFile("./out.txt", zeroSlice)
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	//cells := hbase.GetRowKey("gps_geo", "fwvv1gww")
	//for _, cell := range cells {
	//	fmt.Println(cell.String())
	//	fmt.Println(string(cell.Value))
	//	fmt.Println(cell.CellType)
	//	fmt.Println(string(cell.Family))
	//	fmt.Println(string(cell.Qualifier))
	//	fmt.Println(string(cell.Row))
	//	fmt.Println(string(cell.Tags))
	//	fmt.Println(cell.Timestamp)
	//}

	scanner := hbase.ScanAll("gps_geo")

	// 启动 goroutine
	for i:=0;i< workerNum;i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("goroutine %d start\n", utils.GetGID())
			defer wg.Done()
			for record := range channel {
				err := consumerOneRecord(record)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
	}

	startTime := time.Now()
	for {
		row, err := scanner.Next()
		if err != nil {
			if err == io.EOF {
				finishScan(startTime)
				break
			}
			log.Fatal(err)
		}
		// 生产
		channel <- row
		cnt ++

		if cnt % 10000 ==0 {
			fmt.Printf("current cnt is %d, zero cnt is %d\n ", cnt, len(zeroSlice))
			//if len(zeroSlice)>0 {
			//	fmt.Println(zeroSlice[len(zeroSlice)-1])
			//}
		}
		if cnt == 100000 {
			finishScan(startTime)
			break
		}
	}
}
