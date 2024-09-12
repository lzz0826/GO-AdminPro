package utils

import (
	"encoding/binary"
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
