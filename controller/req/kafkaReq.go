package req

type SendMessageToKafkaReq struct {
	IsAsyncProducer bool   `json:"isAsyncProducer"`
	Message         string `json:"message"`
}
