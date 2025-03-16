package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成6位随机数
func GenerateRandomNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 100000
	max := 999999
	return fmt.Sprintf("%06d", r.Intn(max-min+1)+min)
}
