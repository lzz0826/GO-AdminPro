package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 定義 ConsumerGroupHandler 結構，實作三個方法
type ConsumerGroupHandler struct{}

// Setup 在 Consumer group 啟動時執行（可用來初始化資源）
func (h *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer Group Setup")
	return nil
}

// Cleanup 在 Consumer group 結束前執行（可用來釋放資源）
func (h *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer Group Cleanup")
	return nil
}

// ConsumeClaim 處理每個 partition 的訊息
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("收到訊息 | topic: %s | partition: %d | offset: %d\n", msg.Topic, msg.Partition, msg.Offset)
		fmt.Printf("key: %s\n", string(msg.Key))
		fmt.Printf("value: %s\n", string(msg.Value))

		// 手動標記已處理（commit offset）
		session.MarkMessage(msg, "")

		//// 處理成功才 ACK（手動提交 offset）
		//err := process(msg)
		//if err == nil {
		//	session.MarkMessage(msg, "")
		//} else {
		//	log.Printf("處理失敗，不提交 offset: %v", err)
		//}

	}
	return nil
}

func InitConsumers(brokers, topics []string, groupID string) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("建立 ConsumerGroup 失敗: %v", err)
	}
	defer client.Close()

	// 用 context.Background() 正確取得可取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 監聽 SIGINT/SIGTERM 訊號並取消 context
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		cancel()
	}()

	handler := &ConsumerGroupHandler{}
	fmt.Println("Kafka 消費者啟動中...")

	for {
		if err := client.Consume(ctx, topics, handler); err != nil {
			log.Printf("Consume 錯誤: %v", err)
		}
		if ctx.Err() != nil {
			break
		}
	}
	fmt.Println("Kafka 消費者已結束")
}
