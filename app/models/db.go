package model

import "gorm.io/gorm"

// DB exposes the underlying gorm DB instance for internal use cases.
// NOTE: Prefer adding model-level functions instead of using DB directly when possible.
func DB() *gorm.DB {
	return db
}
