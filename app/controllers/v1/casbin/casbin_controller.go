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
// @Param group_by query string false "v0 根据 role key 分组"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /casbin [get]
func GetCasbinList(c *gin.Context) {
    appG := common.Gin{C: c}
    groupBy := c.DefaultQuery("group_by", "")
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

    arr, err := casbinServiceObj.GetAll()
    if utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "服务器错误", err) {
        return
    }

    data := utils.PageResult{
        List: arr,
        Total: total,
        PageSize: n,
    }

    // 筛选 group by
    if groupBy == "v0" {
        groupMap := make(map[string] []interface{})
        for k, v := range arr {
            if arr[k].V0 == v.V0 {
                groupMap[arr[k].V0] = append(groupMap[arr[k].V0], v)
            }
        }
        data.List = groupMap
    }

    appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}