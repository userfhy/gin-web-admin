package reportService

import (
    model "gin-test/app/models"
)

// 上报信息
type ReportStruct struct {
    Name string `json:"name" form:"name" validate:"required,min=1,max=10" minLength:"1",maxLength:"10"`
    Phone string `json:"phone" form:"phone" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
}

// 录入信息
func ReportInformation(report ReportStruct) model.Report {
    return model.CreateReportNewRecord(model.Report{
        Name: report.Name,
        Phone: report.Phone,
    })
}