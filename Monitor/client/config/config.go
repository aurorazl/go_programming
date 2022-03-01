package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MonitorConfig struct {
	ConfigList []OneConfig `yaml:"configList"`
}

type OneConfig struct {
	Name      string `yaml:"name"`
	Command   string `yaml:"command"`
	BotUrl    string `yaml:"botUrl"`
	Label     string `yaml:"label"`
	Condition string `yaml:"condition"`
	Message   string `yaml:"message"`
}

func LoadConfig(fileName string) MonitorConfig {
	fmt.Println("start to load config")
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("read config file error %v", err)
	}
	var config MonitorConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}
