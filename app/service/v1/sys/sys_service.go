package sysService

import (
	model "gin-web-admin/app/models"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	store *data.Store
}

const (
	asyncRoutesCachePrefix = "sys:routes:"
	asyncRoutesCacheTTL    = 10 * time.Minute
)

var defaultService *Service

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func SetDefaultService(s *Service) {
	defaultService = s
}

func serviceInstance() *Service {
	if defaultService == nil {
		panic("sys service not initialized")
	}
	return defaultService
}

type AsyncRouteMeta struct {
	Title      string `json:"title"`
	Icon       string `json:"icon,omitempty"`
	ShowLink   bool   `json:"showLink"`
	Rank       int    `json:"rank,omitempty"`
	KeepAlive  bool   `json:"keepAlive,omitempty"`
	FrameSrc   string `json:"frameSrc,omitempty"`
	ActivePath string `json:"activePath,omitempty"`
}

type AsyncRoute struct {
	Path      string         `json:"path"`
	Name      string         `json:"name,omitempty"`
	Component string         `json:"component,omitempty"`
	Redirect  string         `json:"redirect,omitempty"`
	Meta      AsyncRouteMeta `json:"meta"`
	Children  []*AsyncRoute  `json:"children,omitempty"`
}

func GetAsyncRoutes(roleKey string) ([]*AsyncRoute, error) {
	return serviceInstance().GetAsyncRoutes(roleKey)
}

func (s *Service) GetAsyncRoutes(roleKey string) ([]*AsyncRoute, error) {
	if roleKey == "" {
		return []*AsyncRoute{}, nil
	}

	if routes, ok, err := s.getCachedRoutes(roleKey); err == nil && ok {
		return routes, nil
	} else if err != nil {
		logging.Warnf("get async routes cache failed: %v", err)
	}

	menus, err := s.getMenusForRole(roleKey)
	if err != nil {
		return nil, err
	}

	nodeMap := make(map[int]*AsyncRoute)
	childrenMap := make(map[int][]*AsyncRoute)

	for _, m := range menus {
		ar := &AsyncRoute{
			Path:      m.Path,
			Name:      m.Name,
			Component: m.Component,
			Redirect:  m.Redirect,
			Meta: AsyncRouteMeta{
				Title:      m.Title,
				Icon:       m.Icon,
				ShowLink:   m.ShowLink,
				Rank:       m.Rank,
				KeepAlive:  m.KeepAlive,
				FrameSrc:   m.FrameSrc,
				ActivePath: m.ActivePath,
			},
		}
		nodeMap[m.ID] = ar
		childrenMap[m.ParentID] = append(childrenMap[m.ParentID], ar)
	}

	for id, node := range nodeMap {
		node.Children = childrenMap[id]
	}

	routes := childrenMap[0]

	s.setCachedRoutes(roleKey, routes)

	return routes, nil
}

// ----- internal helpers -----

type menuRow struct {
	ID         int
	ParentID   int
	Path       string
	Name       string
	Component  string
	Redirect   string
	Title      string
	Icon       string
	Rank       int
	ShowLink   bool
	KeepAlive  bool
	FrameSrc   string
	ActivePath string
}

func (s *Service) getMenusForRole(roleKey string) ([]menuRow, error) {
	if roleKey == "" {
		return []menuRow{}, nil
	}

	role, err := model.GetRoleByKey(roleKey)
	if err != nil {
		return nil, err
	}
	if role == nil || role.RoleId == 0 {
		return []menuRow{}, nil
	}

	return s.getMenusFromDBByRole(role.RoleId)
}

func (s *Service) getMenusFromDBByRole(roleID uint) ([]menuRow, error) {
	var menuIDs []int
	db := s.db()
	if err := db.Model(&model.RoleMenu{}).
		Where("role_id = ?", roleID).
		Pluck("menu_id", &menuIDs).Error; err != nil {
		return nil, err
	}
	if len(menuIDs) == 0 {
		return []menuRow{}, nil
	}

	allMenus, err := model.GetAllMenus(map[string]any{})
	if err != nil {
		return nil, err
	}

	byID := make(map[int]*model.Menu, len(allMenus))
	for _, m := range allMenus {
		byID[m.ID] = m
	}

	selected := make(map[int]struct{})
	for _, id := range menuIDs {
		selected[id] = struct{}{}
		cur := byID[id]
		for cur != nil && cur.ParentID != 0 {
			selected[cur.ParentID] = struct{}{}
			cur = byID[cur.ParentID]
		}
	}

	filtered := make([]*model.Menu, 0, len(selected))
	for id := range selected {
		if m := byID[id]; m != nil {
			filtered = append(filtered, m)
		}
	}

	ids := make([]int, 0, len(filtered))
	for _, m := range filtered {
		ids = append(ids, m.ID)
	}

	var sorted []*model.Menu
	if err := db.Where("id IN ?", ids).Order("`rank` ASC").Find(&sorted).Error; err != nil && err != gorm.ErrRecordNotFound {
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
			KeepAlive:  m.KeepAlive,
		}
		if m.Rank != nil {
			row.Rank = *m.Rank
		}
		rows = append(rows, row)
	}
	return rows
}

func (s *Service) getCachedRoutes(roleKey string) ([]*AsyncRoute, bool, error) {
	var routes []*AsyncRoute
	found, err := gredis.GetJSON(asyncRoutesCacheKey(roleKey), &routes)
	return routes, found, err
}

func (s *Service) setCachedRoutes(roleKey string, routes []*AsyncRoute) {
	gredis.SetJSONAsync(asyncRoutesCacheKey(roleKey), routes, asyncRoutesCacheTTL)
}

func asyncRoutesCacheKey(roleKey string) string {
	return asyncRoutesCachePrefix + roleKey
}

// InvalidateRouteCache 清理菜单路由缓存
func InvalidateRouteCache() {
	gredis.DeleteByPrefixAsync(asyncRoutesCachePrefix)
}

func (s *Service) db() *gorm.DB {
	if s != nil && s.store != nil && s.store.DB() != nil {
		return s.store.DB()
	}
	return model.DB()
}
