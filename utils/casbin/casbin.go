package casbin

import (
	model "gin-web-admin/app/models"

	"github.com/casbin/casbin/v2"
)

var (
	CasbinEnforcer *casbin.Enforcer
)

func SetupCasbin() *casbin.Enforcer {
	CasbinEnforcer = model.SetupCasbin()
	return CasbinEnforcer
}
