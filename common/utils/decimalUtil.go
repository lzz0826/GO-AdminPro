package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func SumInDecimal(values ...interface{}) float64 {
	result := decimal.New(0, 0)
	for _, v := range values {
		var value decimal.Decimal
		switch cv := v.(type) {
		case int:
			value = decimal.NewFromInt(int64(cv))
		case int64:
			value = decimal.NewFromInt(cv)
		case uint:
			value = decimal.NewFromInt(int64(cv))
		case float64:
			value = decimal.NewFromFloat(cv)
		case float32:
			value = decimal.NewFromFloat32(cv)
		case decimal.Decimal:
			value = cv
		default:
			panic(fmt.Sprintln("Invalid argument in SumInDecimal():", v))
		}
		result = result.Add(value)
	}
	fResult, _ := result.Rat().Float64()
	return fResult
}
