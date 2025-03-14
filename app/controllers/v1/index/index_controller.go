package indexController

import (
	"encoding/json"
	"gin-web-admin/common"
	"gin-web-admin/common/sse"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var SSEService = sse.NewSSE()

// @Summary		Ping
// @Description	Test Ping
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			Test
// @Success		200	{object}	common.Response
// @Router			/test/ping [get]
func Ping(c *gin.Context) {
	appG := common.Gin{C: c}
	appG.Response(http.StatusOK, code.SUCCESS, "pong", nil)
}

// @Summary	Base64 Decode
// @Produce	json
// @Security	ApiKeyAuth
// @Tags		Test
// @Param		base64	query		string	true	"base64 string"
// @Success	200		{object}	common.Response
// @Failure	500		{object}	common.Response
// @Router		/test/font [get]
func Test(c *gin.Context) {
	appG := common.Gin{C: c}

	base64 := c.DefaultQuery("base64", "")

	// 替换字符串
	base64String := strings.Replace(base64, " ", "+", -1)

	// base64 解码
	arrByte, err := utils.Base64Decode(base64String)
	if err != nil {
		appG.Response(http.StatusOK, code.InvalidParams, "base64解码失败"+err.Error(), nil)
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

// @Summary	Test SSE
// @Produce	text/event-stream
// @Tags		Test
// @Router		/test/events [get]
func Stream(c *gin.Context) {

}

// @Summary Send message to specific client
// @Accept  json
// @Produce json
// @Tags    Test
// @Param   message body sse.Message true "Message Content"
// @Router  /send [post]
func SendStream(c *gin.Context) {
	// 启动全局广播
	var req struct {
		ClientID string `json:"clientId"`
		Event    string `json:"event"`
		Data     any    `json:"data"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := SSEService.Send(req.ClientID, sse.Message{
		Event: req.Event,
		Data:  req.Data,
	}); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary	Test SSE Client Count
// @Produce	json
// @Tags		Test
// @Router		/test/count [get]
func SSEClientCount(c *gin.Context) {
	c.JSON(200, gin.H{"count": SSEService.ClientCount()})
}
