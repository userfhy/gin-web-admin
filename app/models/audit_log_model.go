package model

import "gin-web-admin/utils/logging"

type AuditLog struct {
	BaseModel
	Category string `gorm:"size:20;index" json:"category"`
	UserID   uint   `gorm:"index" json:"user_id"`
	Username string `gorm:"size:60" json:"username"`
	IP       string `gorm:"size:64" json:"ip"`
	Path     string `gorm:"size:255" json:"path"`
	Method   string `gorm:"size:10" json:"method"`
	Status   int    `gorm:"type:int" json:"status"`
	Action   string `gorm:"size:120" json:"action"`
	Message  string `gorm:"type:text" json:"message"`
}

func (AuditLog) TableName() string {
	return TablePrefix + "audit_log"
}

func CreateAuditLog(log AuditLog) {
	if err := db.Create(&log).Error; err != nil {
		// 写入失败时不影响主流程，打印错误便于排查
		// 使用 logging 包避免循环依赖
		logging.Warnf("create audit log failed: %v", err)
	}
}
