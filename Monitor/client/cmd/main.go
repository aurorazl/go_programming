package main

/**
1。读取配置文件
2。执行command
3。判断条件，是否告警
*/

import (
	"Monitor/client/config"
	"Monitor/client/utils"
	"flag"
	"fmt"
	"github.com/apaxa-go/eval"
	"log"
	"strings"
)

func judgeCondition(condition string) bool {
	expr, err := eval.ParseString(condition, "")
	if err != nil {
		log.Fatalf("parse condition failed, %v", err)
	}
	r, err := expr.EvalToInterface(nil)
	if err != nil {
		log.Fatalf("eval failed, %v", err)
	}
	v, ok := r.(bool)
	if ok {
		return v
	}
	return false
}

func sendAlert(config config.OneConfig) {
	fmt.Printf("%s条件不满足，开始发送告警信息\n", config.Name)
	headers := map[string]string{"Content-Type": "application/json"}
	rawBody := fmt.Sprintf("{\"msg_type\":\"text\",\"content\":{\"text\":\"[%s] %s 的数据还没生成\"}}", config.Label, config.Name)
	_, err := utils.DoPost(config.BotUrl, headers, rawBody)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var fileName = flag.String("config", "config.yaml", "配置文件路径")
	flag.Parse()
	monitorConfig := config.LoadConfig(*fileName)
	for _, config := range monitorConfig.ConfigList {
		output, err := utils.ExecCommand(config.Command)
		if err != nil {
			log.Fatalf("exec command fail,reason is %v", err)
		}
		if !judgeCondition(fmt.Sprintf(config.Condition, strings.Trim(output, "\n"))) {
			sendAlert(config)
		}
	}
}
