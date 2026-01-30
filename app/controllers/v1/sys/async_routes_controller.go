package sysController

import (
	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// GetAsyncRoutes returns async routes generated from gin_menu.
// It filters menus by current user's role via gin_role_menu.
// GET /get-async-routes
func GetAsyncRoutes(c *gin.Context) {
	// claims set by JWT middleware
	claimsAny, _ := c.Get("claims")
	claims, _ := claimsAny.(*utils.Claims)
	roleKey := ""
	if claims != nil {
		roleKey = claims.RoleKey
	}

	menus, err := GetMenusForRole(roleKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    []any{},
		})
		return
	}

	// Build tree by parentId
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

	roots := childrenMap[0]

	// IMPORTANT: front-end expects code === 0
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    roots,
	})
}

// ----- Internal menu shape for route build -----

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

// GetMenusForRole loads menus from gin_menu filtered by role_key -> gin_role_menu.
func GetMenusForRole(roleKey string) ([]menuRow, error) {
	// if no roleKey, return empty
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

	return getMenusFromDBByRole(role.RoleId)
}
