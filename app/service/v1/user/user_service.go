package userService

import (
    model "gin-test/app/models"
)

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