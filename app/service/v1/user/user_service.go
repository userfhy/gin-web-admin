package userService

import (
	"fmt"
	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/logging"
	"strings"
	"time"
)

// RefreshAccessTokenStruct 刷新令牌结构体
type RefreshAccessTokenStruct struct {
	RefreshToken string `json:"refreshToken" form:"refresh_token" validate:"required"`
}

// AuthStruct 用户登录结构体
type AuthStruct struct {
	Username string `json:"username" form:"username" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
	Password string `json:"password" form:"password" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
}

type ChangePasswordStruct struct {
	OldPassword string `json:"oldpassword" form:"oldpassword" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
	NewPassword string `json:"newpassword" form:"newpassword" validate:"required,min=6,max=20" minLength:"6",maxLength:"20"`
}

// 添加用户
type AddUserStruct struct {
	AuthStruct
	RoleId uint `json:"role_id" validate:"omitempty,numeric,min=0"`
}

type UserStruct struct {
	ID       int    `json:"id"`
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Sex      string `form:"sex"`
	Status   string `form:"status" validate:"omitempty,numeric,min=0"`

	PageNum  int
	PageSize int
}

type TestList struct {
	Index int `json:"index"`
	*model.Auth
}

func (u *UserStruct) getConditionMaps() map[string]any {
	maps := make(map[string]any)
	maps["deleted_at is"] = nil
	if u.Username != "" {
		maps["username like"] = "%" + u.Username + "%"
	}

	if u.Status == "0" {
		maps["status ="] = 0
	} else if u.Status == "1" {
		maps["status ="] = 1
	}

	return maps
}

// SetLoggedUserInfo 设置登录用户信息
func SetLoggedUserInfo(userId uint, refreshToken string) error {
	wheres := map[string]any{
		"id =": userId,
	}

	updates := map[string]any{
		"logged_in_at":  time.Now(),
		"refresh_token": refreshToken,
	}

	err, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if err != nil {
		return fmt.Errorf("更新用户登录信息失败: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("未找到要更新的用户信息")
	}
	return nil
}

func RefreshAccessToken(RefreshToken string) (map[string]any, error) {
	data := make(map[string]any)
	_, err := utils.ValidateToken(RefreshToken)
	if err != nil {
		return data, err
	}
	// 判断 token 是否正确
	user, _ := model.GetUser(map[string]any{"refresh_token": RefreshToken})
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
		return data, fmt.Errorf("%s", code.GetMsg(code.AccessTokenFailure))
	}

	data["expires"] = expireTime.Format("2006/01/02 15:04:05")
	data["accessToken"] = accessToken
	data["refreshToken"] = user.RefreshToken

	return data, nil
}

func ChangeUserPassword(userId uint, newPassword string) bool {
	wheres := make(map[string]any)
	wheres["id ="] = userId

	updates := make(map[string]any)
	updates["password"] = utils.EncodeUserPassword(newPassword)
	_, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if rowsAffected == 0 {
		logging.Println("修改用户密码失败！")
		return false
	}
	return true
}

func JoinBlockList(userId uint, jwt string) {
	_ = model.CreateBlockList(userId, jwt)
	_, _ = model.Update(model.Auth{}, map[string]any{"id =": userId}, map[string]any{"refresh_token": userId})
}

func InBlockList(jwt string) (int64, error) {
	wheres := make(map[string]any)
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
