package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Admin struct {
	ID                   uint   `json:"id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

// 生成JWT
func GenerateJWTAdmin(username, secretKey string, id uint) (string, error) {
	claims := User{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

// 解析jwt
func ParseJwtAdmin(tokenstring, secretKey string) (*User, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*User); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
