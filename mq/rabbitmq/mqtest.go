package rabbitmq

import (
	"fmt"
	"log"
)

// Rmq01 是全域 MQClient 實例，可供其他程式呼叫發送訊息
var Rmq01 *MQClient

func init() {
	// 建立 MQ 配置，包含兩個 topic 類型的交換機與對應的隊列綁定
	cfg := MQConfig{
		URL: "amqp://twg:123456@127.0.0.1:5672", // RabbitMQ 連線字串

		// 聲明兩個交換機（輸贏分開處理）
		Exchanges: []ExchangeConfig{
			{Name: TUT_LOSE_BETSLIP_EXCHANGE, Kind: "topic"},
			{Name: TUT_WIN_BETSLIP_EXCHANGE, Kind: "topic"},
		},

		// 綁定兩個隊列到各自的交換機，並設定對應的 routing key 與處理邏輯
		Bindings: []QueueBinding{
			{
				QueueName:    LOSE_BETSLIP_QUEUE,        // 輸注單的隊列名稱
				RoutingKey:   LOSE_BETSLIP_ROUTING_KEY,  // 綁定的 routing key
				ExchangeName: TUT_LOSE_BETSLIP_EXCHANGE, // 綁定的 交換機名
				HandlerFunc: func(body []byte) error { // 訊息處理邏輯
					fmt.Println("[LOSE] =>", string(body))
					return nil
				},
			},
			{
				QueueName:    WIN_BETSLIP_QUEUE,        // 贏注單的隊列名稱
				RoutingKey:   WIN_BETSLIP_ROUTING_KEY,  // 綁定的 routing key
				ExchangeName: TUT_WIN_BETSLIP_EXCHANGE, // 綁定的 交換機名
				HandlerFunc: func(body []byte) error {
					fmt.Println("[WIN] =>", string(body))
					return nil
				},
			},
		},
	}

	// 建立 MQ 客戶端（包含連線、通道、交換機與隊列綁定）
	mqClient, err := NewMQClient(cfg)
	if err != nil {
		log.Fatal(err) // 若連線或初始化失敗，立即中止
	}

	Rmq01 = mqClient

	// 啟動所有綁定隊列的消費者（透過 goroutine 監聽並執行對應的 HandlerFunc）
	Rmq01.StartConsumers()

	// defer Close 不建議放在這裡，會立即關閉連線。
	// 正確做法是應由 main() 或程式結束時顯式呼叫 rmq01.Close()
	// defer mqClient.Close()
}
