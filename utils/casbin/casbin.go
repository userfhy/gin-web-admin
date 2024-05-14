package casbin

import (
	model "gin-web-admin/app/models"

	"github.com/casbin/casbin/v2"
)

var (
	CasbinEnforcer *casbin.SyncedEnforcer
)

func SetupCasbin() *casbin.SyncedEnforcer {
	CasbinEnforcer = model.SetupCasbin()
	return CasbinEnforcer
}
