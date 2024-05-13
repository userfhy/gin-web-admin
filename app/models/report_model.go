package model

type Report struct {
	BaseModel
	ActivityId int    `json:"activity_id"`
	Name       string `gorm:"Size:20" json:"name"`
	Phone      string `gorm:"Size:30;index:idx_phone" json:"phone"`
	Ip         string `gorm:"Size:80" json:"ip"`
}

func (Report) TableName() string {
	return TablePrefix + "report"
}

func GetReportUserCount(r Report) int64 {
	var count int64
	db.Model(&Report{}).Where(&r).Count(&count)
	return count
}

func CreateReportNewRecord(r Report) Report {
	//db.NewRecord(r)
	db.Create(&r)
	return r
}
