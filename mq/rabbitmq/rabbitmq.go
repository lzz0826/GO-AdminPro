package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type MQClient struct {
	conn    *amqp.Connection // RabbitMQ 的 TCP 連線，透過 Dial() 建立
	channel *amqp.Channel    // 透過 conn 建立的 AMQP 通道，用來進行發送/接收等操作
	config  MQConfig         // MQ 配置，包含交換機、隊列與綁定等設定
}

func NewMQClient(cfg MQConfig) (*MQClient, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	client := &MQClient{conn: conn, channel: ch, config: cfg}

	// 宣告所有交換機
	for _, ex := range cfg.Exchanges {
		if err := ch.ExchangeDeclare(ex.Name, ex.Kind, true, false, false, false, nil); err != nil {
			return nil, fmt.Errorf("ExchangeDeclare [%s] failed: %w", ex.Name, err)
		}
	}

	// 宣告並綁定隊列
	for _, bind := range cfg.Bindings {
		q, err := ch.QueueDeclare(bind.QueueName, true, false, false, false, nil)
		if err != nil {
			return nil, fmt.Errorf("QueueDeclare [%s] failed: %w", bind.QueueName, err)
		}
		if err := ch.QueueBind(q.Name, bind.RoutingKey, bind.ExchangeName, false, nil); err != nil {
			return nil, fmt.Errorf("QueueBind [%s -> %s] failed: %w", q.Name, bind.ExchangeName, err)
		}
	}

	return client, nil
}

// 生產者發消息
func (c *MQClient) PublishToExchange(exchange, routingKey string, body []byte) error {
	return c.channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

// PublishToQueue 直接將訊息發送到指定隊列（使用 RabbitMQ default exchange）
func (c *MQClient) PublishToQueue(queueName string, body []byte) error {
	return c.channel.Publish(
		"",        // default exchange
		queueName, // routingKey = queueName
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

// 消費者持續監聽
func (c *MQClient) StartConsumers() {
	for _, bind := range c.config.Bindings {
		// 開始消費指定隊列的訊息
		msgs, err := c.channel.Consume(
			bind.QueueName, // queue：要消費的隊列名稱
			"",             // consumer：消費者的名稱（留空表示由 RabbitMQ 自動產生）
			true,           // autoAck：是否自動 ACK（true = 收到即視為成功，不建議，應設為 false 手動 ack）
			false,          // exclusive：是否為排他性消費者（true = 其他 consumer 無法同時監聽此隊列）
			false,          // noLocal：如果設為 true，自己發送的消息自己不會收到（一般為 false，AMQP 通常不支援）
			false,          // noWait：是否不等待 server 回應就立即返回（false = 等待確認）
			nil,            // args：其他額外參數（例如 AMQP extension，可為 nil）
		)

		if err != nil {
			log.Printf("Failed to consume %s: %v", bind.QueueName, err)
			continue
		}
		// 啟動一個 goroutine 處理該隊列的訊息
		go func(msgs <-chan amqp.Delivery, handler func([]byte) error) {
			for d := range msgs {
				if err := handler(d.Body); err != nil {
					log.Printf("Handler error: %v", err)
				}
			}
		}(msgs, bind.HandlerFunc) // 傳入該隊列的訊息通道與處理邏輯
	}
}

func (c *MQClient) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
