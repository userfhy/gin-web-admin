package deptController

import (
	deptService "gin-web-admin/app/service/v1/dept"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/com"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *deptService.Service
}

func NewHandler(service *deptService.Service) *Handler {
	if service == nil {
		panic("dept handler requires non-nil service")
	}
	return &Handler{service: service}
}

// GET /dept?tree=1
func (h *Handler) GetDeptList(c *gin.Context) {
	appG := common.Gin{C: c}

	tree := c.DefaultQuery("tree", "")
	if tree == "1" {
		data, err := h.service.GetDeptTree()
		if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "获取部门树失败", err) {
			return
		}
		appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
		return
	}

	data, err := h.service.GetDeptList()
	if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "获取部门列表失败", err) {
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}

// POST /dept
func (h *Handler) CreateDept(c *gin.Context) {
	appG := common.Gin{C: c}
	var payload deptService.CreateDeptStruct
	if err := c.ShouldBindJSON(&payload); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	err := h.service.CreateDept(payload)
	if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "创建部门失败", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}

// PUT /dept/:id
func (h *Handler) UpdateDept(c *gin.Context) {
	appG := common.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	var payload deptService.UpdateDeptStruct
	if err := c.ShouldBindJSON(&payload); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	err := h.service.UpdateDept(id, payload)
	if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "更新部门失败", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}

// DELETE /dept/:id
func (h *Handler) DeleteDept(c *gin.Context) {
	appG := common.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	err := h.service.DeleteDept(id)
	if err != nil {
		// 子部门存在：禁止删除
		if err.Error() == "has child dept" {
			appG.Response(http.StatusOK, code.ERROR, "存在子部门，禁止删除", nil)
			return
		}
		if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "删除部门失败", err) {
			return
		}
	}

	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}
