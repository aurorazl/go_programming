package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type RootConfig struct {
	DBConf DBConfig `yaml:"db"`
}

type DBConfig struct {
	ClickhouseConf ClickhouseConfig `yaml:"clickhouse"`
	RedisConf      RedisConfig      `yaml:"redis"`
	MysqlConf      MysqlConfig      `yaml:"mysql"`
}

type ClickhouseConfig struct {
	Driver   string `yaml:"driver"`
	IP       string `yaml:"ip"`
	Port     string `yarml:port`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:database`
}

type RedisConfig struct {
	IP       string `yaml:"ip"`
	Port     string `yarml:port`
	Database int    `yaml:"database"`
	Password string `yaml:"password"`
}

type MysqlConfig struct {
	IP       string `yaml:"ip"`
	Port     string `yarml:port`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

var Config RootConfig

func LoadConfig(fileName string) RootConfig {
	fmt.Println("start to load config")
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("read config file error %v", err)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println("load config done")
	return Config
}

func init() {
	LoadConfig("./config.yaml")
}
