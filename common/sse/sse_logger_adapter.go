package sse

import (
	"gin-web-admin/utils/logging"
	"strings"
)

// LogrusAdapter 实现SSE的Logger接口
type LogrusAdapter struct{}

func (a *LogrusAdapter) Printf(format string, v ...interface{}) {
	// 根据日志级别选择不同的记录方法
	switch {
	case isErrorLog(format, v...):
		logging.L().GetEntry().Errorf(format, v...)
	case isWarnLog(format, v...):
		logging.L().GetEntry().Warnf(format, v...)
	default:
		logging.L().GetEntry().Infof(format, v...)
	}
}

// 判断是否为错误日志（根据SSE实现中的关键字）
func isErrorLog(format string, _ ...any) bool {
	return containsAny(format, "error", "failed", "drop")
}

// 判断是否为警告日志
func isWarnLog(format string, _ ...interface{}) bool {
	return containsAny(format, "warn", "full", "timeout")
}

func containsAny(s string, substrs ...string) bool {
	for _, sub := range substrs {
		if strings.Contains(strings.ToLower(s), sub) {
			return true
		}
	}
	return false
}
