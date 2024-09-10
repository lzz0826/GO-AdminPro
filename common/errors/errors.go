package errors

import (
	"AdminPro/common/enum"
	"fmt"
)

// 錯誤類型
const (
	cateSys = "sys"
	cateIO  = "io"
	cateBiz = "biz"
)

type Errx struct {
	code int    // 錯誤碼
	msg  string // 錯誤訊息
	cate string // 錯誤類型
}

func (e Errx) Msg() string {
	return e.msg
}

func (e Errx) Cate() string {
	return e.cate
}

func (e Errx) Code() int {
	return e.code
}

// 複寫原本的 error.Error() 裡的反回的 string
func (e Errx) Error() string {
	return fmt.Sprintf("failed to %s,code:%d,cate:%s", e.msg, e.code, e.cate)
}

func (e Errx) SetMsg(msg string) Errx {
	e.msg = msg
	return e
}

func newErrx(code int, msg string, cate string) Errx {
	err := Errx{
		code: code,
		msg:  msg,
		cate: cate,
	}
	return err
}

// 系统错误
func NewSysErrx(code int, msg string) Errx {
	return newErrx(code, msg, cateSys)
}

// 业务逻辑错误
func NewBizErrx(code int, msg string) Errx {
	return newErrx(code, msg, cateBiz)
}

// IO错误
func NewIOErrx(code int, msg string) Errx {

	return newErrx(code, msg, cateIO)
}

var (
	ErrSystem     = NewSysErrx(int(enum.ERROR), enum.GetResponseMsg(enum.ERROR))
	SelectFail    = NewBizErrx(int(enum.SelectFail), enum.GetResponseMsg(enum.SelectFail))
	ErrIOReadFail = NewIOErrx(int(enum.IO_ERROR), enum.GetResponseMsg(enum.IO_ERROR))
)
