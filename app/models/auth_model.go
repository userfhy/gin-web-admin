package model

import (
    "gin-test/utils"
    "github.com/jinzhu/gorm"
    "time"
)

type Auth struct {
    gorm.Model
    LoggedInAt time.Time
    Username string `gorm:"Size:20" json:"user_name"`
    Password string `gorm:"Size:50" json:"password"`
}

func CheckAuth(username, password string) (bool, uint) {
    var auth Auth
    db.Select("id").Where(Auth{
        Username : username,
        Password : utils.EncodeMD5(password),
    }).First(&auth)

    if auth.ID > 0 {
        // 记录登录时间
        db.Model(&auth).Update(Auth{LoggedInAt: time.Now()})
        return true, auth.ID
    }

    return false, 0
}