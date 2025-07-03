package mq

import (
	"AdminPro/mq/rabbitmq"
	"github.com/gin-gonic/gin"
)

func SetupRabbitmqController(ctx *gin.Context) {
	betType := ctx.Query("type")
	var req struct {
		Message string `json:"message"`
	}

	req.Message = "testetsetest"

	var routingKey, exchange string
	if betType == "win" {
		exchange = "tut.win_betslip"
		routingKey = "win_betslip-routing-key"
	} else {
		exchange = "tut.lose_betslip"
		routingKey = "lose_betslip-routing-key"
	}

	err := rabbitmq.Rmq01.PublishToExchange(exchange, routingKey, []byte(req.Message))
	if err != nil {
		ctx.JSON(500, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "sent"})
}
