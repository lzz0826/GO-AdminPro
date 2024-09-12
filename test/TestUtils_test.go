package test

import (
	"AdminPro/common/utils"
	"fmt"
	"log"
	"testing"
)

func TestTestAes(t *testing.T) {
	//key := []byte("1234567890abcdef") // 16 字節密钥
	//key := []byte("1234567890abcdef12345678") // 24 字節密钥
	//key := []byte("1234567890abcdef1234567890abcdef")  // 32 字節密钥
	key := []byte("your_32_byte_key")
	iv := []byte("your_16_byte_key")
	plainText := []byte("This is a secret message!")
	encrypted, err := utils.AesCBCPk7EncryptBase64(plainText, key, iv)
	if err != nil {
		fmt.Println("Encryption failed:", err)
	} else {
		fmt.Println("Encrypted (Base64):", encrypted)
	}

	encryptedText := encrypted
	decrypted, err := utils.AesCBCPk7DecryptBase64(encryptedText, key, iv)
	if err != nil {
		fmt.Println("Decryption failed:", err)
	} else {
		fmt.Println("Decrypted message:", string(decrypted))
	}
}

func TestTestBytes(t *testing.T) {
	// 示例 1: UTF-8 到 UTF-16 转换
	utf8Data := []byte("Hello, 世界") // UTF-8 编码的字节数据
	utf16Data, _, err := utils.Utf82Utf16(utf8Data)
	if err != nil {
		log.Fatalf("Error converting UTF-8 to UTF-16: %v", err)
	}
	fmt.Printf("UTF-8 to UTF-16: %v\n", utf16Data)

	// 示例 2: UTF-16 到 UTF-8 转换
	utf16Back, _, err := utils.Utf162Utf8(utf16Data)
	if err != nil {
		log.Fatalf("Error converting UTF-16 to UTF-8: %v", err)
	}
	fmt.Printf("UTF-16 back to UTF-8: %s\n", string(utf16Back))

	// 示例 3: 字节到整数转换
	bytesData := []byte{0x00, 0x01} // 2 字节数据
	intValue := utils.Bytes2Int(2, bytesData)
	fmt.Printf("Bytes to Int (2 bytes): %d\n", intValue)

	// 示例 4: 整数到字节转换
	intValue = 258                           // 示例整数值
	bytesData = utils.Int2Bytes(2, intValue) // 转换为 2 字节
	fmt.Printf("Int to Bytes (2 bytes): %v\n", bytesData)
}

func TestVerify(t *testing.T) {
	// 定义一组需要验证的字段规则
	rules := []utils.Validator{
		{Min: 3, Max: 20, Field: "name", Value: "Alice", Flags: "alpha"},            // 姓名，要求是字母，长度3到20
		{Min: 1, Max: 3, Field: "age", Value: "25", Flags: "digit"},                 // 年龄，要求是数字，长度1到3
		{Min: 5, Max: 50, Field: "email", Value: "test@example.com", Flags: "mail"}, // 邮箱，格式验证
		{Min: 0, Max: 10, Field: "price", Value: "19.99", Flags: "float"},           // 价格，要求是浮点数
		{Field: "status", Value: "true", Flags: "bool"},                             // 状态，要求是布尔值
	}
	// 调用验证方法
	field, isValid := utils.Verify(rules)
	// 检查验证结果
	if isValid {
		fmt.Println("所有字段验证通过")
	} else {
		fmt.Printf("字段 %s 验证失败\n", field)
	}
}

func TestRand(t *testing.T) {
	realRandNumber := utils.RealRandNumber(20)
	fmt.Println("realRandNumber : ", realRandNumber)

	realRand := utils.RealRand(20)
	fmt.Println("realRand : ", realRand)

	num := utils.RandNum(6, 8)
	fmt.Println("num : ", num)

}

func TestPath(t *testing.T) {

	host := utils.IsValidHost("http://www.youtube.com/watch?v=I5iBi1kcGfI/")
	fmt.Println("host : ", host)

	hostname := "https://example.com"
	relativePath := "api/v1/users"
	relativePath2 := "test/v6/order"
	// 拼接域名和相对路径
	fullUrl := utils.BindUrl(hostname, relativePath, relativePath2)
	fmt.Println(fullUrl) // 输出: https://example.com/api/v1/users

	originalUrl := "https://oldsite.com/api/v1/resource"
	newHost := "https://newsite.com"

	// 替换资源的域名
	newUrl := utils.ReplaceHost(originalUrl, newHost)
	fmt.Println(newUrl) // 输出: https://newsite.com/api/v1/resource

}

func TestMd5(t *testing.T) {
	salt := []byte("random_salt")
	value := []byte("my_password")
	// 使用盐值加密
	md5SaltEncrypt := utils.Md5SaltEncrypt(salt, value)
	fmt.Println("Md5SaltEncrypt:", md5SaltEncrypt)

	// 使用普通 MD5 加密
	md5Encrypt := utils.Md5Encrypt(value)
	fmt.Println("Md5Encrypt:", md5Encrypt)

	value2 := "my_password"
	// 将字符串加密为 MD5
	md5EncodeToString := utils.Md5EncodeToString(value2)
	fmt.Println("Md5EncodeToString:", md5EncodeToString)

	phone := "1234567890"
	// 手机号相关的 MD5 加密
	phoneMd5 := utils.PhoneMd5(phone)
	fmt.Println("PhoneMd5:", phoneMd5)

}
