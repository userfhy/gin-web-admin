package model

import "github.com/jinzhu/gorm"

type Test struct {
    ID   int    `gorm:"primary_key" json:"id"`
    Name string `json:"name"`
}

// 通过TableName方法将User表命名为`profiles`
func (Test) TableName() string {
    return "test"
}

// GetArticles gets a list of articles based on paging constraints
func GetTestUsers(pageNum int, pageSize int) ([]*Test, error) {
    var testUsers [] *Test
    err := db.Offset(pageNum).Limit(pageSize).Find(&testUsers).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    
    return testUsers, nil
}
