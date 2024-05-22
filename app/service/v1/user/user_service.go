package userService

import (
	"fmt"
	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"log"
	"strings"
	"time"
)

type RefreshAccessTokenhStruct struct {
	RefreshToken string `json:"refreshToken" form:"refresh_token" validate:"required"`
}

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
	RoleId uint `json:"role_id" validate:"omitempty,numeric,min=0"`
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
	maps["deleted_at is"] = nil
	return maps
}

// 设置登录用户信息
func SetLoggedUserInfo(userId uint, refreshToken string) error {
	wheres := make(map[string]interface{})
	wheres["id"] = userId

	updates := make(map[string]interface{})
	updates["logged_in_at"] = time.Now()
	updates["refresh_token"] = refreshToken
	err, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		log.Println("设置登录信息失败！")
	}
	return nil
}

func RefreshAccessToken(RefreshToken string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	_, err := utils.ValidateToken(RefreshToken)
	if err != nil {
		return data, err
	}
	// 判断 token 是否正确
	user, _ := model.GetUser(map[string]interface{}{"refresh_token": RefreshToken})
	if user.ID == 0 {
		return data, fmt.Errorf("该access_token对应的用户信息不存在")
	}

	claims := utils.Claims{
		UserId:   user.ID,
		Username: user.Username,
		RoleKey:  user.Role.RoleKey,
		IsAdmin:  user.Role.IsAdmin,
	}

	accessToken, expireTime, err := utils.GenerateToken(claims)
	if err != nil {
		return data, fmt.Errorf(code.GetMsg(code.AccessTokenFailure))
	}

	data["expires"] = expireTime.Format("2006/01/02 15:04:05")
	data["accessToken"] = accessToken
	data["refreshToken"] = user.RefreshToken

	return data, nil
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
	_, _ = model.Update(model.Auth{}, map[string]interface{}{"id": userId}, map[string]interface{}{"refresh_token": userId})
}

func InBlockList(jwt string) (int64, error) {
	wheres := make(map[string]interface{})
	wheres["jwt ="] = jwt
	return model.GetTotal(model.JwtBlacklist{}, wheres)
}

func CreateUser(newUser AddUserStruct) error {
	return model.CreatUser(model.Auth{
		Username: strings.TrimSpace(newUser.Username),
		Password: utils.EncodeUserPassword(newUser.Password),
		RoleId:   newUser.RoleId,
	})
}

func (u *UserStruct) Count() (int64, error) {
	return model.GetTotal(model.Auth{}, u.getConditionMaps())
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
