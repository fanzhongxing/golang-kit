package math

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	// 使用当前时间的纳秒表示来设置随机数生成器的种子
	seeded = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomStr 随机字符串
func RandomStr(length int) string {
	if length <= 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seeded.Intn(len(charset))]
	}
	return string(b)
}

// RandomNum 生成随机数字字符串
func RandomNum(length int) string {
	if length <= 0 {
		return "0"
	}
	// 随机生成一个1到999...之间的数字
	endNum := 0
	for i := 0; i < length; i++ {
		endNum = endNum*10 + 9
	}
	rn := seeded.Intn(endNum) + 1
	// 格式化数字为n位数的字符串，不足四位前面补0
	format := "%0" + fmt.Sprintf("%dd", length)
	return fmt.Sprintf(format, rn)
}
