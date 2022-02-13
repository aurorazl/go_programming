package main

import (
	"KafkaConsumer/config"
	"KafkaConsumer/database"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type bill struct {
	BillType string `json:"bill_type"`
	Date     string `json:date`
}

func execOnOneMessageArrival() {
	msg := database.ConsumeOneMessage()
	var oneMessage bill
	json.Unmarshal(msg.Value, &oneMessage)
	fmt.Println(oneMessage.BillType, oneMessage.Date)
	command := fmt.Sprintf("%s %s", getTask(oneMessage.BillType).Command, oneMessage.Date)
	execCommand(command)
}

func execOnBatchMessageArrival() {
	conf := config.Config.KafkaConf
	msgs := database.ConsumeBatchMessage(conf.BatchSize, conf.BatchIntervalInSecond*1000)
	if len(msgs) == 0 {
		return
	}
	fmt.Printf("%s: current message batch size is %d\n", time.Now().String(), len(msgs))
	taskMap := make(map[string]string)
	var oneMessage bill
	for _, msg := range msgs {
		json.Unmarshal(msg.Value, &oneMessage)
		task := getTask(oneMessage.BillType)
		if task != nil {
			if taskMap[task.BillType] == "" {
				taskMap[task.BillType] = fmt.Sprintf("%s %s", task.Command, oneMessage.Date)
			} else {
				taskMap[task.BillType] = fmt.Sprintf("%s %s", taskMap[task.BillType], oneMessage.Date)
			}
		}
	}
	for _, command := range taskMap {
		split := strings.Split(command, " ")
		if len(split) > 1 {
			execCommand(split[0], split[1:]...)
		} else {
			execCommand(split[0])
		}
	}
	database.DoCommit()
}

func getTask(billType string) *config.UpdateTaskConfig {
	tasks := config.Config.TasksConf
	for _, task := range tasks {
		if strings.HasPrefix(billType, task.BillType) {
			return &task
		}
	}
	return nil
}

func execCommand(command string, args ...string) {
	fmt.Printf("exec command: %s %s\n", command, args)
	cmd := exec.Command(command, args...)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("exec command error: %s\n", err)
	}
}

func main() {
	stop := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Printf("receiving signal: %s\n", sig)
		stop = true
	}()

	for !stop {
		execOnBatchMessageArrival()
	}
}
