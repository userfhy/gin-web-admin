package middleware

import (
	"fmt"
	"gin-web-admin/utils"
	"gin-web-admin/utils/casbin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		user := claims.(*utils.Claims)

		username := user.Username
		roleKey := user.RoleKey
		obj := c.Request.URL.EscapedPath()
		act := c.Request.Method

		// 验证路由权限
		check, _ := casbin.CasbinEnforcer.Enforce(roleKey, obj, act)
		if !check {
			//log.Println("权限没有通过")
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  fmt.Sprintf("[%s]对应权限[%s]没有[%s]路由的[%s]权限", username, roleKey, obj, act),
				"data": gin.H{},
			})
			c.Abort()
			return
		}

		c.Next()
		//log.Println(user, obj, act)
	}
}
