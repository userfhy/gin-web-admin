package userService

import (
    model "gin-test/app/models"
)

// 用户登录
type UserLoginStruct struct {
    // 用户名
    Username string `json:"username" form:"username" validate:"required,min=4" minLength:"4"`
    // 密码
    Password string `json:"password" form:"password" validate:"required,min=6" minLength:"6"`
}

type UserStruct struct {
    ID   int
    Name string
    
    PageNum  int
    PageSize int
}

func (u *UserStruct) getMaps() map[string]interface{} {
    maps := make(map[string]interface{})

    return maps
}

func (u *UserStruct) Count() (int, error) {
    return model.GetUserTotal(u.getMaps())
}

func (u *UserStruct) GetAll() ([]*model.User, error) {
    var (
        testUsers [] *model.User
    )

    testUsers, err := model.GetTestUsers(u.PageNum, u.PageSize)
    if err != nil {
        return nil, err
    }

    return testUsers, nil
}