package menuController

import (
	"net/http"
	"strconv"

	model "gin-web-admin/app/models"
	menuService "gin-web-admin/app/service/v1/menu"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *menuService.Service
}

func NewHandler(service *menuService.Service) *Handler {
	if service == nil {
		panic("menu handler requires non-nil service")
	}
	return &Handler{service: service}
}

// 新增菜单
func (h *Handler) CreateMenu(c *gin.Context) {
	appG := common.Gin{C: c}

	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := h.service.CreateMenu(menu); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

// 更新菜单
func (h *Handler) UpdateMenu(c *gin.Context) {
	appG := common.Gin{C: c}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := h.service.UpdateMenu(id, data); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "更新失败: "+err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

// 删除菜单
func (h *Handler) DeleteMenu(c *gin.Context) {
	appG := common.Gin{C: c}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	if err := h.service.DeleteMenu(id); err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}

// 获取菜单列表（分页）
func (h *Handler) GetMenuList(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	// 构建查询条件，可根据你前端传参处理
	where := make(map[string]any)
	if menuType := c.Query("menu_type"); menuType != "" {
		where["menu_type"] = menuType
	}

	menus, err := h.service.GetMenuList(pg.Clone(), where)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "查询失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "查询成功", menus)
}

// @Summary		菜单列表
// @Description	get menu list
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			SYS
// @Success		200	{object}	common.Response
// @Router			/menu/menu_list [get]
func (h *Handler) GetAllMenus(c *gin.Context) {
	appG := common.Gin{C: c}

	where := map[string]any{}
	menus, err := h.service.GetAllMenus(where)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "获取成功", menus)
}
