package request

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
)

type RequestTCP struct {
	Channel           net.Conn //用户连接
	Bt                []byte   //请求的二进制数据
	RequestCode       int32    //请求code
	Protocol          int32    //请求协议 自订前4字节
	UserId            int32    //用户ID
	LanguageId        int32    //语言ID
	ClientPlatform    int32    //客户端平台
	ClientBuildNumber int32    //构建编号
	CustomId          int32    //自定义ID
	ProductId         int32    // 产品ID
	Size              int32    // 大小
	Token             string   //token
	ItemCount         int32    // 业务条目数
	context.Context
}

func (r *RequestTCP) InitRequestTCP(conn net.Conn, buffer []byte) bool {
	r.Channel = conn
	r.Bt = buffer
	return r.parse()
}

func (r *RequestTCP) parse() bool {
	//if len(r.Bt) < 30 {
	//	return false
	//}
	buf := bytes.NewReader(r.Bt)

	// 假设是按照网络字节序（大端）
	binary.Read(buf, binary.BigEndian, &r.Protocol)          // 读取协议
	binary.Read(buf, binary.BigEndian, &r.UserId)            // 读取用户ID
	binary.Read(buf, binary.BigEndian, &r.LanguageId)        // 读取语言ID
	binary.Read(buf, binary.BigEndian, &r.ClientPlatform)    // 读取客户端平台
	binary.Read(buf, binary.BigEndian, &r.ClientBuildNumber) // 读取客户端构建号
	binary.Read(buf, binary.BigEndian, &r.CustomId)          // 读取自定义ID
	binary.Read(buf, binary.BigEndian, &r.ProductId)         // 读取产品ID
	binary.Read(buf, binary.BigEndian, &r.RequestCode)       // 读取请求码
	binary.Read(buf, binary.BigEndian, &r.Size)              // 读取大小

	r.Context = context.Background() // 初始化上下文

	if r.RequestCode == 0 || r.Size == 0 {
		fmt.Println("RequestCode 或 Size 非法")
		return false
	}

	// 根据具体的业务逻辑处理 Token
	r.Token = SetToken(r.Bt)
	return true
}

func SetToken(bt []byte) string {
	// 假设 Token 是从数据中某个部分读取的
	return "example-token"
}

// 将 RequestTCP 转换为 []byte
func (r *RequestTCP) ToBytes() []byte {
	buf := new(bytes.Buffer)

	// 假设是按照网络字节序（大端）写入数据
	binary.Write(buf, binary.BigEndian, int32(r.Protocol))          // 写入协议
	binary.Write(buf, binary.BigEndian, int32(r.UserId))            // 写入用户ID
	binary.Write(buf, binary.BigEndian, int32(r.LanguageId))        // 写入语言ID
	binary.Write(buf, binary.BigEndian, int32(r.ClientPlatform))    // 写入客户端平台
	binary.Write(buf, binary.BigEndian, int32(r.ClientBuildNumber)) // 写入客户端构建号
	binary.Write(buf, binary.BigEndian, int32(r.CustomId))          // 写入自定义ID
	binary.Write(buf, binary.BigEndian, int32(r.ProductId))         // 写入产品ID
	binary.Write(buf, binary.BigEndian, int32(r.RequestCode))       // 写入请求码
	binary.Write(buf, binary.BigEndian, int32(r.Size))              // 写入大小

	//binary.Write 不能直接处理 string

	// 返回生成的字节数组
	return buf.Bytes()
}

//使用占位
//func (r *RequestTCP) parse() bool {
//	if r.Bt == nil || len(r.Bt) < 30 {
//		return false
//	}
//	if r.Bt[0] != task.HEADER_INDICATER_0 || r.Bt[1] != task.HEADER_INDICATER_1 ||
//		r.Bt[2] != task.HEADER_INDICATER_2 || r.Bt[3] != task.HEADER_INDICATER_3 {
//
//		fmt.Printf("data Protocal illegal")
//		glog.Infof("data Protocal illegal")
//		return false
//	}
//
//	//从 bytes中依照索引取得 TODO
//	r.Protocol = int(r.Bt[task.RP_PROTOCOL] & 0xff) //0~3字节
//	r.UserId = utils.ByteArrayToInt(r.Bt, task.RP_USER_ID_1)
//	r.LanguageId = int(r.Bt[task.RP_LANGUAGE_ID]) & 0xff
//	r.ClientPlatform = int(r.Bt[task.RP_CLIENT_PLATFORM]) & 0xff
//	r.ClientBuildNumber = int(r.Bt[task.RP_CLIENT_BUILD_NUMBER]) & 0xff
//	r.CustomId = utils.ByteArrayToShortInt(r.Bt, task.RP_CUSTOM_ID_1)
//	r.ProductId = utils.ByteArrayToShortInt(r.Bt, task.RP_PRODUCT_ID_1)
//	r.RequestCode = utils.ByteArrayToShortInt(r.Bt, task.RP_REQUEST_CODE_HIGH)
//	r.Size = utils.ByteArrayToShortInt(r.Bt, task.RP_SIZE_HIGH)
//	r.Context = context.Background() // 用来传递TraceId
//
//	if r.RequestCode == 0 || r.Size == 0 {
//		glog.Infof("data RequestCode illegal or data size = 0")
//		return false
//	}
//
//	r.Token = SetToken(r.Bt)
//	//处理验证TOKEN
//	//userId, err := token.TokenManagerInstance.Verify(r.Context, r.Token, r.UserId, "")
//	//if userId <= 0 || err != nil {
//	//	glog.ErrorfWithContext(r.Context, "token verify fail, token=%s, userId=%d, err=%v", r.Token, userId, err)
//	//	return false
//	//}
//	return true
//}
//
//func SetToken(data []byte) string {
//	if len(data) < task.RP_TOKEN_HIGH+64 {
//		glog.Errorf("token verify fail, data:%v", data)
//		return ""
//	}
//
//	utf8Byte, _, err := utils.Utf162Utf8(data[task.RP_TOKEN_HIGH : task.RP_TOKEN_HIGH+64])
//	if err != nil {
//		glog.Infof("parse token fail")
//		return ""
//	}
//	return strings.TrimSpace(string(utf8Byte))
//}
