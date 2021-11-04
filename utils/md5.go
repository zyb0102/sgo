package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5加密
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
