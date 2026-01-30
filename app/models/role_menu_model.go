package model

import (
	"errors"

	"gorm.io/gorm"
)

// RoleMenu 角色-菜单关联表
type RoleMenu struct {
	RoleId uint `gorm:"column:role_id;primaryKey" json:"role_id"`
	MenuId int  `gorm:"column:menu_id;primaryKey" json:"menu_id"`
}

func (RoleMenu) TableName() string {
	return TablePrefix + "role_menu"
}

func GetRoleMenuIds(roleId uint) ([]int, error) {
	var rows []RoleMenu
	err := db.Where("role_id = ?", roleId).Find(&rows).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	ids := make([]int, 0, len(rows))
	for _, r := range rows {
		ids = append(ids, r.MenuId)
	}
	return ids, nil
}

// ReplaceRoleMenus 全量替换某角色的菜单权限
func ReplaceRoleMenus(roleId uint, menuIds []int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleId).Delete(&RoleMenu{}).Error; err != nil {
			return err
		}
		if len(menuIds) == 0 {
			return nil
		}
		rows := make([]RoleMenu, 0, len(menuIds))
		for _, mid := range menuIds {
			rows = append(rows, RoleMenu{RoleId: roleId, MenuId: mid})
		}
		return tx.Create(&rows).Error
	})
}
