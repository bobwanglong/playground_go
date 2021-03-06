package main

import (
	"crypto/aes"
	"encoding/base64"
	"flag"
	"log"
)

func main() {
	s := flag.String("es", "", "待加密的数据")
	k := flag.String("k", "tsl0123456789tsl", "加密的密钥")
	encStr := flag.String("ds", "", "待解密的数据")
	flag.Parse()
	key := []byte(*k)

	if *s != "" && *k != "" {
		origData := []byte(*s)

		log.Println("原文：", string(origData), string(key))
		log.Println("------------------ ECB加密模式 --------------------")
		encrypted := AesEncryptECB(origData, key)
		// log.Println("密文(hex)：", hex.EncodeToString(encrypted))
		log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	}
	if *k != "" && *encStr != "" {
		log.Println("------------------ ECB解密模式 --------------------")
		ds, _ := base64.StdEncoding.DecodeString(*encStr)
		decrypted := AesDecryptECB([]byte(ds), key)
		log.Println("解密结果：", string(decrypted))
	}
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
