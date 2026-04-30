package sysService

import (
	"fmt"
	"sort"
	"strings"

	model "gin-web-admin/app/models"
)

type BackendRoute struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type PermissionVisualizationVO struct {
	Roles         []PermissionRole      `json:"roles"`
	Menus         []PermissionMenu      `json:"menus"`
	APIRules      []PermissionAPIRule   `json:"apiRules"`
	BackendRoutes []BackendRoute        `json:"backendRoutes"`
	Scans         PermissionScanSummary `json:"scans"`
}

type PermissionRole struct {
	ID         uint                `json:"id"`
	Name       string              `json:"name"`
	Key        string              `json:"key"`
	IsAdmin    bool                `json:"isAdmin"`
	Status     int                 `json:"status"`
	MenuIDs    []int               `json:"menuIds"`
	MenuCount  int                 `json:"menuCount"`
	APIRules   []PermissionAPIRule `json:"apiRules"`
	APICount   int                 `json:"apiCount"`
	IssueCount int                 `json:"issueCount"`
}

type PermissionMenu struct {
	ID        int      `json:"id"`
	ParentID  int      `json:"parentId"`
	MenuType  int      `json:"menuType"`
	Title     string   `json:"title"`
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Component string   `json:"component"`
	Auths     string   `json:"auths"`
	RoleKeys  []string `json:"roleKeys"`
}

