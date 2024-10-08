package test

import (
	"AdminPro/common/utils"
	"fmt"
	"log"
	"testing"
)

func TestPickAll(t *testing.T) {
	// 模拟的字节数组（实际使用中应从数据源获取）
	bt2 := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	// 模拟的int2二维数组，itemId 和 数据类型的对应关系
	int2 := [][]int{
		{1, utils.TYPE_INT_1},        // itemId 1，类型是1字节整数
		{2, utils.TYPE_INT_4},        // itemId 2，类型是4字节整数
		{3, utils.TYPE_STRING_UTF16}, // itemId 3，类型是UTF-16字符串
	}

	// 调用PickAll函数
	result := utils.PickAll(bt2, int2)
	fmt.Printf("PickAll结果:%d", result[2])

	// 输出解析结果
	fmt.Printf("PickAll结果: %+v\n", result)
}

func TestPickAll2(t *testing.T) {
	// 模拟的int2二维数组，包含itemId、值和类型
	int2 := [][]interface{}{
		{1, 10, utils.TYPE_INT_1},            // itemId 1，值为10，类型是1字节整数
		{2, 123456, utils.TYPE_INT_4},        // itemId 2，值为123456，类型是4字节整数
		{3, "示例文本", utils.TYPE_STRING_UTF16}, // itemId 3，值为"示例文本"，类型是UTF-16字符串
	}

	// 请求码（随意设定）
	requestCode := 100

	// 调用PackAll函数
	packedData := utils.PackAll(int2, requestCode)

	// 输出打包后的字节数组
	fmt.Printf("PackAll结果: %+v\n", packedData)
}

/*
测试 functionUtil
*/
func TestByteToInt(t *testing.T) {
	// 示例字节数组
	byteArray := []byte{
		0x44, 0x5A, 0x50, 0x4B, // 协议头 'DZPK'
		//0x01, 0x00, // short int 示例 等同 1, 0
		1, 0, // short int 示例
		0x00, 0x00, 0x00, 0x0F, // int 示例
	}
	result := byteArray[1:5]
	protocolHeader := string(result)
	fmt.Printf("协议头: %s\n", protocolHeader)

	// 从字节数组中提取 short int (2个字节)
	shortValue := utils.ByteArrayToShortInt(byteArray, 4)
	fmt.Printf("Extracted short int: %d\n", shortValue)

	// 从字节数组中提取 int (4个字节)
	intValue := utils.ByteArrayToInt(byteArray, 6)
	fmt.Printf("Extracted int: %d\n", intValue)
}

func TestIntToBytes(t *testing.T) {
	// 示例 int 数据
	intValue := 1025
	shortValue := 513

	// 将 int 转换为 4 字节数组
	intBytes := utils.IntToByteArray(intValue)
	fmt.Printf("Int %d 转换为字节数组: %v\n", intValue, intBytes)
	// 将 short int 转换为 2 字节数组
	shortBytes := utils.ShortToByteArray(shortValue)
	fmt.Printf("Short int %d 转换为字节数组: %v\n", shortValue, shortBytes)

	// 打印字节数组的十六进制表示
	fmt.Printf("Int 转换为字节数组的十六进制表示: %X\n", intBytes)
	fmt.Printf("Short 转换为字节数组的十六进制表示: %X\n", shortBytes)
}

func TestStringToBytes(t *testing.T) {
	// 字符串转换为 UTF-16 编码的字节数组
	str := "Hello, World!"
	utf16Bytes := utils.StringToBytesUNICODE(str)
	fmt.Printf("UTF-16 encoded bytes: %v\n", utf16Bytes)

	bytes := utils.BytesTOStringUNICODE(utf16Bytes)
	// 将 UTF-16 字节数组转换回字符串（
	fmt.Printf("Recovered string: %s\n", bytes)

}

/*
测试 aesUtil
*/
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

/*
测试 bytesUtil
*/
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

/*
测试 validateUtil
*/
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

/*
 测试 randUtil
*/

func TestRand(t *testing.T) {
	realRandNumber := utils.RealRandNumber(20)
	fmt.Println("realRandNumber : ", realRandNumber)

	realRand := utils.RealRand(20)
	fmt.Println("realRand : ", realRand)

	num := utils.RandNum(6, 8)
	fmt.Println("num : ", num)

}

/*
 测试 pathUtil
*/

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

/*
 测试 Md5Util
*/

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
