package casbinController

import (
    casbinService "gin-test/app/service/v1/casbin"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "net/http"
)

// @Summary 规则列表
// @Description 获取规则列表
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Casbin
// @Param p query int true "page number"
// @Param n query int true "page limit"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /casbin [get]
func GetCasbinList(c *gin.Context) {
    appG := common.Gin{C: c}
    err, errStr, p, n := utils.GetPage(c)
    if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, errStr, err) {
        return
    }

    casbinServiceObj := casbinService.CasbinStruct{
        PageNum: p,
        PageSize: n,
    }

    total, err := casbinServiceObj.Count()
    if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "获取页数失败", err) {
        return
    }

    userArr, err := casbinServiceObj.GetAll()
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