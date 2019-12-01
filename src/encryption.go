package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	//"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	r "math/rand"
	"time"
)

/**
 * Generate Key
 */
func GenerateKey() string {
	r.Seed(time.Now().UnixNano())
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()")
	b := make([]rune, 6)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

func CreateHash(key string) []byte {
	//hasher := md5.New()
	//hasher.Write([]byte(key))
	//return hex.EncodeToString(hasher.Sum(nil))
	return []byte("qwertyuiopasdfghjklzxcvbnm123456")
}

func Encrypt(key, pass []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := hex.EncodeToString(pass)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(key, pass []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(pass) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv, pass := pass[:aes.BlockSize], pass[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(pass, pass)
	data, err := hex.DecodeString(string(pass))
	if err != nil {
		return nil, err
	}

	return data, nil
}
