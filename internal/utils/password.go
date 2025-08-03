package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// 生成随机盐值
func GenerateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// 使用SHA-256哈希算法对密码加盐
func HashPassword(password string, salt string) string {
	// 将密码和盐值组合
	combined := password + salt
	// 创建SHA-256哈希
	hash := sha256.Sum256([]byte(combined))
	// 将哈希值转换为base64编码的字符串
	return base64.StdEncoding.EncodeToString(hash[:])
}

// 验证密码
func VerifyPassword(password string, salt string, hashedPassword string) bool {
	// 对输入的密码进行相同的哈希处理
	computedHash := HashPassword(password, salt)
	// 比较计算出的哈希值与存储的哈希值
	return computedHash == hashedPassword
}
