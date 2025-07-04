package kafka

import (
	"AdminPro/internal/glog"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

// 全域變數 AsyncProducer，用來存放 Kafka 的異步生產者實例
var AsyncProducer sarama.AsyncProducer

// 可以在InitProducer2多個
var AsyncProducer2 sarama.AsyncProducer

// InitAsyncProducer 初始化 Kafka 異步生產者
func InitAsyncProducer(addrArr []string) error {
	// 建立 Sarama 的 Kafka 配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待所有副本都確認收到後才視為發送成功（最高可靠性）
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 隨機分配分區（partition）
	config.Producer.Return.Successes = true                   // 啟用成功訊息回傳
	config.Producer.Return.Errors = true                      // 啟用失敗訊息回傳
	config.Version = sarama.V2_1_0_0                          // 指定 Kafka 協議版本（要與 Kafka server 相容）

	// Kafka broker 的地址陣列，可設定多個以容錯
	//addrArr := strings.Split("localhost:9092,localhost:9093", ",")
	//addrArr := strings.Split(os.Getenv("KAFKA_ADDRS"), ",")

	var err error
	// 建立異步生產者（AsyncProducer）
	AsyncProducer, err = sarama.NewAsyncProducer(addrArr, config)
	if err != nil {
		// 建立失敗就回傳錯誤
		return err
	}

	glog.Info("Kafka producer 初始化成功")

	// 建立 goroutine 處理訊息發送成功或失敗的回報
	go func() {
		for {
			select {
			//監聽成功發送後的消息
			case msg := <-AsyncProducer.Successes():
				// 讀出 Topic、Partition、Offset 和 Message Key/Value
				topic := msg.Topic
				partition := msg.Partition
				offset := msg.Offset

				var keyStr, valStr string
				if msg.Key != nil {
					keyBytes, _ := msg.Key.Encode()
					keyStr = string(keyBytes)
				}
				if msg.Value != nil {
					valBytes, _ := msg.Value.Encode()
					valStr = string(valBytes)
				}

				fmt.Println("Kafka 異步發送成功!!")
				glog.Infof("Kafka 發送成功 | Topic: %s | Partition: %d | Offset: %d | Key: %s | Value: %s",
					topic, partition, offset, keyStr, valStr)
			case fail := <-AsyncProducer.Errors():
				glog.Errorf("Kafka 發送失敗：%v", fail.Err)
			}
		}
	}()

	return nil
}

// SendAsyncMessageToKafka 發送一筆訊息到指定的 topic
// @param topic  Kafka topic 名稱
// @param key    分區鍵（可選，用來決定訊息落在哪個 partition）
// @param message 訊息內容
func SendAsyncMessageToKafka(topic, key, message string) {
	// 如果尚未初始化 AsyncProducer，則跳過發送
	if AsyncProducer == nil {
		glog.Errorf("Kafka AsyncProducer 尚未初始化")
		return
	}

	// 將訊息送入 AsyncProducer 的 Input channel 異不需要另外監聽AsyncProducer.Successes()
	AsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: topic,                         // 目標 topic
		Key:   sarama.StringEncoder(key),     // 分區鍵（可為空字串）
		Value: sarama.StringEncoder(message), // 訊息內容（字串轉換為 sarama Encoder）
	}
}

// CloseAsyncKafkaProducer 關閉 Kafka 異步生產者（釋放資源）
func CloseAsyncKafkaProducer() {
	if AsyncProducer != nil {
		// 關閉 Producer，會關掉 input/output channel，並清理連線
		if err := AsyncProducer.Close(); err != nil {
			log.Printf("Kafka  AsyncProducer close failed: %v", err)
		}
	}
}
