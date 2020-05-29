package model

import (
    "gin-test/utils"
    "github.com/jinzhu/gorm"
    "time"
)

type Auth struct {
    BaseModel
    RoleType int `gorm:"Size:4;DEFAULT:0;NOT NULL;" json:"role_type"`
    LoggedInAt JSONTime `json:"logged_in_at"`
    Username string `gorm:"Size:20;UNIQUE_INDEX;NOT NULL;" json:"user_name"`
    Password string `gorm:"Size:50;NOT NULL;" json:"-"`
}

func CheckAuth(username, password string) (bool, uint) {
    var auth Auth
    db.Select("id").Where(Auth{
        Username : username,
        Password : utils.EncodeMD5(password),
    }).First(&auth)

    if auth.ID > 0 {
        // 记录登录时间
        db.Model(&auth).Update("logged_in_at", time.Now())
        return true, auth.ID
    }

    return false, 0
}

func GetUserTotal(tableStruct interface{}, maps interface{}) (int, error) {
    var count int
    if err := db.Model(tableStruct).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }

    return count, nil
}

// GetTestUsers gets a list of users based on paging constraints
func GetUsers(tableStruct interface{}, pageNum int, pageSize int, maps interface{}) ([]*Auth, error) {
    var user [] *Auth
    err := db.Find(&user).Where(maps).Offset(pageNum).Limit(pageSize).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return user, nil
}