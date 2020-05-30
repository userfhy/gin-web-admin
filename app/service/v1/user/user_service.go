package userService

import (
    model "gin-test/app/models"
    "log"
    "time"
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

type TestList struct {
    Index int `json:"index"`
    *model.Auth
}

func (u *UserStruct) getConditionMaps() map[string]interface{} {
    maps := make(map[string]interface{})
    //maps["deleted_at"] = nil
    return maps
}

// 设置登录时间
func SetLoggedTime(userId uint)  {
    wheres := make(map[string]interface{})
    wheres["id"] = userId

    updates := make(map[string]interface{})
    updates["logged_in_at"] = time.Now()
    _, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
    if rowsAffected == 0 {
        log.Println("设置登录时间失败！")
    }
}

func JoinBlockList(userId uint, jwt string) {
    _ = model.CreatCreateBlockList(userId, jwt)
}

func InBlockList(jwt string) (int, error){
    wheres := make(map[string]interface{})
    wheres["jwt"] = jwt
    return model.GetTotal(model.JwtBlacklist{}, wheres)
}

func (u *UserStruct) Count() (int, error) {
    return model.GetTotal(model.Auth{}, u.getConditionMaps())
}

func (u *UserStruct) GetAll() ([]*model.Auth, error) {
    Users, err := model.GetUsers(u.PageNum, u.PageSize, u.getConditionMaps())
    if err != nil {
        return nil, err
    }

    return Users, nil
}