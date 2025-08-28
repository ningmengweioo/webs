package common

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (uint, error) {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWTSecret, nil
	})

	if err != nil {
		return 0, err
	}

	// 验证token并提取claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 检查token是否过期
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return 0, errors.New("token expired")
			}
		}

		// 提取用户ID
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid user_id in token")
		}

		return uint(userID), nil
	}

	return 0, errors.New("invalid token")
}

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization header中获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未提供认证token",
			})
			c.Abort()
			return
		}

		// 提取token（格式: Bearer token）
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token格式错误",
			})
			c.Abort()
			return
		}

		// 解析token
		userID, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 将用户ID存储在上下文中
		c.Set("userID", userID)
		c.Next()
	}
}
