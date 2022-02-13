package database

import (
	"KafkaConsumer/config"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)

var consumer *kafka.Consumer

func init() {
	conf := config.Config.KafkaConf
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  conf.Broker,
		"group.id":           "myGroup",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": conf.AutoCommit,
	})
	if err != nil {
		panic(err)
	}
	c.Subscribe(conf.Topic, nil)
	consumer = c
}

/*
for macos:
	brew install librdkafka pkg-config
	goland配置go build参数-tags dynamic，环境变量：PKG_CONFIG_PATH=/opt/homebrew/opt/openssl@1.1/lib/pkgconfig

*/
func ConsumeOneMessage() *kafka.Message {
	// 阻塞等待，直到来了一条消息
	msg, err := consumer.ReadMessage(-1)
	if err != nil {
		// The client will automatically try to recover from all errors.
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}
	return msg
}

func ConsumeBatchMessage(batchSize int, timeoutMs int) []*kafka.Message {
	currentCnt := 0
	start := time.Now()
	ret := []*kafka.Message{}
	for currentCnt < batchSize && time.Since(start).Milliseconds() <= int64(timeoutMs) {
		ev := consumer.Poll(1000)
		if ev == nil {
			continue
		}
		switch e := ev.(type) {
		case *kafka.Message:
			currentCnt++
			ret = append(ret, e)
		case kafka.Error:
			// Errors should generally be considered
			// informational, the client will try to
			// automatically recover.
			// But in this example we choose to terminate
			// the application if all brokers are down.
			fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
	return ret
}

func DoCommit() ([]kafka.TopicPartition, error) {
	commit, err := consumer.Commit()
	return commit, err
}
