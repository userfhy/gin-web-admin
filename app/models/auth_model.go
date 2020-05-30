package model

import (
    "gin-test/utils"
    "github.com/jinzhu/gorm"
)

type Auth struct {
    BaseModel
    RoleType int `gorm:"Size:4;DEFAULT:0;NOT NULL;" json:"role_type"`
    LoggedInAt JSONTime `json:"logged_in_at"`
    Username string `gorm:"Size:20;UNIQUE_INDEX;NOT NULL;" json:"user_name"`
    Password string `gorm:"Size:50;NOT NULL;" json:"-"`
}

func CheckAuth(username string, password string) (bool, uint) {
    var auth Auth
    db.Select("id").Where(Auth{
        Username : username,
        Password : utils.EncodeMD5(password),
    }).First(&auth)

    if auth.ID > 0 {
        return true, auth.ID
    }

    return false, 0
}

// GetTestUsers gets a list of users based on paging constraints
func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*Auth, error) {
    var user [] *Auth
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&user).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return user, nil
}