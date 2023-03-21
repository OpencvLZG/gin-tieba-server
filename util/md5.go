package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	b := []byte(str)
	md5.New()
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
