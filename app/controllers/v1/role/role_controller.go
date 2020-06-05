package roleController

import (
    roleService "gin-test/app/service/v1/role"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/code"
    "gin-test/utils/com"
    "github.com/gin-gonic/gin"
    "net/http"
)

// @Summary 删除角色
// @Description 删除角色
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Role
// @Param role_id path int true "role_id"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /role/{role_id} [delete]
func DeleteRole(c *gin.Context) {
    appG := common.Gin{C: c}
    roleId, err := com.StrTo(c.Param("role_id")).Uint()
    if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
        return
    }

    deleteSuccessful := roleService.DeleteRole(roleId)
    if !deleteSuccessful {
        appG.Response(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}

// @Summary 添加角色
// @Description 添加角色
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Role
// @Param role_id path int true "role_id"
// @Param payload body roleService.CreateRoleStruct true "添加"、
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /role [post]
func CreateRole(c *gin.Context) {
    appG := common.Gin{C: c}

    var createRole roleService.CreateRoleStruct
    err := c.ShouldBindJSON(&createRole)

    if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
        return
    }

    err = roleService.CreateRole(createRole)
    if utils.HandleError(c, http.StatusInternalServerError, http.StatusInternalServerError, "添加新角色失败！", err) {
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "ok", createRole)
}

// @Summary 修改角色
// @Description 修改角色信息
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Role
// @Param role_id path int true "role_id"
// @Param payload body roleService.UpdateRoleStruct true "修改角色"、
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /role/{role_id} [put]
func UpdateRole(c *gin.Context) {
    appG := common.Gin{C: c}
    roleId := com.StrTo(c.Param("role_id")).MustInt()

    var updateRole roleService.UpdateRoleStruct
    err := c.ShouldBindJSON(&updateRole)

    if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
        return
    }

    changeSuccessful := roleService.UpdateRole(roleId, updateRole)
    if !changeSuccessful {
        appG.Response(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
        return
    }


    appG.Response(http.StatusOK, code.SUCCESS, "ok", updateRole)
}

// @Summary 角色列表
// @Description 获取角色表
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Role
// @Param p query int true "page number"
// @Param n query int true "page limit"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /role [get]
func GetRoles(c *gin.Context) {
    appG := common.Gin{C: c}
    err, errStr, p, n := utils.GetPage(c)
    if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, errStr, err) {
        return
    }

    roleServiceObj := roleService.RoleStruct{
        PageNum: p,
        PageSize: n,
    }

    total, err := roleServiceObj.Count()
    if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "获取页数失败", err) {
        return
    }

    userArr, err := roleServiceObj.GetAll()
    if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "服务器错误", err) {
        return
    }

    data := utils.PageResult{
        List: userArr,
        Total: total,
        PageSize: n,
    }
    appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}