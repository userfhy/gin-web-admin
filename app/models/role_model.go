package model

import "github.com/jinzhu/gorm"

type Role struct {
	RoleId    uint      `gorm:"primary_key" json:"role_id"` // 角色编码
	CreatedAt JSONTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt JSONTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *JSONTime `sql:"index" json:"deleted_at"`
	RoleName  string    `gorm:"type:varchar(128);" json:"role_name"` // 角色名称
	IsAdmin   bool      `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"is_admin"`
	Status    int       `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	RoleKey   string    `gorm:"type:varchar(128);UNIQUE_INDEX;" json:"role_key"` //角色代码
	RoleSort  int       `gorm:"type:int(4);" json:"role_sort"`                   //角色排序
	Remark    string    `gorm:"type:varchar(255);" json:"remark"`                //备注
	Params    string    `gorm:"-" json:"params"`
	MenuIds   []int     `gorm:"-" json:"menu_ids"`
}

func (Role) TableName() string {
	return TablePrefix + "role"
}

func CreateRole(role Role) error {
	db.NewRecord(role)
	res := db.Create(&role)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func GetRoles(pageNum int, pageSize int, whereSql string, values []interface{}) ([]*Role, error) {
	var role []*Role
	err := db.Select("*").Where(whereSql, values...).Offset(pageNum).Limit(pageSize).Find(&role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return role, nil
}
