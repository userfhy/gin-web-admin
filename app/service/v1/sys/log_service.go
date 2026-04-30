package sysService

import (
	"encoding/json"
	"fmt"
	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AuditLogQuery struct {
	Pagination utils.Pagination
	Category   string
	Username   string
	Module     string
	IP         string
	Status     *int
	StartTime  *time.Time
	EndTime    *time.Time
}

type LoginLogRow struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	IP        string `json:"ip"`
	Status    int    `json:"status"`
	Behavior  string `json:"behavior"`
	LoginTime string `json:"loginTime"`
}

type OperationLogRow struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Module        string `json:"module"`
	Summary       string `json:"summary"`
	IP            string `json:"ip"`
	Status        int    `json:"status"`
	OperatingTime string `json:"operatingTime"`
}

type SystemLogRow struct {
	ID          uint   `json:"id"`
	Category    string `json:"category"`
	Username    string `json:"username"`
	Module      string `json:"module"`
	URL         string `json:"url"`
	Method      string `json:"method"`
	IP          string `json:"ip"`
	Status      int    `json:"status"`
	Action      string `json:"action"`
	Message     string `json:"message"`
	TakesTime   int64  `json:"takesTime"`
	RequestTime string `json:"requestTime"`
}

type SystemLogDetail struct {
	SystemLogRow
	RequestHeaders  map[string]any `json:"requestHeaders"`
	RequestBody     any            `json:"requestBody"`
	ResponseHeaders map[string]any `json:"responseHeaders"`
	ResponseBody    any            `json:"responseBody"`
}

var durationValuePattern = regexp.MustCompile(`([0-9]+(?:\.[0-9]+)?)(ns|us|µs|ms|s|m|h)`)

func (s *Service) GetLoginLogList(query AuditLogQuery) (utils.PageResult, error) {
	items, total, err := s.listAuditLogs(query)
	if err != nil {
		return utils.PageResult{}, err
	}

	rows := make([]LoginLogRow, 0, len(items))
	for _, item := range items {
		rows = append(rows, LoginLogRow{
			ID:        item.ID,
			Username:  item.Username,
			IP:        item.IP,
			Status:    normalizeLogStatus(item.Status),
			Behavior:  fallbackText(item.Message, item.Action),
			LoginTime: formatJSONTime(item.CreatedAt),
		})
	}
	return query.Pagination.Result(rows, total), nil
}

func (s *Service) GetOperationLogList(query AuditLogQuery) (utils.PageResult, error) {
	items, total, err := s.listAuditLogs(query)
	if err != nil {
		return utils.PageResult{}, err
	}

	rows := make([]OperationLogRow, 0, len(items))
	for _, item := range items {
		rows = append(rows, OperationLogRow{
			ID:            item.ID,
			Username:      item.Username,
			Module:        logModule(item.Path, item.Category),
			Summary:       fallbackText(item.Action, item.Message),
			IP:            item.IP,
			Status:        normalizeLogStatus(item.Status),
			OperatingTime: formatJSONTime(item.CreatedAt),
		})
	}
	return query.Pagination.Result(rows, total), nil
}

func (s *Service) GetSystemLogList(query AuditLogQuery) (utils.PageResult, error) {
	items, total, err := s.listAuditLogs(query)
	if err != nil {
		return utils.PageResult{}, err
	}

	rows := make([]SystemLogRow, 0, len(items))
	for _, item := range items {
		rows = append(rows, buildSystemLogRow(item))
	}
	return query.Pagination.Result(rows, total), nil
}

func (s *Service) GetSystemLogDetail(id uint) (*SystemLogDetail, error) {
	if id == 0 {
		return nil, fmt.Errorf("id is required")
	}

	var item model.AuditLog
	err := s.db().Where("id = ?", id).First(&item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("system log not found")
		}
		return nil, err
	}

	row := buildSystemLogRow(&item)
	return &SystemLogDetail{
		SystemLogRow:    row,
		RequestHeaders:  parseLogObject(item.RequestHeaders),
		RequestBody:     parseLogPayload(item.RequestBody),
		ResponseHeaders: parseLogObject(item.ResponseHeaders),
		ResponseBody:    parseLogPayloadWithFallback(item.ResponseBody, item.Message),
	}, nil
}

func (s *Service) DeleteAuditLogs(category string, ids []uint) error {
	db := s.db().Model(&model.AuditLog{})
	if category != "" {
		db = db.Where("category = ?", category)
	}
	if len(ids) > 0 {
		db = db.Where("id IN ?", ids)
	}
	if category == "" && len(ids) == 0 {
		db = db.Session(&gorm.Session{AllowGlobalUpdate: true})
	}
	return db.Delete(&model.AuditLog{}).Error
}

