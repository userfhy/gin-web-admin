package sysController

import (
	"net/http"

	sysService "gin-web-admin/app/service/v1/sys"
	"gin-web-admin/common"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *sysService.Service
}

func NewHandler(service *sysService.Service) *Handler {
	if service == nil {
		panic("sys handler requires non-nil service")
	}
	return &Handler{service: service}
}

var Routers gin.RoutesInfo

// @Summary		菜单列表
// @Description	get router list
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			SYS
// @Success		200	{object}	common.Response
// @Router			/sys/menu_list [get]
func (h *Handler) GetMenuList(c *gin.Context) {}

// @Summary		后端存在路由列表
// @Description	get router list
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			SYS
// @Success		200	{object}	common.Response
// @Router			/sys/router [get]
func (h *Handler) GetRouterList(c *gin.Context) {
	appG := common.Gin{C: c}

	type Router struct {
		Path   string `json:"path"`
		Method string `json:"method"`
	}

	data := make([]Router, 0, len(Routers))

	for _, route := range Routers {
		data = append(data, Router{
			Method: route.Method,
			Path:   route.Path,
		})
	}

	appG.Response(http.StatusOK, code.SUCCESS, "获取存在路由列表成功", data)
}
