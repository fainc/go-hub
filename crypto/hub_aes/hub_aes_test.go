package hub_aes

import (
	"fmt"
	"testing"
)

var (
	testCipherText = "测试加密文本"
	testAesKey     = "2byXZvtRqq0t6uXapUomDGZTfvJtebR9"
	testSpecMode   = "CTR"
	testSpecIv     = "1234567890123456"
)

func TestDefault(t *testing.T) {
	s, err := Encrypt(testCipherText, testAesKey) // 便捷方法内置返回base64数据
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
	d, err := Decrypt(s, testAesKey) // 便捷方法内置数据base64 decode
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)
}

func TestSpec(t *testing.T) {
	s, err := Encrypt(testCipherText, testAesKey, testSpecMode, testSpecIv) // 便捷方法内置返回base64数据
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
	d, err := Decrypt(s, testAesKey, testSpecMode, testSpecIv) // 便捷方法内置数据base64 decode
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)
}

func TestCBC(t *testing.T) {
	s, err := EncryptCBC([]byte(testCipherText), []byte(testAesKey), []byte(testSpecIv))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(encodeBase64(s))                                    // 传输或存储时一般需要 base64 encode
	d, err := DecryptCBC(s, []byte(testAesKey), []byte(testSpecIv)) // 以base64数据传输或存储时解密前需要 base64 decode
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(d))
}

func TestCTR(t *testing.T) {
	s, err := EncryptCTR([]byte(testCipherText), []byte(testAesKey), []byte(testSpecIv))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(encodeBase64(s))                                    // 传输或存储时一般需要 base64 encode
	d, err := DecryptCTR(s, []byte(testAesKey), []byte(testSpecIv)) // 以base64数据传输或存储时解密前需要 base64 decode
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(d))
}
