package model

import (
    "github.com/jinzhu/gorm"
)

type Report struct {
    gorm.Model
    ActivityId int `json:"activity_id"`
    Name string `gorm:"Size:20" json:"name"`
    Phone string `gorm:"Size:30" json:"phone"`
}

func CreateReportNewRecord(r Report) Report {
    //db.NewRecord(r)
    db.Create(&r)
    return r
}
