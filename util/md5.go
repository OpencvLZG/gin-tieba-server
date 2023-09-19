package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func MD5(str string) string {
	b := []byte(str)
	md5.New()
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateKey() string {
	now := time.Now().UnixNano()
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%d", now)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateUID() string {
	// 确定字符集和长度
	rand.Seed(time.Now().UnixNano())
	const uidLen = 8
	const charset = "0123456789"

	// 生成随机字符串并使用 strconv 包将其转换为 int64
	var uidBuilder string
	for i := 0; i < uidLen; i++ {
		uidBuilder += string(charset[rand.Intn(len(charset))])
	}
	//uid, _ := strconv.ParseInt(uidBuilder, 10, 64)
	return uidBuilder
}

func GenerateID() int64 {
	// 确定字符集和长度
	rand.Seed(time.Now().UnixNano())
	const uidLen = 8
	const charset = "0123456789"

	// 生成随机字符串并使用 strconv 包将其转换为 int64
	var uidBuilder string
	for i := 0; i < uidLen; i++ {
		uidBuilder += string(charset[rand.Intn(len(charset))])
	}
	uid, _ := strconv.ParseInt(uidBuilder, 10, 64)
	return uid
}
