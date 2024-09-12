package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密, salt 为盐
func Md5SaltEncrypt(salt, value []byte) string {
	m5 := md5.New()
	m5.Write(value)
	m5.Write(salt)
	resByte := m5.Sum(nil)

	return hex.EncodeToString(resByte)
}

// md5 加密, salt 为盐
func Md5Encrypt(value []byte) string {
	m5 := md5.New()
	m5.Write(value)
	resByte := m5.Sum(nil)

	return hex.EncodeToString(resByte)
}

func Md5EncodeToString(s string) string {
	hexCode := md5.Sum([]byte(s))
	return hex.EncodeToString(hexCode[:])
}

func PhoneMd5(s string) string {
	return Md5Encrypt([]byte(s[1:4] + s + s[3:6] + "chat"))
}
