package model

import (
    "github.com/jinzhu/gorm"
)

type Report struct {
    gorm.Model
    ActivityId int `json:"activity_id"`
    Name string `gorm:"Size:20" json:"name"`
    Phone string `gorm:"Size:30;index:idx_phone" json:"phone"`
    Ip string `gorm:"Size:80" json:"ip"`
}

func CreateReportNewRecord(r Report) Report {
    //db.NewRecord(r)
    db.Create(&r)
    return r
}
