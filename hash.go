package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
	"time"
)

// HashPassword 生成安全的哈希化密码
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// HashCheckPassword 检查哈希密码与提交密码的一致性
func HashCheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NewUUID 生成微秒级唯一的 ID
func NewUUID(size int) string {
	ts := fmt.Sprintf("%d", time.Now().UnixMicro())
	tsSize := len(ts)
	if size < tsSize {
		panic(fmt.Sprintf("UUID 最小尺寸为 %d", tsSize))
	}
	if size-tsSize > 0 {
		ts = RandomString(size-tsSize) + ts
	}
	return ts
}

// RandomHashId 生成随机的 n 位 hash id
func RandomHashId(n int) string {
	id := make([]byte, n)
	for i := range id {
		id[i] = letterBytes[rand.IntN(len(letterBytes))]
	}
	return string(id)
}
