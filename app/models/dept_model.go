package model

import (
	"errors"

	"gorm.io/gorm"
)

type Dept struct {
	ID       int    `gorm:"primaryKey;comment:主键ID" json:"id"`
	ParentID int    `gorm:"comment:父部门ID" json:"parentId"`
	DeptName string `gorm:"type:varchar(128);comment:部门名称" json:"deptName"`
	OrderNum int    `gorm:"comment:排序" json:"orderNum"`
	Leader   string `gorm:"type:varchar(64);comment:负责人" json:"leader"`
	Phone    string `gorm:"type:varchar(32);comment:联系电话" json:"phone"`
	Email    string `gorm:"type:varchar(128);comment:邮箱" json:"email"`
	Status   int    `gorm:"type:int(1);DEFAULT:1;NOT NULL;comment:状态(1启用0停用)" json:"status"`
	Remark   string `gorm:"type:varchar(255);comment:备注" json:"remark"`
}

func (Dept) TableName() string {
	return TablePrefix + "dept"
}

func GetDept(where map[string]any) (*Dept, error) {
	var dept Dept
	err := db.Where(where).First(&dept).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &dept, nil
}

func GetAllDepts(where map[string]any) ([]*Dept, error) {
	var depts []*Dept
	query := db.Model(&Dept{})
	if len(where) > 0 {
		query = query.Where(where)
	}
	err := query.Order("order_num ASC").Order("id ASC").Find(&depts).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return depts, nil
}

func CreateDept(dept Dept) error {
	return db.Create(&dept).Error
}

func UpdateDept(id int, data map[string]any) error {
	return db.Model(&Dept{}).Where("id = ?", id).Updates(data).Error
}

func DeleteDept(id int) error {
	// 删除策略：如果存在子部门则禁止删除
	var cnt int64
	if err := db.Model(&Dept{}).Where("parent_id = ?", id).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("has child dept")
	}

	return db.Delete(&Dept{}, id).Error
}