func (s *Service) listAuditLogs(query AuditLogQuery) ([]*model.AuditLog, int64, error) {
	db := s.db().Model(&model.AuditLog{})
	db = applyAuditLogFilters(db, query)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []*model.AuditLog
	if err := db.Order("id DESC").Scopes(query.Pagination.Scope()).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func applyAuditLogFilters(db *gorm.DB, query AuditLogQuery) *gorm.DB {
	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}
	if username := strings.TrimSpace(query.Username); username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}
	if module := strings.TrimSpace(query.Module); module != "" {
		db = db.Where("path LIKE ? OR action LIKE ?", "%"+module+"%", "%"+module+"%")
	}
	if ip := strings.TrimSpace(query.IP); ip != "" {
		db = db.Where("ip LIKE ?", "%"+ip+"%")
	}
	if query.Status != nil {
		if *query.Status == 1 {
			db = db.Where("status >= 200 AND status < 400")
		} else {
			db = db.Where("status < 200 OR status >= 400")
		}
	}
	if query.StartTime != nil {
		db = db.Where("created_at >= ?", *query.StartTime)
	}
	if query.EndTime != nil {
		db = db.Where("created_at <= ?", *query.EndTime)
	}
	return db
}

func buildSystemLogRow(item *model.AuditLog) SystemLogRow {
	return SystemLogRow{
		ID:          item.ID,
		Category:    item.Category,
		Username:    item.Username,
		Module:      logModule(item.Path, item.Category),
		URL:         item.Path,
		Method:      item.Method,
		IP:          item.IP,
		Status:      item.Status,
		Action:      item.Action,
		Message:     item.Message,
		TakesTime:   extractLatencyMs(item.Message),
		RequestTime: formatJSONTime(item.CreatedAt),
	}
}

func formatJSONTime(t model.JSONTime) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func fallbackText(primary, secondary string) string {
	if strings.TrimSpace(primary) != "" {
		return primary
	}
	return secondary
}

func normalizeLogStatus(status int) int {
	if status >= 200 && status < 400 {
		return 1
	}
	return 0
}

func logModule(path, category string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return category
	}
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return category
	}
	parts := strings.Split(trimmed, "/")
	if len(parts) >= 3 && parts[0] == "v1" && parts[1] == "api" {
		return parts[2]
	}
	if len(parts) > 0 {
		return parts[0]
	}
	return category
}

func extractLatencyMs(message string) int64 {
	match := durationValuePattern.FindString(message)
	if match == "" {
		return 0
	}
	d, err := time.ParseDuration(strings.ReplaceAll(match, "µs", "us"))
	if err != nil {
		return 0
	}
	return d.Milliseconds()
}

func parseLogObject(raw string) map[string]any {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return map[string]any{}
	}
	var result map[string]any
	if err := json.Unmarshal([]byte(raw), &result); err == nil && result != nil {
		return result
	}
	return map[string]any{
		"raw": raw,
	}
}

func parseLogPayload(raw string) any {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return map[string]any{}
	}
	var result any
	if err := json.Unmarshal([]byte(raw), &result); err == nil {
		return result
	}
	return raw
}

func parseLogPayloadWithFallback(raw, fallback string) any {
	payload := parseLogPayload(raw)
	if text, ok := payload.(string); ok && strings.TrimSpace(text) == "" {
		return map[string]any{"message": fallback}
	}
	if object, ok := payload.(map[string]any); ok && len(object) == 0 && strings.TrimSpace(fallback) != "" {
		return map[string]any{"message": fallback}
	}
	return payload
}

func ParseLogStatus(value string) (*int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, nil
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("status 参数必须为数字")
	}
	if v != 0 && v != 1 {
		return nil, fmt.Errorf("status 参数必须为 0 或 1")
	}
	return &v, nil
}

func ParseDateRange(values []string) (*time.Time, *time.Time, error) {
	if len(values) == 0 {
		return nil, nil, nil
	}
	parseOne := func(raw string) (*time.Time, error) {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			return nil, nil
		}
		layouts := []string{
			time.RFC3339,
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05",
			"2006-01-02",
		}
		for _, layout := range layouts {
			if t, err := time.ParseInLocation(layout, raw, time.Local); err == nil {
				return &t, nil
			}
		}
		return nil, fmt.Errorf("无效时间格式: %s", raw)
	}

	start, err := parseOne(values[0])
	if err != nil {
		return nil, nil, err
	}
	var end *time.Time
	if len(values) > 1 {
		end, err = parseOne(values[1])
		if err != nil {
			return nil, nil, err
		}
	}
	return start, end, nil
}
