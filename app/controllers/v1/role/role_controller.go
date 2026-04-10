package roleController

import (
	"net/http"

	roleService "gin-web-admin/app/service/v1/role"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/com"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *roleService.Service
}

func NewHandler(service *roleService.Service) *Handler {
	if service == nil {
		panic("role handler requires non-nil service")
	}
	return &Handler{service: service}
}

// @Summary		删除角色
// @Description	删除角色
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			Role
// @Param			role_id	path		int	true	"role_id"
// @Success		200		{object}	common.Response
// @Failure		500		{object}	common.Response
// @Router			/role/{role_id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	appG := common.Gin{C: c}
	roleId, err := com.StrTo(c.Param("role_id")).Uint()
	if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
		return
	}

	deleteSuccessful := h.service.DeleteRole(roleId)
	if !deleteSuccessful {
		appG.Response(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}

// @Summary		添加角色
// @Description	添加角色
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			Role
// @Param			role_id	path		int								true	"role_id"
// @Param			payload	body		roleService.CreateRoleStruct	true	"添加"、
// @Success		200		{object}	common.Response
// @Failure		500		{object}	common.Response
// @Router			/role [post]
func (h *Handler) CreateRole(c *gin.Context) {
	appG := common.Gin{C: c}

	var createRole roleService.CreateRoleStruct
	err := c.ShouldBindJSON(&createRole)

	if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
		return
	}

	err, parameterErrorStr := common.CheckBindStructParameter(createRole, c)
	if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, parameterErrorStr, err) {
		return
	}

	err = h.service.CreateRole(createRole)
	if utils.HandleError(c, http.StatusInternalServerError, http.StatusInternalServerError, "添加新角色失败！", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", createRole)
}

// @Summary		修改角色
// @Description	修改角色信息
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			Role
// @Param			role_id	path		int								true	"role_id"
// @Param			payload	body		roleService.UpdateRoleStruct	true	"修改角色"、
// @Success		200		{object}	common.Response
// @Failure		500		{object}	common.Response
// @Router			/role/{role_id} [put]
func (h *Handler) UpdateRole(c *gin.Context) {
	appG := common.Gin{C: c}
	roleId := com.StrTo(c.Param("role_id")).MustInt()

	var updateRole roleService.UpdateRoleStruct
	err := c.ShouldBindJSON(&updateRole)

	if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
		return
	}

	changeSuccessful := h.service.UpdateRole(roleId, updateRole)
	if !changeSuccessful {
		appG.Response(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", updateRole)
}

// @Summary		角色列表
// @Description	获取角色表
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			Role
// @Param			p	query		int	true	"page number"
// @Param			n	query		int	true	"page limit"
// @Success		200	{object}	common.Response
// @Failure		500	{object}	common.Response
// @Router			/role [get]
func (h *Handler) GetRoles(c *gin.Context) {
	appG := common.Gin{C: c}
	pg, err := utils.GetPagination(c)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	roleServiceObj := roleService.RoleStruct{
		Pagination: pg.Clone(),
	}

	filterBuilder := query.NewBuilder().IsNull("deleted_at")
	if err := filterBuilder.FromQuery(c, query.RuleSet{
		"role_name": {Field: "role_name", Op: query.OpLike},
		"role_key":  {Field: "role_key", Op: query.OpLike},
		"status":    {Field: "status", Op: query.OpEqual, Parser: query.IntParser()},
	}); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	roleServiceObj.Conditions = filterBuilder.Build()

	total, err := roleServiceObj.Count()
	if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "获取页数失败", err) {
		return
	}

	userArr, err := roleServiceObj.GetAll()
	if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "服务器错误", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", pg.Result(userArr, total))
}
