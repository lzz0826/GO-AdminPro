package utils

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
	"unicode/utf8"
)

// CheckFunc 校驗函數
// minLen: 最小長度
// maxLen: 對大長度
// str: 需要效驗的字段
// regexpStr: 正則表達式
// return error: 返還錯誤
func CheckFunc(minLen, maxLen int, str, regexpStr string) error {
	strLen := len(str)
	if strLen < minLen || strLen > maxLen {
		return fmt.Errorf("the length is invalid, %v-%v", minLen, maxLen)
	}

	// 判斷正則表達式是否有錯誤
	regCom, err := regexp.Compile(regexpStr)
	if err != nil {
		tmpStr := fmt.Sprintf("expression of regexp=%v is err: %v", regexpStr, err)
		return errors.New(tmpStr)
	}

	// 對 string 進行校驗
	matchFlag := regCom.MatchString(str)
	if !matchFlag {
		tmpStr := fmt.Sprintf("params not match, is invalid")
		return errors.New(tmpStr)
	}

	return nil
}

// NumLetter 數字和字母
func NumLetter(minLen, maxLen int, str string) error {
	regexpStr := "^[a-zA-Z0-9_-]*$"

	return CheckFunc(minLen, maxLen, str, regexpStr)
}

// NumCheck 數字
func NumCheck(minLen, maxLen int, str string) error {
	regexpStr := "^[0-9]*$"

	return CheckFunc(minLen, maxLen, str, regexpStr)
}

// NetterCheck 英文字母
func NetterCheck(minLen, maxLen int, str string) error {
	regexpStr := "^[a-zA-Z]*$"

	return CheckFunc(minLen, maxLen, str, regexpStr)
}

// 中國大陸手機號碼
func ChinaPhoneCheck(minLen, maxLen int, phone string) error {
	regexpStr := "^1[0-9]*$"
	return CheckFunc(minLen, maxLen, phone, regexpStr)
}

// 驗證qq
func CheckQQ(minLen, maxLen int, str string) error {
	regexpStr := "^[1-9][0-9]*$"

	return CheckFunc(minLen, maxLen, str, regexpStr)
}

// PwdCheck 密碼校驗 以數字和字母開頭 包含 _ -
func PwdCheck(minLen, maxLen int, str string) error {
	regexpStr := "^[a-zA-Z0-9][a-zA-Z0-9_-]*$"

	return CheckFunc(minLen, maxLen, str, regexpStr)
}

// 验证真实姓名 可以是中文或英文
func CheckRealName(realName string, min int, max int) bool {
	realName = strings.Trim(realName, "")
	count := utf8.RuneCountInString(realName)
	if count < min || count > max {
		return false
	}
	match_cn, err_cn := regexp.MatchString("^[\u4e00-\u9fa5]+([·?][\u4e00-\u9fa5]+)*$", realName)
	match_en, err_en := regexp.MatchString("^[a-zA-Z]+([\\s·?]?[a-zA-Z]+)+$", realName)
	if (!match_cn || err_cn != nil) && (!match_en || err_en != nil) {
		return false
	}
	return true
}

func isAlpha(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

// 判斷字符串是否為英文字母
func Ctype_alpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !isAlpha(r) {
			return false
		}
	}
	return true
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// 判斷字符串是否為數字
func IsDigit(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) {
			return false
		}
	}
	return true
}

// 判斷字符串是否為 字母 + 數字
func IsAlnum(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) && !isAlpha(r) {
			return false
		}
	}
	return true
}

func CheckIp(ip string) bool {
	ipNet := net.ParseIP(ip)
	return ipNet != nil
}

func CheckIpRange(ipRange string) bool {
	ipArr := strings.Split(ipRange, "-")
	if len(ipArr) != 2 {
		return false
	}
	ipLeftArr := strings.Split(ipArr[0], ".")
	if len(ipLeftArr) != 4 {
		return false
	}
	ipNetLeft := net.ParseIP(ipArr[0])
	if ipNetLeft == nil {
		return false
	}
	ipRight := strings.Join([]string{ipLeftArr[0], ipLeftArr[1], ipLeftArr[2], ipArr[1]}, ".")
	ipNetRight := net.ParseIP(ipRight)
	if ipNetRight == nil {
		return false
	}
	return true
}
