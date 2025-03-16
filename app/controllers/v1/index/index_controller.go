package indexController

import (
	"encoding/json"
	"errors"
	"gin-web-admin/common"
	"gin-web-admin/common/sse"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/logging"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志系统
	logging.Setup("sse-service", &logging.Option{
		LogLevel:   "debug",
		Formatter:  "text",
		OutputPath: "",
	})
}

// 初始化SSE服务
var SSEService = sse.NewSSE(
	sse.WithConfig(sse.Config{
		HeartbeatInterval: 10 * time.Second,
		ChannelBufferSize: 50,
		WriteTimeout:      30 * time.Second,
		MaxClientMessages: 10000,
	}),
	sse.WithSendFailHandler(func(clientID string, msg sse.Message) {
		logging.Printf("重要消息发送失败: client=%s event=%s", clientID, msg.Event)
	}),
	sse.WithLogger(&sse.LogrusAdapter{}),
)

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

// 请求/响应结构体
type SendRequest struct {
	ClientID string `json:"clientId" binding:"omitempty,uuid4"` // 使用uuid4格式校验
	Event    string `json:"event" binding:"required,alphanum"`  // 必须字母数字组合
	Data     any    `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// @Summary 发送消息到指定客户端
// @Description 支持点对点消息和广播消息（clientId留空时广播）
// @Accept  json
// @Produce json
// @Tags    Test
// @Param   message body SendRequest true "消息内容"
// @Success 204 "消息已接受"
// @Failure 400 {object} ErrorResponse "请求格式错误"
// @Failure 404 {object} ErrorResponse "客户端不存在"
// @Failure 503 {object} ErrorResponse "服务不可用"
// @Router  /test/send [post]
func SendStream(c *gin.Context) {
	var req SendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Warn("无效的请求格式: ", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    4001,
			Message: "请求格式不符合规范",
		})
		return
	}

	// 参数校验
	if req.Event == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    4002,
			Message: "事件类型不能为空",
		})
		return
	}

	// logging.Info("发送消息请求",
	// 	zap.String("client_id", req.ClientID),
	// 	zap.String("event", req.Event),
	// 	zap.Any("data", req.Data))

	// 处理广播逻辑
	if req.ClientID == "" {
		SSEService.Broadcast(sse.Message{
			Event: req.Event,
			Data:  req.Data,
		})
		c.Status(http.StatusNoContent)
		return
	}

	// 点对点发送
	if err := SSEService.Send(req.ClientID, sse.Message{
		Event: req.Event,
		Data:  req.Data,
	}); err != nil {
		if errors.Is(err, sse.ErrClientNotFound) {
			logging.Warn("目标客户端不存在", zap.String("client_id", req.ClientID))
			c.JSON(http.StatusNotFound, ErrorResponse{
				Code:    4041,
				Message: "目标客户端不存在或已离线",
			})
		} else {
			logging.Error("消息发送失败",
				zap.String("client_id", req.ClientID),
				zap.Error(err))
			c.JSON(http.StatusServiceUnavailable, ErrorResponse{
				Code:    5031,
				Message: "消息服务暂时不可用",
			})
		}
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
