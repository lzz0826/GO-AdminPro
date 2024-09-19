package taskTest

import (
	"context"
	"net"
)

type Request struct {
	Channel           net.Conn //用户连接
	Bt                []byte   //请求的二进制数据
	RequestCode       int      //请求code
	Protocol          int      //请求协议
	UserId            int      //用户ID
	LanguageId        int      //语言ID
	ClientPlatform    int      //客户端平台
	ClientBuildNumber int      //构建编号
	CustomId          int      //自定义ID
	ProductId         int      // 产品ID
	Size              int      // 大小
	Token             string   //token
	ItemCount         int      // 业务条目数
	context.Context
}

func (r *Request) Init(conn net.Conn, buffer []byte) bool {
	r.Channel = conn
	r.Bt = buffer
	return r.parse()
}

func (r *Request) parse() bool {
	//TODO 验证用户请求

	//if r.Bt == nil || len(r.Bt) < 30 {
	//	return false
	//}
	//if r.Bt[0] != Protocal.HEADER_INDICATER_0 || r.Bt[1] != Protocal.HEADER_INDICATER_1 ||
	//	r.Bt[2] != Protocal.HEADER_INDICATER_2 || r.Bt[3] != Protocal.HEADER_INDICATER_3 {
	//	glog.Infof("data Protocal illegal")
	//	return false
	//}
	//
	//r.Protocol = int(r.Bt[Protocal.RP_PROTOCOL] & 0xff)
	//r.UserId = function.ByteArrayToInt(r.Bt, Protocal.RP_USER_ID_1)
	//r.LanguageId = int(r.Bt[Protocal.RP_LANGUAGE_ID]) & 0xff
	//r.ClientPlatform = int(r.Bt[Protocal.RP_CLIENT_PLATFORM]) & 0xff
	//r.ClientBuildNumber = int(r.Bt[Protocal.RP_CLIENT_BUILD_NUMBER]) & 0xff
	//r.CustomId = function.ByteArrayToShortInt(r.Bt, Protocal.RP_CUSTOM_ID_1)
	//r.ProductId = function.ByteArrayToShortInt(r.Bt, Protocal.RP_PRODUCT_ID_1)
	//r.RequestCode = function.ByteArrayToShortInt(r.Bt, Protocal.RP_REQUEST_CODE_HIGH)
	//r.Size = function.ByteArrayToShortInt(r.Bt, Protocal.RP_SIZE_HIGH)
	//r.Context = context.Background() // ????TraceId
	//
	//if r.RequestCode == 0 || r.Size == 0 {
	//	glog.Infof("data RequestCode illegal or data size = 0")
	//	return false
	//}
	//
	//r.Token = SetToken(r.Bt)
	//userId, err := token.TokenManagerInstance.Verify(r.Context, r.Token, r.UserId, "")
	//if userId <= 0 || err != nil {
	//	glog.ErrorfWithContext(r.Context, "token verify fail, token=%s, userId=%d, err=%v", r.Token, userId, err)
	//	return false
	//}
	return true
}

func SetToken(data []byte) string {
	//if len(data) < Protocal.RP_TOKEN_HIGH+64 {
	//	glog.Errorf("token verify fail, data:%v", data)
	//	return ""
	//}
	//
	//utf8Byte, _, err := utils.Utf162Utf8(data[Protocal.RP_TOKEN_HIGH : Protocal.RP_TOKEN_HIGH+64])
	//if err != nil {
	//	glog.Infof("parse token fail")
	//	return ""
	//}
	//return strings.TrimSpace(string(utf8Byte))

	return ""
}
