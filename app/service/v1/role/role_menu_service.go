package roleService

import model "gin-web-admin/app/models"

type SaveRoleMenuIdsReq struct {
	MenuIds []int `json:"menuIds"`
}

func GetRoleMenuIds(roleId uint) ([]int, error) {
	return model.GetRoleMenuIds(roleId)
}

func SaveRoleMenuIds(roleId uint, menuIds []int) error {
	return model.ReplaceRoleMenus(roleId, menuIds)
}
