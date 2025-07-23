package service

import (
	model "gin-web-admin/app/models"
)

func CreateMenu(menu model.Menu) error {
	return model.CreateMenu(menu)
}

func UpdateMenu(id int, data map[string]any) error {
	return model.UpdateMenu(id, data)
}

func DeleteMenu(id int) error {
	return model.DeleteMenu(id)
}

func GetMenu(where map[string]any) (*model.Menu, error) {
	return model.GetMenu(where)
}

func GetMenuList(pageNum, pageSize int, where map[string]any) ([]*model.Menu, error) {
	return model.GetMenuList(pageNum, pageSize, where)
}

func GetAllMenus(where map[string]any) ([]*model.Menu, error) {
	return model.GetAllMenus(where)
}