type PermissionAPIRule struct {
	ID          uint   `json:"id"`
	RoleKey     string `json:"roleKey"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ExistsRoute bool   `json:"existsRoute"`
}

type PermissionIssue struct {
	Code    string `json:"code"`
	Level   string `json:"level"`
	Message string `json:"message"`
	RoleKey string `json:"roleKey,omitempty"`
	MenuID  int    `json:"menuId,omitempty"`
	APIID   uint   `json:"apiId,omitempty"`
	Path    string `json:"path,omitempty"`
	Method  string `json:"method,omitempty"`
}

type PermissionScanSummary struct {
	MenuWithoutRole []PermissionIssue `json:"menuWithoutRole"`
	MenuMissingPeer []PermissionIssue `json:"menuMissingPeer"`
	RoleWithoutMenu []PermissionIssue `json:"roleWithoutMenu"`
	RoleWithoutAPI  []PermissionIssue `json:"roleWithoutApi"`
	CasbinNoRole    []PermissionIssue `json:"casbinNoRole"`
	CasbinNoRoute   []PermissionIssue `json:"casbinNoRoute"`
	RoleMenuOrphan  []PermissionIssue `json:"roleMenuOrphan"`
}

func (s *Service) GetPermissionVisualization(routes []BackendRoute) (*PermissionVisualizationVO, error) {
	roles, err := s.listPermissionRoles()
	if err != nil {
		return nil, err
	}
	menus, err := s.listPermissionMenus()
	if err != nil {
		return nil, err
	}
	roleMenus, err := s.listRoleMenus()
	if err != nil {
		return nil, err
	}
	apiRules, err := s.listPermissionAPIRules()
	if err != nil {
		return nil, err
	}

	return buildPermissionVisualization(roles, menus, roleMenus, apiRules, routes), nil
}

func (s *Service) listPermissionRoles() ([]model.Role, error) {
	var roles []model.Role
	err := s.db().Order("role_sort ASC, role_id ASC").Find(&roles).Error
	return roles, err
}

func (s *Service) listPermissionMenus() ([]model.Menu, error) {
	var menus []model.Menu
	err := s.db().Order("`rank` ASC, id ASC").Find(&menus).Error
	return menus, err
}

func (s *Service) listRoleMenus() ([]model.RoleMenu, error) {
	var rows []model.RoleMenu
	err := s.db().Find(&rows).Error
	return rows, err
}

func (s *Service) listPermissionAPIRules() ([]model.CasbinRuleM, error) {
	var rows []model.CasbinRuleM
	err := s.db().Where("ptype = ? OR ptype = ''", "p").Order("id ASC").Find(&rows).Error
	return rows, err
}

func buildPermissionVisualization(roles []model.Role, menus []model.Menu, roleMenus []model.RoleMenu, apiRules []model.CasbinRuleM, routes []BackendRoute) *PermissionVisualizationVO {
	roleByID := make(map[uint]model.Role, len(roles))
	roleByKey := make(map[string]model.Role, len(roles))
	for _, role := range roles {
		roleByID[role.RoleId] = role
		roleByKey[role.RoleKey] = role
	}

	menuByID := make(map[int]model.Menu, len(menus))
	menuRoleKeys := make(map[int][]string)
	roleMenuIDs := make(map[uint][]int)
	for _, menu := range menus {
		menuByID[menu.ID] = menu
	}
	for _, row := range roleMenus {
		roleMenuIDs[row.RoleId] = append(roleMenuIDs[row.RoleId], row.MenuId)
		if role, ok := roleByID[row.RoleId]; ok {
			menuRoleKeys[row.MenuId] = appendUniqueString(menuRoleKeys[row.MenuId], role.RoleKey)
		}
	}

	routeSet := make(map[string]struct{}, len(routes))
	for _, route := range routes {
		routeSet[routeKey(route.Method, route.Path)] = struct{}{}
	}

	permissionAPIRules := make([]PermissionAPIRule, 0, len(apiRules))
	roleAPIRules := make(map[string][]PermissionAPIRule)
	for _, rule := range apiRules {
		item := PermissionAPIRule{
			ID:          rule.ID,
			RoleKey:     rule.V0,
			Path:        rule.V1,
			Method:      strings.ToUpper(rule.V2),
			ExistsRoute: hasRoute(routeSet, rule.V2, rule.V1),
		}
		permissionAPIRules = append(permissionAPIRules, item)
		roleAPIRules[item.RoleKey] = append(roleAPIRules[item.RoleKey], item)
	}

	permissionMenus := make([]PermissionMenu, 0, len(menus))
	for _, menu := range menus {
		roleKeys := append(make([]string, 0, len(menuRoleKeys[menu.ID])), menuRoleKeys[menu.ID]...)
		sort.Strings(roleKeys)
		permissionMenus = append(permissionMenus, PermissionMenu{
			ID:        menu.ID,
			ParentID:  menu.ParentID,
			MenuType:  menu.MenuType,
			Title:     menu.Title,
			Name:      menu.Name,
			Path:      menu.Path,
			Component: menu.Component,
			Auths:     menu.Auths,
			RoleKeys:  roleKeys,
		})
	}

	scans := buildPermissionScans(roles, menus, roleMenus, roleByID, roleByKey, menuByID, roleMenuIDs, permissionAPIRules)
	permissionRoles := make([]PermissionRole, 0, len(roles))
	for _, role := range roles {
		menuIDs := append(make([]int, 0, len(roleMenuIDs[role.RoleId])), roleMenuIDs[role.RoleId]...)
		sort.Ints(menuIDs)
		apis := append(make([]PermissionAPIRule, 0, len(roleAPIRules[role.RoleKey])), roleAPIRules[role.RoleKey]...)
		permissionRoles = append(permissionRoles, PermissionRole{
			ID:         role.RoleId,
			Name:       role.RoleName,
			Key:        role.RoleKey,
			IsAdmin:    role.IsAdmin,
			Status:     role.Status,
			MenuIDs:    menuIDs,
			MenuCount:  len(menuIDs),
			APIRules:   apis,
			APICount:   len(apis),
			IssueCount: countRoleIssues(role.RoleKey, scans),
		})
	}

	return &PermissionVisualizationVO{
		Roles:         permissionRoles,
		Menus:         permissionMenus,
		APIRules:      permissionAPIRules,
		BackendRoutes: append(make([]BackendRoute, 0, len(routes)), routes...),
		Scans:         scans,
	}
}

func buildPermissionScans(roles []model.Role, menus []model.Menu, roleMenus []model.RoleMenu, roleByID map[uint]model.Role, roleByKey map[string]model.Role, menuByID map[int]model.Menu, roleMenuIDs map[uint][]int, apiRules []PermissionAPIRule) PermissionScanSummary {
	scans := emptyPermissionScanSummary()
	for _, menu := range menus {
		if len(menuRoleIDs(roleMenus, menu.ID)) == 0 {
			scans.MenuWithoutRole = append(scans.MenuWithoutRole, PermissionIssue{
				Code:    "menu_without_role",
				Level:   "warning",
				Message: fmt.Sprintf("菜单未分配给任何角色: %s", fallbackPermissionText(menu.Title, menu.Path)),
				MenuID:  menu.ID,
				Path:    menu.Path,
			})
		}
		if menu.ParentID != 0 {
			if _, ok := menuByID[menu.ParentID]; !ok {
				scans.MenuMissingPeer = append(scans.MenuMissingPeer, PermissionIssue{
					Code:    "menu_missing_parent",
					Level:   "error",
					Message: fmt.Sprintf("菜单父节点不存在: %s", fallbackPermissionText(menu.Title, menu.Path)),
					MenuID:  menu.ID,
					Path:    menu.Path,
				})
			}
		}
	}

	for _, role := range roles {
		if len(roleMenuIDs[role.RoleId]) == 0 {
			scans.RoleWithoutMenu = append(scans.RoleWithoutMenu, PermissionIssue{
				Code:    "role_without_menu",
				Level:   "warning",
				Message: fmt.Sprintf("角色未配置菜单权限: %s", role.RoleKey),
				RoleKey: role.RoleKey,
			})
		}
	}

	apiCountByRole := make(map[string]int)
	for _, rule := range apiRules {
		apiCountByRole[rule.RoleKey]++
		if _, ok := roleByKey[rule.RoleKey]; !ok {
			scans.CasbinNoRole = append(scans.CasbinNoRole, PermissionIssue{
				Code:    "casbin_no_role",
				Level:   "error",
				Message: fmt.Sprintf("API 权限引用了不存在的角色: %s", rule.RoleKey),
				RoleKey: rule.RoleKey,
				APIID:   rule.ID,
				Path:    rule.Path,
				Method:  rule.Method,
			})
		}
		if !rule.ExistsRoute {
			scans.CasbinNoRoute = append(scans.CasbinNoRoute, PermissionIssue{
				Code:    "casbin_no_route",
				Level:   "warning",
				Message: fmt.Sprintf("API 权限对应的后端路由不存在: %s %s", rule.Method, rule.Path),
				RoleKey: rule.RoleKey,
				APIID:   rule.ID,
				Path:    rule.Path,
				Method:  rule.Method,
			})
		}
	}
	for _, role := range roles {
		if apiCountByRole[role.RoleKey] == 0 {
			scans.RoleWithoutAPI = append(scans.RoleWithoutAPI, PermissionIssue{
				Code:    "role_without_api",
				Level:   "warning",
				Message: fmt.Sprintf("角色未配置 API 权限: %s", role.RoleKey),
				RoleKey: role.RoleKey,
			})
		}
	}

	for _, row := range roleMenus {
		if _, ok := roleByID[row.RoleId]; !ok {
			scans.RoleMenuOrphan = append(scans.RoleMenuOrphan, PermissionIssue{
				Code:    "role_menu_missing_role",
				Level:   "error",
				Message: fmt.Sprintf("角色菜单关联引用了不存在的角色: %d", row.RoleId),
				MenuID:  row.MenuId,
			})
		}
		if _, ok := menuByID[row.MenuId]; !ok {
			roleKey := ""
			if role, ok := roleByID[row.RoleId]; ok {
				roleKey = role.RoleKey
			}
			scans.RoleMenuOrphan = append(scans.RoleMenuOrphan, PermissionIssue{
				Code:    "role_menu_missing_menu",
				Level:   "error",
				Message: fmt.Sprintf("角色菜单关联引用了不存在的菜单: %d", row.MenuId),
				RoleKey: roleKey,
				MenuID:  row.MenuId,
			})
		}
	}
	return scans
}

func emptyPermissionScanSummary() PermissionScanSummary {
	return PermissionScanSummary{
		MenuWithoutRole: make([]PermissionIssue, 0),
		MenuMissingPeer: make([]PermissionIssue, 0),
		RoleWithoutMenu: make([]PermissionIssue, 0),
		RoleWithoutAPI:  make([]PermissionIssue, 0),
		CasbinNoRole:    make([]PermissionIssue, 0),
		CasbinNoRoute:   make([]PermissionIssue, 0),
		RoleMenuOrphan:  make([]PermissionIssue, 0),
	}
}

func countRoleIssues(roleKey string, scans PermissionScanSummary) int {
	count := 0
	groups := [][]PermissionIssue{
		scans.RoleWithoutMenu,
		scans.RoleWithoutAPI,
		scans.CasbinNoRole,
		scans.CasbinNoRoute,
		scans.RoleMenuOrphan,
	}
	for _, group := range groups {
		for _, issue := range group {
			if issue.RoleKey == roleKey {
				count++
			}
		}
	}
	return count
}

func routeKey(method, path string) string {
	return strings.ToUpper(strings.TrimSpace(method)) + " " + strings.TrimSpace(path)
}

func hasRoute(routeSet map[string]struct{}, method, path string) bool {
	_, ok := routeSet[routeKey(method, path)]
	return ok
}

func appendUniqueString(list []string, value string) []string {
	value = strings.TrimSpace(value)
	if value == "" {
		return list
	}
	for _, item := range list {
		if item == value {
			return list
		}
	}
	return append(list, value)
}

func menuRoleIDs(rows []model.RoleMenu, menuID int) []uint {
	ids := make([]uint, 0)
	for _, row := range rows {
		if row.MenuId == menuID {
			ids = append(ids, row.RoleId)
		}
	}
	return ids
}

func fallbackPermissionText(primary, secondary string) string {
	if strings.TrimSpace(primary) != "" {
		return primary
	}
	return secondary
}
