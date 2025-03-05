package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

func IsValidEmail(email string) bool {
	// 定义一个正则表达式来匹配常见的邮箱格式
	// 这个正则表达式是基于RFC 5322标准简化而来的
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	reg := regexp.MustCompile(pattern)

	// 使用正则表达式检查字符串是否匹配
	return reg.MatchString(email)
}

const (
	iteration_time = 3         // 迭代次数
	memory         = 32 * 1024 // 内存消耗（32MB但实际并不会一定消耗32MB）
	threads        = 4         // 并行线程数
	keyLen         = 32        // 生成的哈希长度
)

func Encrypt(password string) (string, error) {
	// 生成随机盐值
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// 使用 Argon2id 生成哈希
	hash := argon2.IDKey([]byte(password), salt, iteration_time, memory, threads, keyLen)

	// 将盐值和哈希值编码为字符串
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	// 返回格式：算法版本$盐值$哈希值
	encodedPassword := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, iteration_time, threads, encodedSalt, encodedHash)

	return encodedPassword, nil
}

// ValidatePassword 验证密码
func ValidatePassword(password, encodedPassword string) error {
	// 解析存储的密码字符串
	parts := strings.Split(encodedPassword, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return errors.New("invalid encoded password format")
	}

	// 提取参数
	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil || version != argon2.Version {
		return errors.New("incompatible argon2 version")
	}

	var memory, iterations, threads uint32
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &threads)
	if err != nil {
		return errors.New("invalid argon2 parameters")
	}

	// 解码盐值和哈希值
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}

	// 使用相同的参数重新计算哈希值
	computedHash := argon2.IDKey([]byte(password), salt, iterations, memory, uint8(threads), uint32(len(storedHash)))

	// 比较哈希值
	if !compareHash(storedHash, computedHash) {
		return errors.New("invalid password")
	}

	return nil
}

func compareHash(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
