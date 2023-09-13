package tools

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// https://www.cnblogs.com/XY-Heruo/p/17328842.html
// https://juejin.cn/post/7219651766706159653

// 指定加密密钥
var jwtSecret = []byte("C2vbgqjAeOx6HBLdfm0uYthZXkl4JzTp")

// Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Phone    string `json:"phone"`
	jwt.RegisteredClaims
}

// 根据用户的用户名产生token
func GenerateToken(phone string, exp int) (string, error) {
	claims := Claims{
		phone,
		jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"note_app"},
			Subject:   "note_go",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exp) * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(jwtSecret))

	return s, err
}

// 解析token
func ParseToken(tokenstring string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// RefreshToken 通过 refresh token 刷新 atoken
func RefreshToken(atoken, rtoken string) (string, string, error) {
	// rtoken 无效直接返回
	claim, err := ParseToken(rtoken)
	if err != nil {
		return "", "", err
	}
	access_token, err := GenerateToken(claim.Phone, claim.Password, 2) // 2小时过期
	refresh_token, err := GenerateToken(claim.Phone, claim.Password, 168) // 7天过期

	return access_token, refresh_token, err
}
