package userService

import (
	"fmt"
	"strings"
	"time"

	model "gin-web-admin/app/models"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
	security "gin-web-admin/utils/security"
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
	OldPassword string `json:"oldpassword" form:"oldpassword" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
	NewPassword string `json:"newpassword" form:"newpassword" validate:"required,min=6,max=20" minLength:"6" maxLength:"20"`
}

// 添加用户
type AddUserStruct struct {
	AuthStruct
	RoleId uint `json:"role_id" validate:"omitempty,numeric,min=0"`
}

type UserStruct struct {
	ID         int    `json:"id"`
	Username   string `form:"username"`
	Nickname   string `form:"nickname"`
	Phone      string `form:"phone"`
	Email      string `form:"email"`
	Sex        string `form:"sex"`
	Status     string `json:"status" validate:"omitempty,numeric,min=0"`
	Pagination utils.Pagination
	Conditions map[string]any
}

type TestList struct {
	Index int `json:"index"`
	*model.Auth
}

func (u *UserStruct) getConditionMaps() map[string]any {
	maps := make(map[string]any)
	for k, v := range u.Conditions {
		maps[k] = v
	}
	if _, ok := maps["deleted_at is"]; !ok {
		maps["deleted_at is"] = nil
	}
	if _, ok := maps["username like"]; !ok && u.Username != "" {
		maps["username like"] = "%" + u.Username + "%"
	}

	if _, ok := maps["status ="]; !ok {
		if u.Status == "0" {
			maps["status ="] = 0
		} else if u.Status == "1" {
			maps["status ="] = 1
		}
	}

	return maps
}

type Service struct {
	store *data.Store
}

const (
	jwtBlacklistKeyPrefix = "jwt:blacklist:"
	defaultBlacklistTTL   = 24 * time.Hour
)

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func (s *Service) SetLoggedUserInfo(userId uint, refreshToken string, ip string) (string, error) {
	var (
		oldRefresh = ""
	)
	if user, err := model.GetUser(map[string]any{"id": userId}); err == nil && user != nil {
		oldRefresh = user.RefreshToken
	}

	wheres := map[string]any{
		"id =": userId,
	}

	updates := map[string]any{
		"logged_in_at":  time.Now(),
		"refresh_token": refreshToken,
	}
	if ip != "" {
		updates["last_login_ip"] = ip
		updates["failed_login_count"] = 0
		updates["locked_until"] = nil
	}

	err, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if err != nil {
		return "", fmt.Errorf("更新用户登录信息失败: %w", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("未找到要更新的用户信息")
	}
	return oldRefresh, nil
}

func (s *Service) RefreshAccessToken(refreshToken string) (map[string]any, error) {
	data := make(map[string]any)
	_, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return data, err
	}
	// 判断 token 是否正确
	user, _ := model.GetUser(map[string]any{"refresh_token": refreshToken})
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

func (s *Service) ChangeUserPassword(userId uint, newPassword string) error {
	if err := security.ValidatePasswordComplexity(newPassword); err != nil {
		return err
	}
	wheres := make(map[string]any)
	wheres["id ="] = userId

	updates := make(map[string]any)
	updates["password"] = utils.EncodeUserPassword(newPassword)
	_, rowsAffected := model.Update(&model.Auth{}, wheres, updates)
	if rowsAffected == 0 {
		return fmt.Errorf("修改用户密码失败，用户不存在或未更新")
	}

	return nil
}

func (s *Service) JoinBlockList(userId uint, jwt string) {
	if jwt == "" {
		return
	}
	if err := s.cacheJWTBlacklist(userId, jwt); err != nil {
		logging.Warnf("write jwt blacklist redis failed: %v", err)
	}
	_ = model.CreateBlockList(userId, jwt)
	_, _ = model.Update(model.Auth{}, map[string]any{"id =": userId}, map[string]any{"refresh_token": ""})
}

func (s *Service) InBlockList(jwt string) (int64, error) {
	if jwt == "" {
		return 0, nil
	}
	ok, err := s.redisJWTBlocked(jwt)
	if err != nil {
		logging.Warnf("check jwt blacklist redis failed: %v", err)
	} else if ok {
		return 1, nil
	} else if err == nil {
		// redis 确认不存在，无需访问数据库
		return 0, nil
	}

	wheres := make(map[string]any)
	wheres["jwt ="] = jwt
	return model.GetTotal(model.JwtBlacklist{}, wheres)
}

func (s *Service) cacheJWTBlacklist(userId uint, jwt string) error {
	ttl := blacklistTTL(jwt)
	if ttl <= 0 {
		ttl = defaultBlacklistTTL
	}
	return gredis.SetWithTTL(jwtBlacklistKey(jwt), fmt.Sprintf("%d", userId), ttl)
}

func (s *Service) redisJWTBlocked(jwt string) (bool, error) {
	return gredis.ExistsKey(jwtBlacklistKey(jwt))
}

func jwtBlacklistKey(jwt string) string {
	return jwtBlacklistKeyPrefix + jwt
}

func blacklistTTL(jwt string) time.Duration {
	claims, err := utils.ParseToken(jwt)
	if err != nil || claims == nil || claims.ExpiresAt == nil {
		return defaultBlacklistTTL
	}
	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		return time.Minute
	}
	return ttl
}

func (s *Service) CreateUser(newUser AddUserStruct) error {
	if err := security.ValidatePasswordComplexity(newUser.Password); err != nil {
		return err
	}
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
	Users, err := model.GetUsers(u.Pagination, u.getConditionMaps())
	if err != nil {
		return nil, err
	}

	for i := range Users {
		Users[i].RoleName = Users[i].Role.RoleName
	}

	return Users, nil
}
