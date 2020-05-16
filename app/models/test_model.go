package model

import "github.com/jinzhu/gorm"

type User struct {
    ID   int    `gorm:"primary_key AUTO_INCREMENT" json:"id"`
    Name string `json:"name"`
}

// 通过 TableName 方法将 TestUser 表命名为 `t_user`
func (User) TableName() string {
   return "t_user"
}

func GetUserTotal(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }

    return count, nil
}

// GetTestUsers gets a list of users based on paging constraints
func GetTestUsers(pageNum int, pageSize int) ([]*User, error) {
    var testUsers [] *User
    err := db.Offset(pageNum).Limit(pageSize).Find(&testUsers).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    
    return testUsers, nil
}
