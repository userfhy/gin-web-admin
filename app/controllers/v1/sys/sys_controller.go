package sysController

import (
	"gin-web-admin/common"
	"gin-web-admin/utils/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Routers gin.RoutesInfo

// @Summary		菜单列表
// @Description	get router list
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			SYS
// @Success		200	{object}	common.Response
// @Router			/sys/menu_list [get]
func GetMenuList(c *gin.Context) {

}

// @Summary		后端存在路由列表
// @Description	get router list
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			SYS
// @Success		200	{object}	common.Response
// @Router			/sys/router [get]
func GetRouterList(c *gin.Context) {
	appG := common.Gin{C: c}

	type Router struct {
		Path   string `json:"path"`
		Method string `json:"method"`
	}

	data := make([]Router, 0)

	routers := Routers

	for index := range routers {
		var router Router
		router.Method = Routers[index].Method
		router.Path = Routers[index].Path
		data = append(data, router)
	}

	appG.Response(http.StatusOK, code.SUCCESS, "获取存在路由列表成功", data)
}
