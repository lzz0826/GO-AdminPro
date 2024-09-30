package response

import (
	"bytes"
	"encoding/binary"
)

type ResponseTCP struct {
	Code     int32
	Bt       []byte //回复的二进制数据
	Size     int32
	CustomId int32 //自定义ID
}

func (r *ResponseTCP) Parse() bool {
	buf := bytes.NewReader(r.Bt)
	if err := binary.Read(buf, binary.BigEndian, &r.Code); err != nil {
		return false
	}
	//自定义ID
	if err := binary.Read(buf, binary.BigEndian, &r.CustomId); err != nil {
		return false
	}
	return true
}

// 将 ResponseTCP 转换为 []byte
func (r *ResponseTCP) ToBytes() []byte {
	buf := new(bytes.Buffer)
	// 假设是按照网络字节序（大端）写入数据
	binary.Write(buf, binary.BigEndian, int32(r.Code))
	binary.Write(buf, binary.BigEndian, int32(r.CustomId)) // 写入自定义ID

	// 返回生成的字节数组
	return buf.Bytes()
}
