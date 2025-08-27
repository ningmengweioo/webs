package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret JWT密钥（实际项目中应从配置文件读取）
var JWTSecret = []byte("your_secret_key")

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, username string) (string, error) {
	// 创建claims
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 24小时后过期
		"iat":      time.Now().Unix(),
	}

	// 创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的编码后的字符串token
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
