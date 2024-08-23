package test

import (
	"AdminPro/common/utils"
	"fmt"
	"strconv"
	"testing"
)

func TestSumInDecimal(t *testing.T) {
	decimal := utils.SumInDecimal(1, 442.33, 321.4444)
	// 格式：'f' 表示十进制浮点数
	precision := 2 // 保留两位小数
	bitSize := 64  // 64 位浮点数（float64）
	fmt.Printf(strconv.FormatFloat(decimal, 'f', precision, bitSize))
}
