package casbinController

import (
    casbinService "gin-test/app/service/v1/casbin"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/casbin"
    "gin-test/utils/code"
    "gin-test/utils/com"
    "github.com/gin-gonic/gin"
    "net/http"
)

// @Summary 创建规则
// @Description 创建规则
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Casbin
// @Param payload body casbinService.AddCasbinStruct true "create new user"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /casbin [post]
func CreateCasbin(c *gin.Context) {
    appG := common.Gin{C: c}

    var newCasbin casbinService.AddCasbinStruct
    err := c.ShouldBindJSON(&newCasbin)
    if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
        return
    }

    err, parameterErrorStr := common.CheckBindStructParameter(newCasbin, c)
    if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, parameterErrorStr, err) {
        return
    }

    err = casbinService.CreateCasbin(newCasbin)
    if utils.HandleError(c, http.StatusInternalServerError, http.StatusInternalServerError, "Path添加失败！", err) {
       return
    }

    // 重新生成权限列表
    casbin.SetupCasbin()
    //casbin.SetupCasbin().AddPolicy(newCasbin.V0, newCasbin.V1, newCasbin.V2)

    appG.Response(http.StatusOK, code.SUCCESS, "Path添加成功", nil)
}

// @Summary 修改规则
// @Description 修改规则信息
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Casbin
// @Param id path int true "casbin_id"
// @Param payload body casbinService.AddCasbinStruct true "修改规则"、
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /casbin/{id} [put]
func UpdateCasbin(c *gin.Context) {
    appG := common.Gin{C: c}
    id := com.StrTo(c.Param("id")).MustInt()

    var update casbinService.AddCasbinStruct
    err := c.ShouldBindJSON(&update)

    if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
        return
    }

    changeSuccessful := casbinService.UpdateCasbin(id, update)
    if !changeSuccessful {
        appG.Response(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
        return
    }

    // 重新生成权限列表
    casbin.SetupCasbin()

    appG.Response(http.StatusOK, code.SUCCESS, "ok", update)
}

// @Summary 删除规则
// @Description 删除规则信息
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Casbin
// @Param id path int true "casbin_id"
// @Param payload body casbinService.AddCasbinStruct true "删除规则"、
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /casbin/{id} [delete]
func DeleteCasbin(c *gin.Context) {
    appG := common.Gin{C: c}
    //id := com.StrTo(c.Param("id")).MustInt()

    var update casbinService.AddCasbinStruct
    err := c.ShouldBindJSON(&update)

    if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
        return
    }

    // 重新生成权限列表
    casbin.SetupCasbin().RemovePolicy(update.V0, update.V1, update.V2)

    appG.Response(http.StatusOK, code.SUCCESS, "ok", update)
}

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