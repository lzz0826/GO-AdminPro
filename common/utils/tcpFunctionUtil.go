package utils

import (
	"encoding/binary"
	"unicode/utf16"
)

// ByteArrayToShortInt: 将字节数组中的两个字节转换为 short int。
// 参数 byteArrayData: 需要转换的字节数组。
// 参数 offset: 从字节数组的哪个位置开始转换。
// 返回值: 转换后的 short int 值。
func ByteArrayToShortInt(byteArrayData []byte, offset int) int {
	return (0&0xff)<<24 |
		(0&0xff)<<16 |
		(int(byteArrayData[offset]&0xff) << 8) |
		(int(byteArrayData[offset+1] & 0xff))
}

// ByteArrayToInt: 将字节数组中的四个字节转换为 int。
// 参数 byteArrayData: 需要转换的字节数组。
// 参数 offset: 从字节数组的哪个位置开始转换。
// 返回值: 转换后的 int 值。
func ByteArrayToInt(byteArrayData []byte, offset int) int {
	return (int(byteArrayData[offset]&0xff) << 24) |
		(int(byteArrayData[offset+1]&0xff) << 16) |
		(int(byteArrayData[offset+2]&0xff) << 8) |
		(int(byteArrayData[offset+3] & 0xff))
}

// IntToByteArray: 将 int 转换为 4 字节数组。
// 参数 intData: 需要转换的 int 数据。
// 返回值: 转换后的字节数组，长度为4。
func IntToByteArray(intData int) []byte {
	result := make([]byte, 4) // 创建长度为4的字节数组
	result[0] = byte((intData & 0xFF000000) >> 24)
	result[1] = byte((intData & 0xFF0000) >> 16)
	result[2] = byte((intData & 0xFF00) >> 8)
	result[3] = byte(intData & 0xFF)
	return result
}

// ShortToByteArray: 将 short int 转换为 2 字节数组。
// 参数 intData: 需要转换的 short int 数据（使用 int 表示）。
// 返回值: 转换后的字节数组，长度为2。
func ShortToByteArray(intData int) []byte {
	result := make([]byte, 2)
	result[0] = (byte)((intData & 0xFF00) >> 8)
	result[1] = (byte)(intData & 0xFF)
	return result
}

// StringToBytesUNICODE: 将字符串转换为 UTF-16 编码的字节数组（大端序）。
// 参数 ss: 需要转换的字符串。
// 返回值: 转换后的字节数组，UTF-16 编码，每个字符占用 2 字节。
func StringToBytesUNICODE(ss string) []byte {
	if ss == "" {
		return nil
	}
	// 使用 UTF-16 编码字符串
	utf16Encoded := utf16.Encode([]rune(ss))
	byteBuffer := make([]byte, len(utf16Encoded)*2)
	for i, v := range utf16Encoded {
		byteBuffer[i*2] = byte(v >> 8)
		byteBuffer[i*2+1] = byte(v)
	}
	return byteBuffer
}

// BytesTOStringUNICODE: 将 UTF-16 编码的字节数组转换为字符串。
// 参数 utf16Bytes: 需要转换的 UTF-16 字节数组，长度应为偶数。
// 返回值: 解码后的字符串。
func BytesTOStringUNICODE(utf16Bytes []byte) string {
	// 将 UTF-16 字节数组转换回 uint16 切片
	utf16Words := make([]uint16, len(utf16Bytes)/2)
	for i := 0; i < len(utf16Words); i++ {
		// 使用大端序将字节转换为 uint16
		utf16Words[i] = binary.BigEndian.Uint16(utf16Bytes[i*2:])
	}
	// 使用 utf16.Decode 将 uint16 切片解码回字符串
	recoveredStr := string(utf16.Decode(utf16Words))
	return recoveredStr
}

// StringToBytesUTF16: 与 StringToBytesUNICODE 类似，将字符串转换为 UTF-16 字节数组。
// 参数 ss: 需要转换的字符串。
// 返回值: 转换后的字节数组，UTF-16 编码，每个字符占用 2 字节。
func StringToBytesUTF16(ss string) []byte {
	if ss == "" {
		return nil
	}
	// 使用 UTF-16 编码字符串
	utf16Encoded := utf16.Encode([]rune(ss))
	byteBuffer := make([]byte, len(utf16Encoded)*2)
	for i, v := range utf16Encoded {
		byteBuffer[i*2] = byte(v >> 8)
		byteBuffer[i*2+1] = byte(v)
	}
	return byteBuffer
}
