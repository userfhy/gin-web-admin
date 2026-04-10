package reportService

import (
	model "gin-web-admin/app/models"
	"gin-web-admin/internal/data"
)

// 上报信息
type ReportStruct struct {
	Name       string `json:"name" form:"name" validate:"required,min=1,max=10" minLength:"1" maxLength:"10"`
	Phone      string `json:"phone" form:"phone" validate:"required,numeric,min=4,max=15" minLength:"4" maxLength:"15"`
	ActivityId int    `json:"activity_id" form:"activity_id" validate:"omitempty,numeric,min=1,max=10"`
}

type Service struct {
	store *data.Store
}

var defaultService *Service

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func SetDefaultService(s *Service) {
	defaultService = s
}

func serviceInstance() *Service {
	if defaultService == nil {
		panic("report service not initialized")
	}
	return defaultService
}

func GetReportUserCountByPhoneAndActivityID(mobile string, activityId int) int64 {
	return serviceInstance().GetReportUserCountByPhoneAndActivityID(mobile, activityId)
}

func (s *Service) GetReportUserCountByPhoneAndActivityID(mobile string, activityId int) int64 {
	return model.GetReportUserCount(model.Report{Phone: mobile, ActivityId: activityId})
}

// 录入信息
func ReportInformation(report ReportStruct, ip string) model.Report {
	return serviceInstance().ReportInformation(report, ip)
}

func (s *Service) ReportInformation(report ReportStruct, ip string) model.Report {
	return model.CreateReportNewRecord(model.Report{
		Name:       report.Name,
		Phone:      report.Phone,
		ActivityId: report.ActivityId,
		Ip:         ip,
	})
}
