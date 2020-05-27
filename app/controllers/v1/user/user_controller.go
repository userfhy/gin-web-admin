package userController

import (
    userService "gin-test/app/service/v1/user"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "net/http"
)

// @Summary 用户列表
// @Description 获取用户列表
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param p query int true "page number"
// @Param n query int true "page limit"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /user [get]
func GetUsers(c *gin.Context) {
    appG := common.Gin{C: c}

    _, errStr, p, n := utils.GetPage(c)
    if errStr != "" {
        appG.Response(http.StatusBadRequest, code.InvalidParams, errStr, nil)
        return
    }

    userServiceObj := userService.UserStruct{
        PageNum: p,
        PageSize: n,
    }

    total, err := userServiceObj.Count()
    if err != nil {
        appG.Response(http.StatusInternalServerError, code.ERROR, "获取页数失败" + err.Error(), nil)
        return
    }

    userArr, err := userServiceObj.GetAll()
    if err != nil {
        appG.Response(http.StatusInternalServerError, code.ERROR, "服务器错误" + err.Error(), nil)
        return
    }

    data := utils.PageResult{
        List: userArr,
        Total: total,
        PageSize: n,
    }
    appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}