package sysController

import (
	model "gin-web-admin/app/models"
	"gorm.io/gorm"
)

func getAllMenusFromDB() ([]menuRow, error) {
	menus, err := model.GetAllMenus(map[string]any{})
	if err != nil {
		return nil, err
	}
	return convertMenusToRows(menus), nil
}

func getMenusFromDBByRole(roleId uint) ([]menuRow, error) {
	// 查询该角色绑定的菜单ID，并补齐父链
	var menuIds []int
	if err := model.DB().Model(&model.RoleMenu{}).
		Where("role_id = ?", roleId).
		Pluck("menu_id", &menuIds).Error; err != nil {
		return nil, err
	}
	if len(menuIds) == 0 {
		return []menuRow{}, nil
	}

	// 取出所有菜单用于补父链
	allMenus, err := model.GetAllMenus(map[string]any{})
	if err != nil {
		return nil, err
	}

	byId := make(map[int]*model.Menu, len(allMenus))
	for _, m := range allMenus {
		byId[m.ID] = m
	}

	selected := make(map[int]struct{})
	for _, id := range menuIds {
		selected[id] = struct{}{}
		cur := byId[id]
		for cur != nil && cur.ParentID != 0 {
			selected[cur.ParentID] = struct{}{}
			cur = byId[cur.ParentID]
		}
	}

	filtered := make([]*model.Menu, 0, len(selected))
	for id := range selected {
		if m := byId[id]; m != nil {
			filtered = append(filtered, m)
		}
	}

	// 对齐原 GetAllMenus 的排序方式
	// 这里用 gorm 对 filtered id 再查一次按 rank 排序
	ids := make([]int, 0, len(filtered))
	for _, m := range filtered {
		ids = append(ids, m.ID)
	}

	var sorted []*model.Menu
	err = model.DB().Where("id IN ?", ids).Order("`rank` ASC").Find(&sorted).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return convertMenusToRows(sorted), nil
}

func convertMenusToRows(menus []*model.Menu) []menuRow {
	rows := make([]menuRow, 0, len(menus))
	for _, m := range menus {
		row := menuRow{
			ID:         m.ID,
			ParentID:   m.ParentID,
			Path:       m.Path,
			Name:       m.Name,
			Component:  m.Component,
			Redirect:   m.Redirect,
			Title:      m.Title,
			Icon:       m.Icon,
			ShowLink:   m.ShowLink,
			FrameSrc:   m.FrameSrc,
			ActivePath: m.ActivePath,
		}
		if m.Rank != nil {
			row.Rank = *m.Rank
		}
		row.KeepAlive = m.KeepAlive
		rows = append(rows, row)
	}
	return rows
}
