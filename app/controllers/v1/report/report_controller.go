package reportController

import (
	"net/http"

	reportService "gin-web-admin/app/service/v1/report"
	"gin-web-admin/common"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *reportService.Service
}

func NewHandler(service *reportService.Service) *Handler {
	if service == nil {
		panic("report handler requires non-nil service")
	}
	return &Handler{service: service}
}

// @Summary		Report Information
// @Description	User Report Information
// @Accept			json
// @Produce		json
// @Tags			Report
// @Param			payload	body		reportService.ReportStruct	true	"上报信息"
// @Success		200		{object}	common.Response
// @Router			/report [post]
func (h *Handler) Report(c *gin.Context) {
	appG := common.Gin{C: c}

	// 绑定 payload 到结构体
	var report reportService.ReportStruct
	if err := c.ShouldBindJSON(&report); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	// 验证绑定结构体参数
	if parameterErrorStr, err := common.CheckBindStructParameter(report, c); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, parameterErrorStr, nil)
		return
	}

	// 是否存在
	var count = h.service.GetReportUserCountByPhoneAndActivityID(report.Phone, report.ActivityId)
	if count >= 1 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "已经存在数据，请勿重复报名！", nil)
		return
	}

	if report.ActivityId == 0 {
		report.ActivityId = 1
	}

	// 信息入库
	var reportResult = h.service.ReportInformation(report, c.ClientIP())

	if reportResult.ID == 0 {
		appG.Response(http.StatusInternalServerError, code.ERROR, "录入失败，请稍后再试。", nil)
		return
	}

	m := make(map[string]any)
	m["id"] = reportResult.ID
	m["name"] = report.Name
	//m["created_at"] = utils.TimeToDateTimesString(reportResult.CreatedAt)
	appG.Response(http.StatusOK, code.SUCCESS, "信息录入成功！", m)
}
