package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"}, // Kafka broker 地址
		Topic:     "test-topic",               // 要讀取的 topic
		GroupID:   "test-group",               // Consumer group ID
		Partition: 0,
		MinBytes:  1,    // 最小讀取 1B
		MaxBytes:  10e6, // 最大讀取 10MB
	})

	defer reader.Close()

	fmt.Println("等待接收 Kafka 訊息...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("讀取錯誤:", err)
			continue
		}
		fmt.Printf("收到訊息: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}
}
