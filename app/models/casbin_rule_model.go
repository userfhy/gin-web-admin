package model

import "github.com/jinzhu/gorm"

type CasbinRule struct {
	BaseModel
	PType string `json:"p_type" gorm:"type:varchar(100);"`
	V0    string `json:"v0" gorm:"type:varchar(100);"`
	V1    string `json:"v1" gorm:"type:varchar(100);"`
	V2    string `json:"v2" gorm:"type:varchar(100);"`
	V3    string `json:"v3" gorm:"type:varchar(100);"`
	V4    string `json:"v4" gorm:"type:varchar(100);"`
	V5    string `json:"v5" gorm:"type:varchar(100);"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}

func CreatCasbin(casbin CasbinRule) error {
	db.NewRecord(casbin)
	res := db.Create(&casbin)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func GetCasbinRuleList(pageNum int, pageSize int, whereSql string, values []interface{}) ([]*CasbinRule, error) {
	var casbinRuleList []*CasbinRule
	err := db.Select("*").Where(whereSql, values...).Offset(pageNum).Limit(pageSize).Find(&casbinRuleList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return casbinRuleList, nil
}
