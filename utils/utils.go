package utils

import (
	"math/rand"
	"time"
)

// RandStr 生成随机字符串,传入参数为获取的字符串长度
func RandStr(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < num; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
