package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const (
	EXCHANGE_LOSE_BETSLIP_NAME = "tut.lose_betslip"
	ROUTING_KEY_LOSE_BETSLIP   = "lose_betslip-routing-key"
)

func main() {
	// RabbitMQ 连接参数
	mqURL := "amqp://twg:123456@127.0.0.1:5672"

	// 创建 RabbitMQ 连接
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// 创建 RabbitMQ 通道
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer channel.Close()

	// 声明交换机
	err = channel.ExchangeDeclare(EXCHANGE_LOSE_BETSLIP_NAME, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}

	// 声明队列
	queue, err := channel.QueueDeclare("lose_betslip", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	// 绑定队列到交换机
	err = channel.QueueBind(queue.Name, ROUTING_KEY_LOSE_BETSLIP, EXCHANGE_LOSE_BETSLIP_NAME, false, nil)
	if err != nil {
		log.Fatalf("Failed to bind queue to exchange: %s", err)
	}

	fmt.Println("Initialization complete. Ready to consume and send messages.")

	// 接收消息
	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}

	// 持续监听消息
	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	// 发送消息
	go func() {
		for {
			msgContent := "Hello World!"
			err := channel.Publish(EXCHANGE_LOSE_BETSLIP_NAME, ROUTING_KEY_LOSE_BETSLIP, false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msgContent),
			})
			if err != nil {
				log.Fatalf("Failed to publish message: %s", err)
			}
			fmt.Printf("Sent message: %s\n", msgContent)
			time.Sleep(1 * time.Second)
		}
	}()

	select {}
}
