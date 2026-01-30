package roleController

import (
	roleService "gin-web-admin/app/service/v1/role"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/com"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /role/:role_id/menu_ids
func GetRoleMenuIds(c *gin.Context) {
	appG := common.Gin{C: c}
	roleId, err := com.StrTo(c.Param("role_id")).Uint()
	if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	ids, err := roleService.GetRoleMenuIds(roleId)
	if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "获取角色菜单权限失败", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", ids)
}

// PUT /role/:role_id/menu_ids
func SaveRoleMenuIds(c *gin.Context) {
	appG := common.Gin{C: c}
	roleId, err := com.StrTo(c.Param("role_id")).Uint()
	if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	var payload roleService.SaveRoleMenuIdsReq
	if err := c.ShouldBindJSON(&payload); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	if err := roleService.SaveRoleMenuIds(roleId, payload.MenuIds); utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "保存角色菜单权限失败", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}
