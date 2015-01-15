// Package gotools implements a simple golang tools package.
package gotools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type AES struct {
	// 对称算法的密钥。必须是16,24,32位分别代表128,192,256位加密。
	Key []byte
	// 对称算法的初始化向量。大小为16。
	IV []byte
}

func (a AES) Padding(data []byte) []byte {
	padLen := 16 - len(data)%16
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padText...)
}

func (a AES) UnPadding(data []byte) []byte {
	dataLen := len(data)
	lastData := int(data[dataLen-1])

	return data[:(dataLen - lastData)]
}

func (a AES) Encrypt(text string) (string, error) {
	data := []byte(text)
	dataByte, err := a.EncryptBytes(data)
	if err != nil {
		return "", err
	}
	return string(dataByte), nil
}

func (a AES) Decrypt(text string) (string, error) {
	data := []byte(text)
	dataByte, err := a.DecryptBytes(data)
	if err != nil {
		return "", err
	}
	return string(dataByte), nil
}

func (a AES) EncryptBytes(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	paddingData := a.Padding(data)
	blockMode := cipher.NewCBCEncrypter(block, a.IV)
	result := make([]byte, len(paddingData))
	blockMode.CryptBlocks(result, paddingData)
	return result, nil
}

func (a AES) DecryptBytes(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(data))
	blockMode := cipher.NewCBCDecrypter(block, a.IV)
	blockMode.CryptBlocks(result, data)
	paddingResult := a.UnPadding(result)
	return paddingResult, nil
}
