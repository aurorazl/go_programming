package database

import (
	"KafkaConsumer/config"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

var Conn *kafka.Conn

func init() {
	// 已知缺陷：不能指定group，Reader接口可以指定，但不能batch
	conf := config.Config.KafkaConf
	conn, err := kafka.DialLeader(context.Background(), "tcp", conf.Broker, conf.Topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	Conn = conn
}

func PollBatchMsg(timeoutMs int) {
	Conn.SetReadDeadline(time.Now().Add(time.Duration(timeoutMs) * time.Millisecond))
	batch := Conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
}
