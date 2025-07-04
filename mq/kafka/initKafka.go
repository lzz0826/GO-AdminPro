package kafka

import "log"

func init() {
	//init異步生產者 可以多個 調用裡面的SendAsyncMessageToKafka方法
	err := InitAsyncProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("init producer error: %s", err)
	}

	//init同步步生產者 可以多個 調用裡面的SendSyncMessageToKafka方法
	err = InitSyncProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("init producer error: %s", err)
	}

	//init消費者 可以多個
	go InitConsumers([]string{"localhost:9092"}, []string{"test-topic"}, "my-consumer-group")
}

func Close() {
	if AsyncProducer != nil {
		// 關閉 producer，會自動關掉 input/output/error channel
		err := AsyncProducer.Close()
		if err != nil {
			log.Printf("error closing Kafka producer: %v", err)
		}
	}
}
