package mq

import (
	"AdminPro/mq/rabbitmq"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SetupRabbitmqController(ctx *gin.Context) {
	betType := ctx.Query("betType")

	var routingKey, exchange string
	if betType == "win" {
		exchange = "tut.win_betslip"
		routingKey = "win_betslip-routing-key"
	} else {
		exchange = "tut.lose_betslip"
		routingKey = "lose_betslip-routing-key"
	}

	var User struct {
		UserId   int    `json:"userId"`
		UserName string `json:"userName"`
	}
	User.UserId = 1
	User.UserName = "tony"

	message := rabbitmq.MqMessage{
		//Type: "轉給客服",
		Type: "轉給站內信",
		Data: &User,
	}

	//轉成byte
	marshal, err := json.Marshal(message)
	if err != nil {
		ctx.JSON(400, gin.H{"err": err.Error()})

	}

	err = rabbitmq.Rmq01.PublishToExchange(exchange, routingKey, marshal)
	if err != nil {
		ctx.JSON(500, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "sent"})
}
