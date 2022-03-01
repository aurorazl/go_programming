package mysql

import (
	"GinQPS/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

func init() {
	conf := config.Config.DBConf.MysqlConf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName, conf.Password, conf.IP, conf.Port)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db = db
}

var MysqlDb *sql.DB

func initBySQL() {
	conf := config.Config.DBConf.MysqlConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		conf.UserName, conf.Password, conf.IP, conf.Port))
	if err != nil {
		panic(err)
	}
	MysqlDb = db
	db.SetConnMaxLifetime(time.Second * 10)
}
