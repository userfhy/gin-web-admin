package casbin

import (
	model "gin-web-admin/app/models"

	"github.com/casbin/casbin/v3"
)

// SetupCasbin 初始化 Casbin 并返回实例，调用方自行保存引用。
func SetupCasbin() *casbin.SyncedEnforcer {
	return model.SetupCasbin()
}
