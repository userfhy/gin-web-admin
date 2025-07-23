package menuController

import (
	"net/http"
	"strconv"

	model "gin-web-admin/app/models"
	service "gin-web-admin/app/service/v1/menu"
	"gin-web-admin/common"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
)

// 新增菜单
func CreateMenu(c *gin.Context) {
	appG := common.Gin{C: c}

	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	if err := service.CreateMenu(menu); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "创建成功", nil)
}

// 更新菜单
func UpdateMenu(c *gin.Context) {
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

	if err := service.UpdateMenu(id, data); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "更新失败: "+err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "更新成功", nil)
}

// 删除菜单
func DeleteMenu(c *gin.Context) {
	appG := common.Gin{C: c}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效ID", nil)
		return
	}

	if err := service.DeleteMenu(id); err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}

// 获取菜单列表（分页）
func GetMenuList(c *gin.Context) {
	appG := common.Gin{C: c}

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 构建查询条件，可根据你前端传参处理
	where := make(map[string]any)
	if menuType := c.Query("menu_type"); menuType != "" {
		where["menu_type"] = menuType
	}

	menus, err := service.GetMenuList(pageNum, pageSize, where)
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
func GetAllMenus(c *gin.Context) {
	appG := common.Gin{C: c}

	where := map[string]any{}
	menus, err := service.GetAllMenus(where)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取失败", nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "获取成功", menus)
}
