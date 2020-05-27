package middleware

import (
    "fmt"
    "gin-test/utils"
    "gin-test/utils/casbin"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CasbinHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, _ := c.Get("claims")
        user := claims.(*utils.Claims)

        username := user.Username
        obj := c.Request.URL.EscapedPath()
        act := c.Request.Method

        // 验证路由权限
        check, _ := casbin.CasbinEnforcer.Enforce(username, obj, act)
        if !check {
             //log.Println("权限没有通过")
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": http.StatusUnauthorized,
                "msg": fmt.Sprintf("[%s]没有[%s]路由的[%s]权限", username, obj, act),
                "data": gin.H{},
            })
            c.Abort()
            return
        }

        c.Next()
        //log.Println(user, obj, act)
    }
}