package mq

/**
 * 輸注單的mq
 */

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	QueueLoseBetslip      = "lose_betslip"
	ExchangeLoseBetslip   = "tut.lose_betslip"
	RoutingKeyLoseBetslip = "lose_betslip-routing-key"
)

var conn *amqp.Connection
var channel *amqp.Channel
var queue amqp.Queue

func InitLoseBetslipQueue() {
	fmt.Println("註冊: " + QueueLoseBetslip + "...")
	// 创建 RabbitMQ 连接
	var err error

	conn, err = connectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	// 创建 RabbitMQ 通道
	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	// 声明交换机
	err = channel.ExchangeDeclare(ExchangeLoseBetslip, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}

	// 声明队列
	queue, err = channel.QueueDeclare(QueueLoseBetslip, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	// 绑定队列到交换机
	err = channel.QueueBind(queue.Name, RoutingKeyLoseBetslip, ExchangeLoseBetslip, false, nil)
	if err != nil {
		log.Fatalf("Failed to bind queue to exchange: %s", err)
	}

	go receiveLoseBetSlip()

	fmt.Println("Initialization complete. Ready to consume and sendLoseBetSlip messages.")
}

// SendLoseBetSlip 连接rabbitmq server
func SendLoseBetSlip(msgContent string) {
	err := channel.Publish(ExchangeLoseBetslip, RoutingKeyLoseBetslip, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})

	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}
}

func receiveLoseBetSlip() {
	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	failOnErr(err, "")

	forever := make(chan bool)

	go func() {
		//fmt.Println(*msgs)
		for d := range msgs {
			s := BytesToString(&(d.Body))
			fmt.Printf("receve msg is :%s -- \n", *s)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
