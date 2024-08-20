package mq

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
)

func connectRabbitMQ() (*amqp.Connection, error) {
	mqDSN := fmt.Sprintf("amqp://%s:%s@%s:%s",
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.pwd"),
		viper.GetString("rabbitmq.host"),
		viper.GetString("rabbitmq.port"),
	)

	// 使用你的 RabbitMQ 詳細資訊更新連線 URL
	conn, err := amqp.Dial(mqDSN)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
