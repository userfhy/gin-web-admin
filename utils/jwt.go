package utils

import (
	"gin-test/utils/setting"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"user_name"`
	RoleKey  string `json:"role_key"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userClaims Claims) (string, error) {
	nowTime := time.Now()
	zeroTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location())
	// 明天零点
	expireTime := zeroTime.Add(time.Hour * 24)

	claims := Claims{
		userClaims.UserId,
		userClaims.Username,
		userClaims.RoleKey,
		userClaims.IsAdmin,
		jwt.StandardClaims{
			Issuer:    "gin-test",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
