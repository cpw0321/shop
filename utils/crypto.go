// Copyright 2020 The shop Authors

// Package utils implements utils.
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// PKCS7UnPadding ...
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesCBCDecrypt ...
func AesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}

	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedData := make([]byte, len(encryptData))
	mode.CryptBlocks(decryptedData, encryptData)
	decryptedData = PKCS7UnPadding(decryptedData)
	return decryptedData, nil
}

// Md5 md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Encode base64编码
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode base64解码
func Base64Decode(str string) string {
	decodestr, _ := base64.StdEncoding.DecodeString(str)
	return string(decodestr)
}
