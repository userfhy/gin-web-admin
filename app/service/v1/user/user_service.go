package userService

import (
	model "gin-test/app/models"
	"gin-test/utils"
	"log"
	"strings"
	"time"
)

// 用户登录
type AuthStruct struct {
	Username string `json:"user_name" form:"user_name" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
	Password string `json:"password" form:"password" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
}

type ChangePasswordStruct struct {
	OldPassword string `json:"old_password" form:"old_password" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required,min=6,max=20" minLength:"6",maxLength:"20"`
}

// 添加用户
type AddUserStruct struct {
	AuthStruct
	RoleId int `json:"role_id" validate:"omitempty,numeric,min=0"`
}

type UserStruct struct {
	ID int `json:"id"`

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
func SetLoggedTime(userId uint) {
	wheres := make(map[string]interface{})
	wheres["id"] = userId

	updates := make(map[string]interface{})
	updates["logged_in_at"] = time.Now()
	_, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if rowsAffected == 0 {
		log.Println("设置登录时间失败！")
	}
}

func ChangeUserPassword(userId uint, newPassword string) bool {
	wheres := make(map[string]interface{})
	wheres["id"] = userId

	updates := make(map[string]interface{})
	updates["password"] = utils.EncodeUserPassword(newPassword)
	_, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if rowsAffected == 0 {
		log.Println("修改用户密码失败！")
		return false
	}
	return true
}

func JoinBlockList(userId uint, jwt string) {
	_ = model.CreateBlockList(userId, jwt)
}

func InBlockList(jwt string) (int, error) {
	wheres := make(map[string]interface{})
	wheres["jwt"] = jwt
	whereSql, values, err := model.BuildCondition(wheres)
	if err != nil {
		return 0, err
	}
	return model.GetTotal(model.JwtBlacklist{}, whereSql, values)
}

func CreateUser(newUser AddUserStruct) error {
	return model.CreatUser(model.Auth{
		Username: strings.TrimSpace(newUser.Username),
		Password: utils.EncodeUserPassword(newUser.Password),
		RoleId:   newUser.RoleId,
	})
}

func (u *UserStruct) Count() (int, error) {
	whereSql, values, err := model.BuildCondition(u.getConditionMaps())
	if err != nil {
		return 0, err
	}
	return model.GetTotal(model.Auth{}, whereSql, values)
}

func (u *UserStruct) GetAll() ([]*model.Auth, error) {
	Users, err := model.GetUsers(u.PageNum, u.PageSize, u.getConditionMaps())
	if err != nil {
		return nil, err
	}

	for i := range Users {
		Users[i].RoleName = Users[i].Role.RoleName
	}

	return Users, nil
}
