package userService

import (
    model "gin-test/app/models"
)

// 用户登录
type AuthStruct struct {
    Username string `json:"user_name" form:"user_name" validate:"required,min=1,max=20" minLength:"1",maxLength:"20"`
    Password string `json:"password" form:"password" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
}

type UserStruct struct {
    ID   int `json:"id"`
    
    PageNum  int
    PageSize int
}

func (u *UserStruct) getMaps() map[string]interface{} {
    maps := make(map[string]interface{})
    maps["deleted_at"] = nil
    return maps
}

func (u *UserStruct) Count() (int, error) {
    return model.GetUserTotal(u.getMaps())
}

func (u *UserStruct) GetAll() ([]*model.Auth, error) {
    Users, err := model.GetUsers(u.PageNum, u.PageSize, u.getMaps())
    if err != nil {
        return nil, err
    }

    return Users, nil
}