package utils

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"log"
)

// Utf82Utf16 将 UTF-8 编码的字节数据转换为 UTF-16 编码的字节数据。
func Utf82Utf16(data []byte) ([]byte, int, error) {
	charset := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM) // 使用大端序

	resultArr, n, err := transform.Bytes(charset.NewEncoder(), data)
	if err != nil {
		log.Fatalf("Error encoding string: %v", err)
		return nil, 0, err
	}
	return resultArr, n, nil
}

// Utf162Utf8 将 UTF-16 编码的字节数据转换为 UTF-8 编码的字节数据。
func Utf162Utf8(data []byte) ([]byte, int, error) {
	//str := "20010001"
	charset := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM) // 使用小端序

	// 使用 UTF16 編碼器將字串轉換為字節數組
	resultArr, n, err := transform.Bytes(charset.NewDecoder(), data)
	if err != nil {
		log.Fatalf("Error encoding string: %v", err)
		return nil, 0, err
	}
	return resultArr, n, nil
}

// Bytes2Int 将指定字节长度的数据转换为整数。
func Bytes2Int(bytes int, frame []byte) int {
	switch bytes {
	case 1:
		return int(frame[0])
	case 2:
		return int(binary.BigEndian.Uint16(frame))
	case 4:
		return int(binary.BigEndian.Uint32(frame))
	case 8:
		return int(binary.BigEndian.Uint64(frame))
	default:
		return 0
	}
}

// Int2Bytes 将整数转换为指定字节长度的字节数据。
func Int2Bytes(bytes int, length int) []byte {
	var frame []byte
	switch bytes {
	case 1:
		frame = append(frame, byte(length))
	case 2:
		frame = make([]byte, 2)
		binary.BigEndian.PutUint16(frame, uint16(length))
	case 4:
		frame = make([]byte, 4)
		binary.BigEndian.PutUint32(frame, uint32(length))
	case 8:
		frame = make([]byte, 8)
		binary.BigEndian.PutUint64(frame, uint64(length))
	default:

	}
	return frame
}

func TestBytes() {
	// 示例 1: UTF-8 到 UTF-16 转换
	utf8Data := []byte("Hello, 世界") // UTF-8 编码的字节数据
	utf16Data, _, err := Utf82Utf16(utf8Data)
	if err != nil {
		log.Fatalf("Error converting UTF-8 to UTF-16: %v", err)
	}
	fmt.Printf("UTF-8 to UTF-16: %v\n", utf16Data)

	// 示例 2: UTF-16 到 UTF-8 转换
	utf16Back, _, err := Utf162Utf8(utf16Data)
	if err != nil {
		log.Fatalf("Error converting UTF-16 to UTF-8: %v", err)
	}
	fmt.Printf("UTF-16 back to UTF-8: %s\n", string(utf16Back))

	// 示例 3: 字节到整数转换
	bytesData := []byte{0x00, 0x01} // 2 字节数据
	intValue := Bytes2Int(2, bytesData)
	fmt.Printf("Bytes to Int (2 bytes): %d\n", intValue)

	// 示例 4: 整数到字节转换
	intValue = 258                     // 示例整数值
	bytesData = Int2Bytes(2, intValue) // 转换为 2 字节
	fmt.Printf("Int to Bytes (2 bytes): %v\n", bytesData)
}
