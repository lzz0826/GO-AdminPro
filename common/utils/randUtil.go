package utils

import (
	cryptorand "crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/xid"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes    = "0123456789"
	letterBytesLen = len(letterBytes)
	numberBytesLen = len(numberBytes)
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateId() string {
	return xid.New().String()
}

// 真随机 数字+字母大小写
func RealRand(size int) string {
	var buf strings.Builder

	for i := 0; i < size; i++ {
		result, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(letterBytesLen)))
		if err != nil {
			return ""
		}

		index := int(result.Int64())
		buf.WriteString(letterBytes[index : index+1])
	}
	str := buf.String()
	return str
}

/**
 * @Description: 真随机 数字
 * @param size
 * @return string
 */
func RealRandNumber(size int) string {
	var buf strings.Builder
	for i := 0; i < size; i++ {
		result, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(numberBytesLen)))
		if err != nil {
			return ""
		}
		index := int(result.Int64())
		buf.WriteString(numberBytes[index : index+1])
	}
	str := buf.String()
	return str
}

func GetTimeId() (id string) {
	id = fmt.Sprintf("%d%s", time.Now().UnixMilli(), RealRandNumber(3))
	return
}

// RandNum 获取一个 min-max(不含max) 区间的随机数
func RandNum(min, max int) int {
	if max < min {
		tmp := max
		max = min
		min = tmp
	} else if min == max {
		return min
	}
	rand.Seed(GetBjNowTime().UnixNano()) // 随机种子
	num := rand.Intn(max - min)
	return min + num
}

// RandFileName 生成文件名 并发模式下注意文件名会冲突
func RandFileName(fileType string) string {
	id := xid.New()
	fileName := fmt.Sprintf("%v_%d%s",
		id,
		RandNum(100000, 999999),
		strings.ToLower(fileType),
	)
	return fileName
}

func GetRandomStr(length int) string {
	str := uuid.New().String()
	fmt.Println(str)
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomNum(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
