package kafka

import (
	"AdminPro/internal/glog"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

// 新增全域同步 Producer
var SyncProducer sarama.SyncProducer

// InitSyncProducer 初始化 Kafka 同步生產者
func InitSyncProducer(addrArr []string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_1_0_0

	var err error
	SyncProducer, err = sarama.NewSyncProducer(addrArr, config)
	if err != nil {
		return err
	}

	glog.Info("Kafka SyncProducer 初始化成功")
	return nil
}

// SendSyncMessageToKafka 發送同步消息
func SendSyncMessageToKafka(topic, key, message string) {
	if SyncProducer == nil {
		glog.Errorf("Kafka SyncProducer 尚未初始化")
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := SyncProducer.SendMessage(msg)
	if err != nil {
		glog.Errorf("Kafka 同步發送失敗: %v", err)
		return
	}

	fmt.Println("Kafka 同步發送成功")
	glog.Infof("Kafka 同步發送成功 | Topic: %s | Partition: %d | Offset: %d | Key: %s | Value: %s",
		topic, partition, offset, key, message)
}

// CloseSyncKafkaProducer 關閉 Kafka 異步生產者（釋放資源）
func CloseSyncKafkaProducer() {
	if SyncProducer != nil {
		// 關閉 Producer，會關掉 input/output channel，並清理連線
		if err := SyncProducer.Close(); err != nil {
			log.Printf("Kafka SyncProducer close failed: %v", err)
		}
	}
}
