package indexController

import (
    "encoding/json"
    userService "gin-test/app/service/v1/user"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

// @Summary Ping
// @Description Test Ping
// @Accept  json
// @Produce  json
// @Tags Test
// @Success 200 {object} common.Response
// @Router /test/ping [get]
func Ping(c *gin.Context) {
    appG := common.Gin{C: c}
    appG.Response(http.StatusOK, code.SUCCESS, "pong", nil)
}

// @Summary Base64 Decode
// @Produce  json
// @Tags Test
// @Param base64 query string true "base64 string"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /test/font [get]
func Test(c *gin.Context) {
    appG := common.Gin{C: c}

    base64 := c.DefaultQuery("base64", "")

    // 替换字符串
    base64String := strings.Replace(base64, " ", "+", -1)

    // base64 解码
    arrByte, err := utils.Base64Decode(base64String)
    if err != nil {
        appG.Response(http.StatusOK, code.INVALID_PARAMS, "base64解码失败" + err.Error(), nil)
        return
    }

    // 结构体
    var imgTextArray []common.ImgText
    err = json.Unmarshal(arrByte, &imgTextArray)

    if utils.HandleError(c, http.StatusInternalServerError, err) {
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "文字解析成功", imgTextArray)
}

// @Summary Get Users
// @Description Test db Connections.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Test
// @Param p query int true "page number"
// @Param n query int true "page limit"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /test/test_users [get]
func GetTestUsers(c *gin.Context) {
    appG := common.Gin{C: c}

    _, errStr, p, n := utils.GetPage(c)
    if errStr != "" {
        appG.Response(http.StatusBadRequest, code.INVALID_PARAMS, errStr, nil)
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

    data := make(map[string]interface{})
    data["lists"] = userArr
    data["total"] = total

    appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}

// @Summary User Login
// @Description Test db Connections.
// @Accept json
// @Produce json
// @Tags Test
// @Param payload body userService.UserLoginStruct true "user login"
// @Success 200 {object} common.Response
// @Router /test/login [post]
func UserLogin(c *gin.Context) {
    appG := common.Gin{C: c}

    // 绑定 query 参数到结构体
    var userLogin userService.UserLoginStruct
    if err := c.ShouldBindJSON(&userLogin); err != nil {
        appG.Response(http.StatusBadRequest, code.INVALID_PARAMS, err.Error(), nil)
        return
    }

    // 验证绑定结构体参数
    err, parameterErrorStr := common.CheckBindStructParameter(userLogin, c)
    if err != nil {
        appG.Response(http.StatusBadRequest, code.INVALID_PARAMS, parameterErrorStr, nil)
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "ok", userLogin)
}
