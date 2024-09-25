package request

import (
	"context"
)

type HttpTaskRequest struct {
	RequestCode      int                    `json:"request_code"`      // 请求code 对印任务key
	RequestParameter map[string]interface{} `json:"request_parameter"` // 请求的参数
	Protocol         int                    `json:"protocol"`          // 请求协议
	UserId           int                    `json:"user_id"`           // 用户ID
	Token            string                 `json:"token"`             // token
	ItemCount        int                    `json:"item_count"`        // 业务条目数
	context.Context                         // 继承 context.Context 来处理请求生命周期
}
