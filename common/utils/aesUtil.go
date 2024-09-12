package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

//AES对称加密
//Key（密钥）：
//必須是 16 字節（128 位）、24 字節（192 位）或 32 字節（256 位）。
//密钥必須保密且足夠隨機，避免簡單密钥。
//可以通過哈希算法或密钥派生函數來生成。

//IV（初始化向量）：
//必須是 16 字節（128 位），與 AES 算法的塊大小相同。
//每次加密都應該使用新的隨機 IV。
//IV 不需要保密，但必須與密文一起傳輸，以便解密。

// AesCBCEncrypt aes cbc模式加密 並且 base64加碼
func AesCBCPk7EncryptBase64(orgData, key, iv []byte) (string, error) {
	encryptBytes, err := AesCBCPk7Encrypt(orgData, key, iv)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(encryptBytes)

	return str, nil
}

// AesCBCPk7Decrypt base64解碼, aes cbc模式解碼
func AesCBCPk7DecryptBase64(orgData string, key, iv []byte) ([]byte, error) {
	// 進行 base64 解碼
	orgByte, err := base64.StdEncoding.DecodeString(orgData)
	if err != nil {
		return []byte(""), err
	}

	decryptBytes, err := AesCBCPk7Decrypt(orgByte, key, iv)
	if err != nil {
		return []byte(""), err
	}

	return []byte(decryptBytes), nil
}

// Aes cbc 加密, pkcs7 填充
func AesCBCPk7Encrypt(origData, key []byte, iv []byte) ([]byte, error) {
	if len(origData) < 1 {
		return []byte(""), errors.New("crypted is empty")
	}
	if len(key) < 1 {
		return []byte(""), errors.New("key is empty")
	}
	if len(iv) < 1 {
		return []byte(""), errors.New("iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// PKCS7 填充
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func GetRealString(encodingAesKey string, data string) (string, error) {
	dataTmp, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	var md5Str = Md5EncodeToString(encodingAesKey)
	rs, err := AesCBCPk7Decrypt(dataTmp, getAesKey(md5Str), getIv(md5Str))
	if err != nil {
		return "", err
	}
	return rs, nil
}

func getAesKey(key string) []byte {
	if len(key) != 32 {
		panic("error secret key")
	}
	return []byte(key[2:7] + key[11:15] + key[18:25])
}

func getIv(key string) []byte {
	if len(key) != 32 {
		panic("error secret key")
	}
	return []byte(key[4:9] + key[16:23] + key[25:29])
}

// Aes cbc 解密, pkcs7 填充
func AesCBCPk7Decrypt(encryption, key []byte, iv []byte) (string, error) {
	if len(encryption) < 1 {
		return "", errors.New("encryption is empty")
	}
	if len(key) < 1 {
		return "", errors.New("key is empty")
	}
	if len(iv) < 1 {
		return "", errors.New("iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 加入判斷條件防止 panic
	blockSize := block.BlockSize()
	if len(key) < blockSize {
		return "", errors.New("key too short")
	}
	if len(encryption)%blockSize != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encryption))
	blockMode.CryptBlocks(origData, encryption)
	origData = PKCS7UnPadding(origData, blockSize)
	return string(origData), nil
}

// PKCS7 填充
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	if length-unPadding < 0 {
		return []byte("")
	}
	return plantText[:(length - unPadding)]
}

// 生成隨機的 IV，長度應與 AES 的區塊大小一致（16 字節）
func GenerateRandomIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize) // AES 的區塊大小是 16 字節
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}
