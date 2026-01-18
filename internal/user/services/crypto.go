package services

import (
	"fmt"
	"github.com/AlexanderGrom/componenta/crypt"
)

func CryptoEncrypt(text string, key string) string {
	fmt.Println(key)
	c, _ := crypt.Encrypt(text, key)

	return c
}

func CryptoDecrypt(text string, key string) string {
	fmt.Println(key)
	s, _ := crypt.Decrypt(text, key)

	return s
}
