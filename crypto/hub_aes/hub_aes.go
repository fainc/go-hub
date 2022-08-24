// Package hub_aes 基于标准库的AES加密包 支持CBC/CTR 加解密
package hub_aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"

	"github.com/fainc/go-hub/encode/hub_base64"
)

// Decrypt 便捷AES解密
// 传入 base64 sting, 返回解密后数据 string
// 默认使用CBC模式, 默认iv采用密钥前16位
// 使用 spec[0] 指定模式(CBC/CTR), spec[1] 指定iv
func Decrypt(encryptedStr, aesKey string, spec ...string) (string, error) {
	key := []byte(aesKey)
	if len(key) < 16 {
		return "", errors.New("AES密钥不足16位")
	}
	mode := "CBC"
	if len(spec) > 0 {
		mode = spec[0]
	}
	i := key[0:16]
	if len(spec) > 1 {
		i = []byte(spec[1])
	}
	data, err := decodeBase64(encryptedStr)
	if err != nil {
		return "", err
	}
	switch mode {
	case "CBC":
		ret, err := DecryptCBC(data, key, i)
		return string(ret), err
	case "CTR":
		ret, err := DecryptCTR(data, key, i)
		return string(ret), err
	default:
		return "", errors.New("mode仅支持CBC/CTR")
	}
}

// Encrypt 便捷AES加密
// 传入待加密数据 sting, 返回加密后base64数据 string
// 默认使用CBC模式, 默认iv采用密钥前16位
// 使用 spec[0] 指定模式(CBC/CTR), spec[1] 指定iv
func Encrypt(cipherText, aesKey string, spec ...string) (string, error) {
	key := []byte(aesKey)
	if len(key) < 16 {
		return "", errors.New("AES密钥不足16位")
	}
	mode := "CBC"
	if len(spec) > 0 {
		mode = spec[0]
	}
	i := key[0:16]
	if len(spec) > 1 {
		i = []byte(spec[1])
	}
	switch mode {
	case "CBC":
		ret, err := EncryptCBC([]byte(cipherText), key, i)
		if err != nil {
			return "", err
		}
		return encodeBase64(ret), nil
	case "CTR":
		ret, err := EncryptCTR([]byte(cipherText), key, i)
		if err != nil {
			return "", err
		}
		return encodeBase64(ret), nil
	default:
		return "", errors.New("mode仅支持CBC/CTR")
	}
}

// EncryptCBC 标准CBC模式加密
func EncryptCBC(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, fmt.Errorf("ErrAesDecryptBlockSize:%d", blockSize)
	}
	plainText = pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

// EncryptCTR 标准CTR模式加密
func EncryptCTR(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, fmt.Errorf("ErrAesDecryptBlockSize:%d", blockSize)
	}
	plainText = pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCTR(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.XORKeyStream(cipherText, plainText)
	return cipherText, nil
}

// DecryptCBC AESDecryptCBC AES CBC模式解密
func DecryptCBC(encrypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, fmt.Errorf("ErrAesDecryptBlockSize:%d", blockSize)
	}
	origData := make([]byte, len(encrypted))
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(origData, encrypted)
	return pkcs5UnPadding(origData, block.BlockSize())
}

// DecryptCTR AESDecryptCTR AES CTR模式解密
func DecryptCTR(encrypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, fmt.Errorf("ErrAesDecryptBlockSize:%d", blockSize)
	}
	origData := make([]byte, len(encrypted))
	blockMode := cipher.NewCTR(block, iv)
	blockMode.XORKeyStream(origData, encrypted)
	return pkcs5UnPadding(origData, block.BlockSize())
}

func decodeBase64(in string) ([]byte, error) {
	return hub_base64.Decode([]byte(in))
}
func encodeBase64(in []byte) string {
	return hub_base64.EncodeString(string(in))
}

func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if blockSize <= 0 {
		return nil, errors.New("invalid blocklen")
	}

	if length%blockSize != 0 || length == 0 {
		return nil, errors.New("invalid data len")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, errors.New("invalid padding")
	}

	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, errors.New("invalid padding")
		}
	}

	return src[:(length - unpadding)], nil
}
