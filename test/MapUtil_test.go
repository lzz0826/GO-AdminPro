package test

import (
	"AdminPro/common/utils"
	"fmt"
	"testing"
)

func TestGetOrDefault(t *testing.T) {
	strMap := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	fmt.Println(utils.GetOrDefault(strMap, "a", "default")) //  apple
	fmt.Println(utils.GetOrDefault(strMap, "c", "default")) //  default

	numMap := map[int]float64{
		1: 1.1,
		2: 2.2,
	}
	fmt.Println(utils.GetOrDefault(numMap, 1, 0.0)) //  1.1
	fmt.Println(utils.GetOrDefault(numMap, 3, 0.0)) //  0.0
}
