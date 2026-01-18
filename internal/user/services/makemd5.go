package services

import (
	"crypto/md5"
	"encoding/hex"
)

func Makemd5(text string) string {
	pass := []byte(text)
	hash := md5.New()
	hash.Write(pass)
	md5string := hex.EncodeToString(hash.Sum(nil))
	return md5string
}
