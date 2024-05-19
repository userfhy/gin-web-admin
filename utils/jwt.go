package utils

import (
	"errors"
	"gin-web-admin/utils/setting"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"user_name"`
	RoleKey  string `json:"role_key"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// GenerateToken generates an access token used for auth
func GenerateToken(userClaims Claims) (string, time.Time, error) {
	return generateToken(userClaims, time.Second*5)
}

// GenerateRefreshToken generates a refresh token used for auth
func GenerateRefreshToken(userClaims Claims) (string, time.Time, error) {
	// RefreshToken 7天过期
	var timeExpiresNumber = time.Hour * 24 * 7

	nowTime := time.Now()
	expireTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location()).Add(timeExpiresNumber)
	return generateToken(userClaims, expireTime.Sub(nowTime))
}

// generateToken generates a JWT token with a specified duration from now
func generateToken(userClaims Claims, duration time.Duration) (string, time.Time, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(duration)

	claims := Claims{
		UserId:   userClaims.UserId,
		Username: userClaims.Username,
		RoleKey:  userClaims.RoleKey,
		IsAdmin:  userClaims.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gin-web-admin",
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, expireTime, err
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

// 验证 JWT token 是否有效
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
