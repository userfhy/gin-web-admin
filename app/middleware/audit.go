package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	model "gin-web-admin/app/models"
	"gin-web-admin/utils"

	"github.com/gin-gonic/gin"
)

const maxAuditPayloadBytes = 64 * 1024

type auditResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *auditResponseWriter) Write(data []byte) (int, error) {
	if w.body.Len() < maxAuditPayloadBytes {
		remaining := maxAuditPayloadBytes - w.body.Len()
		if len(data) > remaining {
			w.body.Write(data[:remaining])
		} else {
			w.body.Write(data)
		}
	}
	return w.ResponseWriter.Write(data)
}

// OperationLogger 记录非 GET 请求的操作审计日志。
func OperationLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestBody := readAuditRequestBody(c)
		writer := &auditResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		c.Writer = writer

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

			RequestHeaders:  marshalAuditValue(c.Request.Header),
			RequestBody:     requestBody,
			ResponseHeaders: marshalAuditValue(c.Writer.Header()),
			ResponseBody:    normalizeAuditPayload(writer.body.String()),
		}
		go model.CreateAuditLog(entry)
	}
}

func readAuditRequestBody(c *gin.Context) string {
	if c.Request == nil || c.Request.Body == nil {
		return ""
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return marshalAuditValue(map[string]string{"error": err.Error()})
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	return normalizeAuditPayload(string(body))
}

func normalizeAuditPayload(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	if len(raw) > maxAuditPayloadBytes {
		raw = raw[:maxAuditPayloadBytes]
	}
	var value any
	if err := json.Unmarshal([]byte(raw), &value); err == nil {
		return marshalAuditValue(value)
	}
	return raw
}

func marshalAuditValue(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}
	if len(data) > maxAuditPayloadBytes {
		data = data[:maxAuditPayloadBytes]
	}
	return string(data)
}
