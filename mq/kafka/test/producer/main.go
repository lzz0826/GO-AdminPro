package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"}, // Kafka broker 地址
		Topic:    "test-topic",               // 要發送的 topic
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key-A"),
			Value: []byte("你好 Kafka"),
		},
	)
	if err != nil {
		log.Fatal("寫入失敗:", err)
	}

	fmt.Println("發送成功")
}
