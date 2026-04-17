package middleware

import (
	"errors"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTHandler(userSvc *userService.Service) gin.HandlerFunc {
	if userSvc == nil {
		panic("jwt middleware requires user service")
	}
	return func(c *gin.Context) {
		var rCode int
		var data any

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
				rCode = code.ErrorAuthCheckTokenFail
				if errors.Is(err, jwt.ErrTokenMalformed) {
					// log.Println("That's not even a token")
				} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
					// log.Println("Invalid signature")
				} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
					rCode = code.ErrorAuthCheckTokenTimeout
					// log.Println("Timing is everything")
				}
			}
		}

		jwtCount, _ := userSvc.InBlockList(token)
		if jwtCount >= 1 {
			rCode = code.AuthTokenInBlockList
		}

		if rCode == code.SUCCESS && claims != nil && claims.UserId > 0 {
			active, activeErr := userSvc.HasActiveOnlineSession(claims.UserId, token)
			if activeErr != nil || !active {
				rCode = code.AuthTokenInBlockList
			}
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
		_ = userSvc.SaveOnlineSession(token, claims, c.ClientIP(), c.GetHeader("User-Agent"))
		c.Set("claims", claims)
		c.Next()
	}
}
