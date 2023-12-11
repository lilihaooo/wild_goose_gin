package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 加密
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	return hex.EncodeToString(m.Sum(nil))
}
