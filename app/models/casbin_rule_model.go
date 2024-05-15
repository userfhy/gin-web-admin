package model

import (
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type CasbinRuleM struct {
	BaseModel
	Ptype string `gorm:"type:varchar(300);"`
	V0    string `json:"v0" gorm:"type:varchar(100);uniqueIndex:unique_index"`
	V1    string `json:"v1" gorm:"type:varchar(100);uniqueIndex:unique_index"`
	V2    string `json:"v2" gorm:"type:varchar(100);uniqueIndex:unique_index"`
	V3    string `json:"v3" gorm:"type:varchar(100);uniqueIndex:unique_index"`
	V4    string `json:"v4" gorm:"type:varchar(100);uniqueIndex:unique_index"`
	V5    string `json:"v5" gorm:"type:varchar(100);uniqueIndex:unique_index"`
}

func (CasbinRuleM) TableName() string {
	return "casbin_rule"
}

func SetupCasbin() *casbin.SyncedEnforcer {
	// a, _ := gormadapter.NewAdapter(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Name,
	// ),
	// 	true,
	// )

	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRuleM{})
	e, _ := casbin.NewSyncedEnforcer("conf/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	// _ = e.LoadPolicy()

	// Refresh every 12 hours.
	e.StartAutoLoadPolicy(12 * time.Hour)

	// Check the permission.
	/*    check, _ := e.Enforce("alice", "data1", "read")
	      if check {
	          log.Println("通过权限")
	      } else {
	          log.Println("权限没有通过")
	      }*/

	// Modify the policy.
	// e.AddPolicy("alice2", "data1", "read")
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	// e.SavePolicy()

	return e

}

func CreatCasbin(casbin CasbinRuleM) error {
	res := db.Create(&casbin)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func GetCasbinRuleList(pageNum int, pageSize int, where map[string]interface{}) ([]*CasbinRuleM, error) {
	var casbinRuleList []*CasbinRuleM

	db, _ := BuildCondition(db, where)
	err := db.Select("*").Offset(pageNum).Limit(pageSize).Find(&casbinRuleList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return casbinRuleList, nil
}
