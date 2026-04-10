package middleware

import (
	"net/http"
	"strings"
	"time"

	model "gin-web-admin/app/models"
	"gin-web-admin/utils"

	"github.com/gin-gonic/gin"
)

// OperationLogger 记录非 GET 请求的操作审计日志。
func OperationLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if c.Request.Method == http.MethodGet {
			return
		}

		var (
			userID   uint
			username string
		)
		if claimsValue, exists := c.Get("claims"); exists {
			if claims, ok := claimsValue.(*utils.Claims); ok && claims != nil {
				userID = claims.UserId
				username = claims.Username
			}
		}

		status := c.Writer.Status()
		latency := time.Since(start)
		details := []string{latency.String()}
		if len(c.Errors) > 0 {
			details = append(details, c.Errors.String())
		}
		message := strings.Join(details, " | ")

		entry := model.AuditLog{
			Category: "operation",
			UserID:   userID,
			Username: username,
			IP:       c.ClientIP(),
			Path:     c.FullPath(),
			Method:   c.Request.Method,
			Status:   status,
			Action:   c.Request.Method + " " + c.FullPath(),
			Message:  message,
		}
		go model.CreateAuditLog(entry)
	}
}
