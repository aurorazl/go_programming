package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type RootConfig struct {
	KafkaConf kafkaConfig        `yaml:"kafka"`
	TasksConf []UpdateTaskConfig `yaml:"tasks"`
}

type kafkaConfig struct {
	Broker                string `yaml:"broker"`
	Topic                 string `yaml:"topic"`
	AutoCommit            string `yaml:"auto_commit"`
	BatchSize             int    `yaml:"batch_size"`
	BatchIntervalInSecond int    `yaml:"batch_interval_in_second"`
}

type UpdateTaskConfig struct {
	BillType string `yaml:"bill_type"`
	Command  string `yaml:"command"`
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
