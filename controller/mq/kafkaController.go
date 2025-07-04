package mq

import (
	"AdminPro/common/tool"
	"AdminPro/controller/req"
	"AdminPro/mq/kafka"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessageToKafka(ctx *gin.Context) {
	// 創建結構體，用於存放 JSON 資料
	req := req.SendMessageToKafkaReq{}
	// 使用 ShouldBindJSON 方法綁定 JSON 資料到結構體
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 如果綁定失敗，回應錯誤信息
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	if req.IsAsyncProducer == true {
		kafka.SendAsyncMessageToKafka("test-topic", "myKey", "Hello Tony Kafka 我是異步!"+req.Message)

	} else {
		kafka.SendSyncMessageToKafka("test-topic", "myKey", "Hello Tony Kafka 我是同步!"+req.Message)

	}
	ctx.JSON(http.StatusOK, tool.RespOk(tool.Success.Msg))

}
