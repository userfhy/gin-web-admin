package reportService

import (
	model "gin-web-admin/app/models"
)

// 上报信息
type ReportStruct struct {
	Name       string `json:"name" form:"name" validate:"required,min=1,max=10" minLength:"1",maxLength:"10"`
	Phone      string `json:"phone" form:"phone" validate:"required,numeric,min=4,max=15" minLength:"4",maxLength:"15"`
	ActivityId int    `json:"activity_id" form:"activity_id" validate:"omitempty,numeric,min=1,max=10"`
}

func GetReportUserCountByPhoneAndActivityID(mobile string, activityId int) int64 {
	return model.GetReportUserCount(model.Report{Phone: mobile, ActivityId: activityId})
}

// 录入信息
func ReportInformation(report ReportStruct, ip string) model.Report {
	return model.CreateReportNewRecord(model.Report{
		Name:       report.Name,
		Phone:      report.Phone,
		ActivityId: report.ActivityId,
		Ip:         ip,
	})
}
