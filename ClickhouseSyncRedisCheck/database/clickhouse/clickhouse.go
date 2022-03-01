package clickhouse

import (
	"ClickhouseSyncRedisCheck/config"
	"database/sql"
	"fmt"
	Ch "github.com/ClickHouse/clickhouse-go"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
)

var Conn *sql.DB

func initByClickHouseGo() {
	dbConfig := config.Config
	driver := dbConfig.DBConf.ClickhouseConf.Driver
	ip := dbConfig.DBConf.ClickhouseConf.IP
	port := dbConfig.DBConf.ClickhouseConf.Port
	username := dbConfig.DBConf.ClickhouseConf.UserName
	password := dbConfig.DBConf.ClickhouseConf.Password

	conn, err := sql.Open(driver, fmt.Sprintf("tcp://%s:%s?username=%s&password=%s", ip, port, username, password))
	if err != nil {
		log.Fatal(err)
	}
	Conn = conn
	if err := conn.Ping(); err != nil {
		if exception, ok := err.(*Ch.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
}

var Db *gorm.DB

func initByGorm() {
	dbConfig := config.Config.DBConf.ClickhouseConf
	ip := dbConfig.IP
	database := dbConfig.Database
	port := dbConfig.Port
	username := dbConfig.UserName
	password := dbConfig.Password

	dsn := fmt.Sprintf("tcp://%s:%s?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20&parseTime=True&charset=utf8",
		ip, port, database, username, password)
	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		DSN: dsn,
		//Conn: conn,                       // initialize with existing database conn
		DisableDatetimePrecision:  true,     // disable datetime64 precision, not supported before clickhouse 20.4
		DontSupportRenameColumn:   true,     // rename column not supported before clickhouse 20.4
		SkipInitializeWithVersion: false,    // smart configure based on used version
		DefaultGranularity:        3,        // 1 granule = 8192 rows
		DefaultCompression:        "LZ4",    // default compression algorithm. LZ4 is lossless
		DefaultIndexType:          "minmax", // index stores extremes of the expression
		DefaultTableEngineOpts:    "ENGINE=MergeTree() ORDER BY tuple()",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	Db = db
}

func init() {
	initByGorm()
}

func SelectList(sql string) *sql.Rows {
	if Conn != nil {
		rows, err := Conn.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		return rows
	}
	return nil
}

func SelectMany(models interface{}) {
	//Db.Select("imei,min(first_active_date) as first_active_date").Group("imei").Find(models)
	Db.Find(models)
}

func Close() {
	Conn.Close()
}
