package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"log"
)

var channel *amqp.Channel // 透過 conn 建立的 AMQP 通道，用來進行發送/接收等操作

func main() {
	fmt.Println("主線程啟動...")
	initMQClient()
	startConsumers([]string{"lose_betslip_queue", "win_betslip_queue"})

	// 阻塞主程式（永遠不會結束）
	select {}

}

func initMQClient() {
	conn, err := amqp.Dial("amqp://twg:123456@127.0.0.1:5672")
	if err != nil {
		log.Fatal(err) // 若連線或初始化失敗，立即中止

	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err) // 若連線或初始化失敗，立即中止
	}

	// 建立 MQ
	channel = ch
}

func startConsumers(queueName []string) {
	fmt.Println("開啟監聽對列.....", queueName)
	//監聽隊列
	for _, bind := range queueName {
		// 開始消費指定隊列的訊息
		msgs, err := channel.Consume(bind, "", true, false, false, false, nil)
		if err != nil {
			log.Printf("Failed to consume %s: %v", bind, err)
			continue
		}

		//啟動一個 goroutine 處理該隊列的訊息
		//go func(msgs <-chan amqp.Delivery, handler func([]byte) error) {
		//	for d := range msgs {
		//		if err := handler(d.Body); err != nil {
		//			log.Printf("Handler error: %v", err)
		//		}
		//	}
		//}(msgs, MyHandler) // 傳入該隊列的訊息通道與處理邏輯

		//上面等同 handler func([]byte) error 帶入制定函數
		go consumeMessages(msgs, MyHandler)

	}
}

// func([]byte) error匿名函數 可以帶入自訂的 func
func consumeMessages(msgs <-chan amqp.Delivery, handler func([]byte) error) {
	for d := range msgs {
		if err := handler(d.Body); err != nil {
			log.Printf("Handler error: %v", err)
		}

		//手動ACK
		//err := handler(d.Body)
		//if err != nil {
		//	log.Printf("Handler error: %v", err)
		//	// 可選：可以做 d.Nack(false, true) 來重入隊列
		//	continue
		//}
		//// 處理成功後，手動 ack
		//if err := d.Ack(false); err != nil {
		//	log.Printf("Failed to ACK message: %v", err)
		//}
	}
}

type MqMessage struct {
	Type string      `json:"type"` //用於區分業務邏輯的 type
	Data interface{} `json:"data"`
}

type User struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

func MyHandler(msg []byte) error {
	fmt.Println("收到消息:", string(msg))

	var m MqMessage
	err := json.Unmarshal(msg, &m)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
		return err
	}

	fmt.Println("Type:", m.Type)

	// 把 m.Data 先轉成 JSON 字串
	dataBytes, err := json.Marshal(m.Data)
	if err != nil {
		log.Printf("Marshal inner Data error: %v", err)
		return err
	}

	var u User
	err = json.Unmarshal(dataBytes, &u)
	if err != nil {
		log.Printf("Unmarshal to User error: %v", err)
		return err
	}

	fmt.Printf("解析出 User: ID=%d, Name=%s\n", u.UserId, u.UserName)

	return nil
}
