package indexController

import (
    "encoding/json"
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
// @Security ApiKeyAuth
// @Tags Test
// @Success 200 {object} common.Response
// @Router /test/ping [get]
func Ping(c *gin.Context) {
    appG := common.Gin{C: c}
    appG.Response(http.StatusOK, code.SUCCESS, "pong", nil)
}

// @Summary Base64 Decode
// @Produce  json
// @Security ApiKeyAuth
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
        appG.Response(http.StatusOK, code.InvalidParams, "base64解码失败" + err.Error(), nil)
        return
    }

    // 结构体
    var imgTextArray []common.ImgText
    err = json.Unmarshal(arrByte, &imgTextArray)

    if utils.HandleError(c, http.StatusInternalServerError, http.StatusInternalServerError, "参数绑定失败", err) {
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "文字解析成功", imgTextArray)
}
