package utils

import (
    "gin-test/utils/setting"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
    UserId uint `json:"user_id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string, userId uint) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(3 * time.Hour)

    claims := Claims{
        userId,
        username,
        jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    "gin-test",
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