package roleService

import (
	model "gin-web-admin/app/models"
	sysService "gin-web-admin/app/service/v1/sys"
)

type SaveRoleMenuIdsReq struct {
	MenuIds []int `json:"menuIds"`
}

func (s *Service) GetRoleMenuIds(roleId uint) ([]int, error) {
	return model.GetRoleMenuIds(roleId)
}

func (s *Service) SaveRoleMenuIds(roleId uint, menuIds []int) error {
	if err := model.ReplaceRoleMenus(roleId, menuIds); err != nil {
		return err
	}
	sysService.InvalidateRouteCache()
	return nil
}
