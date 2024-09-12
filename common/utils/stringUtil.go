package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func GetBillNo(prefix string) string {
	return prefix + time.Now().Format(TimeYMDHM) + strings.ToUpper(GetRandomString(5))
}

func CreatePwd(pwd, salt string) string {
	return MD5(pwd + salt)
}

// StrToLower 将字符串转为小写并去掉空格
func StrToLower(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "", -1)
}

// StrToUpper 将字符串转为大写并去掉空格。
func StrToUpper(s string) string {
	return strings.Replace(strings.ToUpper(s), " ", "", -1)
}

func ConvertToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func StringToInt(str string, defalutValue int) int {
	if str == "" {
		return defalutValue
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return n
}

// 提交过来的参数转为小数点两位
func StringToDecimal(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	d64 := Precision(f64, 2, true)
	return d64
}

// FloatPrecision float 精度转换
func FloatPrecision(fStr string, prec int, round bool) (string, error) {
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		return "", err
	}

	f = Precision(f, prec, round)
	str := strconv.FormatFloat(f, 'f', prec, 64)

	return str, nil
}

// FloatPrecisionStr float 转换为 string 精度转换
func FloatPrecisionStr(f float64, prec int, round bool) string {
	ff := Precision(f, prec, round)
	str := strconv.FormatFloat(ff, 'f', prec, 64)

	return str
}

// FloatPrecision float 精度转换
func FloatFPrecision(fStr string, prec int, round bool) (float64, error) {
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		return 0, err
	}

	return Precision(f, prec, round), nil
}

// Precision 支持精度以及是否四舍五入, prec 保留的几位小数 round: true 为四舍五入, false 不是四舍五入
func Precision(f float64, prec int, round bool) float64 {
	// 需要加上对长度的校验, 否则直接用 math.Trunc 会有bug(1.14会变成1.13)
	arr := strings.Split(strconv.FormatFloat(f, 'f', -1, 64), ".")
	if len(arr) < 2 {
		return f
	}
	if len(arr[1]) <= prec {
		return f
	}
	pow10N := math.Pow10(prec)

	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}

	return math.Trunc((f)*pow10N) / pow10N
}

// 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func MapToQuery(m map[string]string) string {
	var sBuild strings.Builder
	for k, v := range m {
		if sBuild.Len() > 0 {
			sBuild.WriteString("&")
		}
		sBuild.WriteString(fmt.Sprint(k, "=", v))
	}
	return sBuild.String()
}

// 判断字符串是否包含汉字 true：包含，false不包含
func ContainsHan(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

// 生成随机字符串
func RandomLogNo() string {
	n := 10
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	suffix := string(b)
	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	formatted += suffix
	return formatted
}

// 获取source的字串,如果start小于0或者end大于source长度则返回""
// start:开始index，从0开始，包括0
// end:结束index，以end结束，但不包括end
func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

/**
* interface类型转成字符串
* @param interface{} 转换的值
* @return string 返回的字符串
 */
func InterfaceToString(inter interface{}) string {

	res := ""
	switch inter := inter.(type) {
	case bool:
		res = fmt.Sprintf("%t", inter)
	case int:
		res = fmt.Sprintf("%d", inter)
	case int64:
		res = fmt.Sprintf("%d", inter)
	case float64:
		res = strconv.FormatFloat(inter, 'f', -1, 64)
	case byte:
		res = fmt.Sprintf("%b", inter)
	case string:
		res = fmt.Sprintf("%s", inter)
	case *bool:
		res = fmt.Sprintf("%p", inter)
	case *int:
		res = fmt.Sprintf("%p", inter)
	case *int64:
		res = fmt.Sprintf("%p", inter)
	case *float64:
		res = fmt.Sprintf("%p", inter)
	case *string:
		res = fmt.Sprintf("%p", inter)
	}
	return res
}
