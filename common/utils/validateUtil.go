package utils

import (
	"regexp"
	"strconv"
)

type Validator struct {
	Min   int
	Max   int
	Field string
	Value string
	Flags string
}

// CheckEmail 检查字符串是否为有效的邮箱格式
func CheckEmail(str string) bool {
	ma, err := regexp.MatchString("^[A-Za-z\\d]+([-_.][A-Za-z\\d]+)*@([A-Za-z\\d]+[-.])+[A-Za-z\\d]{2,4}$", str)
	if err != nil {
		return false
	}
	return ma
}

// CheckBool 检查字符串是否为布尔值
func CheckBool(str string) bool {

	_, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return true
}

// CheckFloat 检查字符串是否为浮点数
func CheckFloat(str string) bool {

	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}

// CheckLength 检查字符串长度是否在指定范围内
func CheckLength(str string, min, max int) bool {

	if min == 0 && max == 0 {
		return true
	}

	n := len(str)
	if n < min || n > max {
		return false
	}

	return true
}

// Verify 进行字段验证
func Verify(rules []Validator) (string, bool) {
	for _, val := range rules {
		if val.Flags == "alpha" && (!Ctype_alpha(val.Value) || !CheckLength(val.Value, val.Min, val.Max)) {
			return val.Field, false
		} else if val.Flags == "digit" && (!IsDigit(val.Value) || !CheckLength(val.Value, val.Min, val.Max)) {
			return val.Field, false
		} else if val.Flags == "alnum" && (!IsAlnum(val.Value) || !CheckLength(val.Value, val.Min, val.Max)) {
			return val.Field, false
		} else if val.Flags == "string" && !CheckLength(val.Value, val.Min, val.Max) {
			return val.Field, false
		} else if val.Flags == "bool" && !CheckBool(val.Value) {
			return val.Field, false
		} else if val.Flags == "mail" && !CheckEmail(val.Value) {
			return val.Field, false
		} else if val.Flags == "float" && !CheckFloat(val.Value) {
			return val.Field, false
		}
	}
	return "", true
}
