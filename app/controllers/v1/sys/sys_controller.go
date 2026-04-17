package sysController

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	sysService "gin-web-admin/app/service/v1/sys"
	"gin-web-admin/common"
	"gin-web-admin/utils"
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

func (h *Handler) GetOnlineUsers(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(200))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	result, err := h.service.GetOnlineUserList(sysService.OnlineUserQuery{
		Pagination: pg.Clone(),
		Username:   c.Query("username"),
		IP:         c.Query("ip"),
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取在线用户失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) ForceOffline(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效用户ID", nil)
		return
	}
	if err := h.service.ForceOffline(uint(id)); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "强制下线成功", nil)
}

func (h *Handler) GetServerMonitor(c *gin.Context) {
	appG := common.Gin{C: c}

	result, err := h.service.GetServerMonitor()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取服务器监控失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) StreamServerMonitor(c *gin.Context) {
	const (
		streamInterval = 5 * time.Second
		serverRetry    = 3 * time.Second
	)

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	writeEvent := func(event string, seq uint64, payload any) bool {
		body, err := json.Marshal(payload)
		if err != nil {
			return false
		}
		_, _ = c.Writer.Write([]byte("id: " + strconv.FormatUint(seq, 10) + "\n"))
		_, _ = c.Writer.Write([]byte("retry: " + strconv.FormatInt(serverRetry.Milliseconds(), 10) + "\n"))
		_, _ = c.Writer.Write([]byte("event: " + event + "\n"))
		_, _ = c.Writer.Write([]byte("data: " + string(body) + "\n\n"))
		flusher.Flush()
		return true
	}

	sendInit := func() bool {
		result, err := h.service.GetServerMonitorStreamInit(serverRetry)
		if err != nil {
			return writeEvent("error", 0, gin.H{
				"code": code.ERROR,
				"msg":  "获取服务器监控失败",
				"meta": gin.H{
					"serverRetryMs": serverRetry.Milliseconds(),
				},
			})
		}
		return writeEvent("server_monitor_init", result.Seq, gin.H{
			"code": code.SUCCESS,
			"msg":  "ok",
			"data": result,
		})
	}

	sendAppend := func() bool {
		result, err := h.service.GetServerMonitorStreamDelta(serverRetry)
		if err != nil {
			return writeEvent("error", 0, gin.H{
				"code": code.ERROR,
				"msg":  "获取服务器监控失败",
				"meta": gin.H{
					"serverRetryMs": serverRetry.Milliseconds(),
				},
			})
		}
		return writeEvent("server_monitor_append", result.Seq, gin.H{
			"code": code.SUCCESS,
			"msg":  "ok",
			"data": result,
		})
	}

	if !sendInit() {
		return
	}

	ticker := time.NewTicker(streamInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.Request.Context().Done():
			return
		case <-ticker.C:
			if !sendAppend() {
				return
			}
		}
	}
}

func (h *Handler) GetLoginLogs(c *gin.Context) {
	h.getAuditLogs(c, "login")
}

func (h *Handler) GetOperationLogs(c *gin.Context) {
	h.getAuditLogs(c, "operation")
}

func (h *Handler) GetSystemLogs(c *gin.Context) {
	appG := common.Gin{C: c}

	query, ok := h.buildAuditLogQuery(c, "")
	if !ok {
		return
	}

	result, err := h.service.GetSystemLogList(query)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取系统日志失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) GetSystemLogDetail(c *gin.Context) {
	appG := common.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "无效日志ID", nil)
		return
	}

	result, err := h.service.GetSystemLogDetail(uint(id))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) DeleteLoginLogs(c *gin.Context) {
	h.deleteAuditLogs(c, "login")
}

func (h *Handler) DeleteOperationLogs(c *gin.Context) {
	h.deleteAuditLogs(c, "operation")
}

func (h *Handler) DeleteSystemLogs(c *gin.Context) {
	h.deleteAuditLogs(c, "")
}

func (h *Handler) getAuditLogs(c *gin.Context, category string) {
	appG := common.Gin{C: c}

	query, ok := h.buildAuditLogQuery(c, category)
	if !ok {
		return
	}

	var (
		result utils.PageResult
		err    error
	)
	switch category {
	case "login":
		result, err = h.service.GetLoginLogList(query)
	case "operation":
		result, err = h.service.GetOperationLogList(query)
	default:
		result, err = h.service.GetSystemLogList(query)
	}
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取日志列表失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", result)
}

func (h *Handler) buildAuditLogQuery(c *gin.Context, category string) (sysService.AuditLogQuery, bool) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c, utils.WithMaxPageSize(200))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return sysService.AuditLogQuery{}, false
	}
	status, err := sysService.ParseLogStatus(c.Query("status"))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return sysService.AuditLogQuery{}, false
	}
	timeRange := c.QueryArray("timeRange")
	if len(timeRange) == 0 {
		startRaw := c.Query("startTime")
		endRaw := c.Query("endTime")
		if startRaw != "" || endRaw != "" {
			timeRange = []string{startRaw, endRaw}
		}
	}
	start, end, err := sysService.ParseDateRange(timeRange)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return sysService.AuditLogQuery{}, false
	}

	return sysService.AuditLogQuery{
		Pagination: pg.Clone(),
		Category:   category,
		Username:   c.Query("username"),
		Module:     c.Query("module"),
		Status:     status,
		StartTime:  start,
		EndTime:    end,
	}, true
}

func (h *Handler) deleteAuditLogs(c *gin.Context, category string) {
	appG := common.Gin{C: c}

	ids := make([]uint, 0)
	rawIDs := c.QueryArray("ids")
	if len(rawIDs) == 0 {
		rawIDs = c.QueryArray("ids[]")
	}
	for _, raw := range rawIDs {
		id, err := strconv.Atoi(raw)
		if err != nil || id <= 0 {
			appG.Response(http.StatusBadRequest, code.InvalidParams, "无效日志ID", nil)
			return
		}
		ids = append(ids, uint(id))
	}

	if err := h.service.DeleteAuditLogs(category, ids); err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "删除日志失败", nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "删除成功", nil)
}
