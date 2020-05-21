package reportController

import (
    reportService "gin-test/app/service/v1/report"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "net/http"
)

// @Summary Report Information
// @Description User Report Information
// @Accept json
// @Produce json
// @Tags Report
// @Param payload body reportService.ReportStruct true "上报信息"
// @Success 200 {object} common.Response
// @Router /report [post]
func Report(c *gin.Context) {
    appG := common.Gin{C: c}

    // 绑定 payload 到结构体
    var report reportService.ReportStruct
    if err := c.ShouldBindJSON(&report); err != nil {
        appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
        return
    }

    // 验证绑定结构体参数
    err, parameterErrorStr := common.CheckBindStructParameter(report, c)
    if err != nil {
        appG.Response(http.StatusBadRequest, code.InvalidParams, parameterErrorStr, nil)
        return
    }

    // 是否存在
    var count = reportService.GetReportUserCountByPhoneAndActivityID(report.Phone, report.ActivityId)
    if count >= 1 {
        appG.Response(http.StatusBadRequest, code.InvalidParams, "已经存在数据，请勿重复报名！", nil)
        return
    }

    if report.ActivityId == 0 {
        report.ActivityId = 1
    }

    // 信息入库
    var reportResult = reportService.ReportInformation(report, c.ClientIP())

    if reportResult.ID == 0 {
       appG.Response(http.StatusInternalServerError, code.ERROR, "录入失败，请稍后再试。", nil)
       return
    }

    m := make(map[string]interface{})
    m["id"] = reportResult.ID
    m["name"] = report.Name
    m["created_at"] = utils.TimeToDateTimesString(reportResult.CreatedAt)
    appG.Response(http.StatusOK, code.SUCCESS, "信息录入成功！", m)
}
