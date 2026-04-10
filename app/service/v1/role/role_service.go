package roleService

import (
	model "gin-web-admin/app/models"
	sysService "gin-web-admin/app/service/v1/sys"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
	"gin-web-admin/utils/logging"
)

type RoleStruct struct {
	Pagination utils.Pagination
	Conditions map[string]any
}

type NewRoleStruct struct {
	RoleKey string `json:"role_key" form:"role_key" validate:"required,min=4,max=10" minLength:"4" maxLength:"10"`
}

type UpdateRoleStruct struct {
	// 角色名称
	RoleName string `json:"role_name" form:"role_name" validate:"required,min=4,max=10" minLength:"4" maxLength:"10"`
	// 备注
	Remark string `json:"remark" form:"remark" validate:"omitempty,min=4,max=100" minLength:"4" maxLength:"100"`
}

type CreateRoleStruct struct {
	NewRoleStruct
	UpdateRoleStruct
}

type Service struct {
	store *data.Store
}

var defaultService *Service

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func SetDefaultService(s *Service) {
	defaultService = s
}

func serviceInstance() *Service {
	if defaultService == nil {
		panic("role service not initialized")
	}
	return defaultService
}

func DeleteRole(roleId uint) bool {
	return serviceInstance().DeleteRole(roleId)
}

func (s *Service) DeleteRole(roleId uint) bool {
	wheres := make(map[string]any)
	wheres["role_id"] = roleId
	_, rowsAffected := model.SoftDelete(&model.Role{RoleId: roleId})
	if rowsAffected == 0 {
		logging.Println("删除Role失败！")
		return false
	}
	sysService.InvalidateRouteCache()
	return true
}

func CreateRole(newRole CreateRoleStruct) error {
	return serviceInstance().CreateRole(newRole)
}

func (s *Service) CreateRole(newRole CreateRoleStruct) error {
	if err := model.CreateRole(model.Role{
		RoleKey:  newRole.RoleKey,
		RoleName: newRole.RoleName,
		Remark:   newRole.Remark,
	}); err != nil {
		return err
	}
	sysService.InvalidateRouteCache()
	return nil
}

func UpdateRole(roleId int, u UpdateRoleStruct) bool {
	return serviceInstance().UpdateRole(roleId, u)
}

func (s *Service) UpdateRole(roleId int, u UpdateRoleStruct) bool {
	wheres := make(map[string]any)
	wheres["role_id ="] = roleId

	updates := make(map[string]any)
	updates["role_name"] = u.RoleName
	updates["remark"] = u.Remark
	_, rowsAffected := model.Update(&model.Role{}, wheres, updates)
	if rowsAffected == 0 {
		logging.Println("修改Role失败！")
		return false
	}
	sysService.InvalidateRouteCache()
	return true
}

func (u *RoleStruct) getConditionMaps() map[string]any {
	maps := make(map[string]any)
	for k, v := range u.Conditions {
		maps[k] = v
	}
	if _, ok := maps["deleted_at is"]; !ok {
		maps["deleted_at is"] = nil
	}
	return maps
}

func (u *RoleStruct) Count() (int64, error) {
	count, err := model.GetTotal(model.Role{}, u.getConditionMaps())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *RoleStruct) GetAll() ([]*model.Role, error) {
	roles, err := model.GetRoles(u.Pagination, u.getConditionMaps())
	if err != nil {
		return nil, err
	}

	return roles, nil
}
