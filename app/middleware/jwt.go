package middleware

import (
	userService "gin-test/app/service/v1/user"
	"gin-test/utils"
	"gin-test/utils/code"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rCode int
		var data interface{}

		rCode = code.SUCCESS

		var token = c.Query("token")

		var authorization = c.GetHeader("Authorization")
		if authorization != "" {
			parts := strings.SplitN(authorization, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}

		var claims *utils.Claims
		var err error
		if token == "" {
			rCode = code.TokenInvalid
		} else {
			claims, err = utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					rCode = code.ErrorAuthCheckTokenTimeout
				default:
					rCode = code.ErrorAuthCheckTokenFail
				}
			}
		}

		jwtCount, _ := userService.InBlockList(token)
		if jwtCount >= 1 {
			rCode = code.AuthTokenInBlockList
		}

		if rCode != code.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": rCode,
				"msg":  code.GetMsg(rCode),
				"data": data,
			})

			c.Abort()
			return
		}
		// 存储登录用户信息
		c.Set("claims", claims)
		c.Next()
	}
}
