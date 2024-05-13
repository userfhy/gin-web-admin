package casbin

import (
	"fmt"
	"gin-web-admin/utils/setting"

	"github.com/casbin/casbin/v2"

	//"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	//"strings"
)

var (
	CasbinEnforcer *casbin.Enforcer
)

func SetupCasbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapter(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	),
		true,
	)
	e, _ := casbin.NewEnforcer("conf/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	_ = e.LoadPolicy()

	// Check the permission.
	/*    check, _ := e.Enforce("alice", "data1", "read")
	      if check {
	          log.Println("通过权限")
	      } else {
	          log.Println("权限没有通过")
	      }*/

	CasbinEnforcer = e
	return e

	// Modify the policy.
	//e.AddPolicy("alice2", "data1", "read")
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	//e.SavePolicy()
}
