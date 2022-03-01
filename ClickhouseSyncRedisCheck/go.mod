module ClickhouseSyncRedisCheck

go 1.16

require (
	github.com/ClickHouse/clickhouse-go v1.5.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/hashicorp/go-version v1.4.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/datatypes v1.0.5
	gorm.io/driver/clickhouse v0.2.2
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.3.1 // indirect
	gorm.io/driver/sqlite v1.3.1 // indirect
	gorm.io/driver/sqlserver v1.3.1 // indirect
	gorm.io/gorm v1.23.1
)
