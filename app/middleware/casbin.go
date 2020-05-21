package middleware

import (
    "gin-test/utils"
    "gin-test/utils/casbin"
    "github.com/gin-gonic/gin"
    "log"
)

func CasbinHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, _ := c.Get("claims")
        user := claims.(*utils.Claims)

        username := user.Username
        obj := c.Request.URL.RequestURI()
        act := c.Request.Method

        // 验证路由权限
        check, _ := casbin.CasbinEnforcer.Enforce(username, obj, act)
        if check {
            log.Println("通过权限")
        } else {
            log.Println("权限没有通过")
        }
        //log.Println(user, obj, act)
    }
}