package middleware

import (
	"fmt"
	"net/http"

	"gin-web-admin/utils"

	"github.com/casbin/casbin/v3"
	"github.com/gin-gonic/gin"
)

func CasbinHandler(enforcer *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		if enforcer == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "casbin enforcer not initialized",
				"data": gin.H{},
			})
			c.Abort()
			return
		}

		claims, _ := c.Get("claims")
		user := claims.(*utils.Claims)

		username := user.Username
		roleKey := user.RoleKey
		obj := c.Request.URL.EscapedPath()
		act := c.Request.Method

		// 验证路由权限
		check, _ := enforcer.Enforce(roleKey, obj, act)
		if !check {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  fmt.Sprintf("[%s]对应权限[%s]没有[%s]路由的[%s]权限", username, roleKey, obj, act),
				"data": gin.H{},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
