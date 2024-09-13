package utils

import (
	"AdminPro/common/tool"
	"fmt"
)

// GlobalError 继承 error 并附加状态信息
type GlobalError struct {
	Err    error
	Status *tool.Status
}

// NewGlobalError
func NewGlobalError(err error, status *tool.Status) *GlobalError {
	return &GlobalError{Err: err, Status: status}
}

// Error 实现 error 接口，返回错误信息
func (ge *GlobalError) Error() string {
	if ge.Status != nil {
		return fmt.Sprintf("error: %v, status: %d - %s", ge.Err, ge.Status.Code, ge.Status.Msg)
	}
	return ge.Err.Error()
}

// 获取状态码
func (ge *GlobalError) Code() int {
	if ge.Status != nil {
		return ge.Status.Code
	}
	return tool.SystemError.Code // 默认为系统错误码
}

// 获取状态信息
func (ge *GlobalError) Msg() string {
	if ge.Status != nil {
		return ge.Status.Msg
	}
	return tool.SystemError.Msg // 默认为系统错误信息
}
