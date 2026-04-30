package userService

import (
	"fmt"
	"net/http"
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

type ResetPasswordStruct struct {
	NewPassword string `json:"newPassword" form:"newPassword" validate:"required,min=6,max=20" minLength:"6" maxLength:"20"`
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

type UserSecurityEvent struct {
	ID        uint   `json:"id"`
	Category  string `json:"category"`
	Action    string `json:"action"`
	Message   string `json:"message"`
	IP        string `json:"ip"`
	Status    int    `json:"status"`
	CreatedAt string `json:"createdAt"`
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
		switch u.Status {
		case "0":
			maps["status ="] = 0
		case "1":
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

func (s *Service) SetLoggedUserInfo(userId uint, refreshToken string, ip string) error {
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

	rowsAffected, err := model.Update(&model.Auth{}, wheres, updates)
	if err != nil {
		return fmt.Errorf("更新用户登录信息失败: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("未找到要更新的用户信息")
	}

	profile, err := s.GetAuthProfile(userId)
	if err != nil {
		logging.Warnf("load auth profile cache failed: %v", err)
		return nil
	}
	if profile != nil && ip != "" {
		profile.LastLoginIP = ip
		if gredis.RedisConn != nil {
			gredis.SetJSONAsync(authProfileCacheKey(userId), profile, authProfileCacheTTL)
		}
	}
	return nil
}

func (s *Service) RefreshAccessToken(refreshToken, ip, userAgent string) (map[string]any, error) {
	data := make(map[string]any)
	_, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return data, err
	}
	claims, err := utils.ParseToken(refreshToken)
	if err != nil || claims == nil || claims.UserId == 0 {
		return data, fmt.Errorf("该access_token对应的用户信息不存在")
	}

	profile, err := s.GetAuthProfile(claims.UserId)
	if err != nil || profile == nil {
		return data, fmt.Errorf("该access_token对应的用户信息不存在")
	}

	tokenClaims := utils.Claims{
		UserId:   profile.UserID,
		Username: profile.Username,
		RoleKey:  profile.RoleKey,
		IsAdmin:  profile.IsAdmin,
	}

	accessToken, expireTime, err := utils.GenerateToken(tokenClaims)
	if err != nil {
		return data, fmt.Errorf("%s", code.GetMsg(code.AccessTokenFailure))
	}
	if gredis.RedisConn != nil {
		if err := s.RefreshOnlineSessionAccessToken(accessToken, refreshToken, &tokenClaims, ip, userAgent); err != nil {
			return data, err
		}
	} else {
		user, getErr := model.GetUser(map[string]any{"id": claims.UserId})
		if getErr != nil || user == nil || user.RefreshToken != refreshToken {
			return data, fmt.Errorf("该access_token对应的用户信息不存在")
		}
	}

	data["expires"] = expireTime.Format("2006/01/02 15:04:05")
	data["accessToken"] = accessToken
	data["refreshToken"] = refreshToken

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
	updates["refresh_token"] = ""
	updates["failed_login_count"] = 0
	updates["locked_until"] = nil
	rowsAffected, err := model.Update(&model.Auth{}, wheres, updates)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("修改用户密码失败，用户不存在或未更新")
	}

	s.InvalidateAllUserSessions(userId, true)
	s.InvalidateAuthProfile(userId)
	s.logSecurityEvent(userId, "", http.StatusOK, "password_changed", "用户修改密码，已强制下线全部会话")

	return nil
}

func (s *Service) ResetUserPassword(userId uint, newPassword string, operator string, ip string) error {
	if userId == 0 {
		return fmt.Errorf("用户ID不能为空")
	}
	if err := security.ValidatePasswordComplexity(newPassword); err != nil {
		return err
	}

	user, err := model.GetUser(map[string]any{"id": userId})
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("用户不存在")
	}

	updates := map[string]any{
		"password":           utils.EncodeUserPassword(newPassword),
		"refresh_token":      "",
		"failed_login_count": 0,
		"locked_until":       nil,
	}
	rowsAffected, err := model.Update(&model.Auth{}, map[string]any{"id =": userId}, updates)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("重置用户密码失败，用户不存在或未更新")
	}

	s.InvalidateAllUserSessions(userId, true)
	s.InvalidateAuthProfile(userId)
	s.logSecurityEvent(userId, ip, http.StatusOK, "password_reset", fmt.Sprintf("管理员 %s 重置密码，已强制下线全部会话", fallbackOperator(operator)))
	return nil
}

func (s *Service) UnlockUser(userId uint, operator string, ip string) error {
	if userId == 0 {
		return fmt.Errorf("用户ID不能为空")
	}
	user, err := model.GetUser(map[string]any{"id": userId})
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("用户不存在")
	}
	rowsAffected, err := model.Update(&model.Auth{}, map[string]any{"id =": userId}, map[string]any{
		"failed_login_count": 0,
		"locked_until":       nil,
	})
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("解锁用户失败，用户不存在或未更新")
	}
	s.InvalidateAuthProfile(userId)
	s.logSecurityEvent(userId, ip, http.StatusOK, "account_unlocked", fmt.Sprintf("管理员 %s 手动解锁账号", fallbackOperator(operator)))
	return nil
}

func (s *Service) GetUserSecurityTimeline(userId uint, pagination utils.Pagination) (utils.PageResult, error) {
	if userId == 0 {
		return pagination.Result([]UserSecurityEvent{}, 0), nil
	}
	db := model.DB().
		Model(&model.AuditLog{}).
		Where("user_id = ? AND category IN ?", userId, []string{"login", "security", "operation"}).
		Where("action IN ? OR path IN ?", []string{
			"login",
			"account_locked",
			"account_unlocked",
			"password_changed",
			"password_reset",
			"force_offline",
		}, []string{
			"/v1/api/login",
			"/v1/api/user/change_password",
		})

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return utils.PageResult{}, err
	}

	var logs []model.AuditLog
	err := db.
		Order("id DESC").
		Scopes(pagination.Scope()).
		Find(&logs).Error
	if err != nil {
		return utils.PageResult{}, err
	}

	events := make([]UserSecurityEvent, 0, len(logs))
	for _, item := range logs {
		events = append(events, UserSecurityEvent{
			ID:        item.ID,
			Category:  item.Category,
			Action:    item.Action,
			Message:   item.Message,
			IP:        item.IP,
			Status:    item.Status,
			CreatedAt: formatSecurityEventTime(item.CreatedAt),
		})
	}
	return pagination.Result(events, total), nil
}

func (s *Service) logSecurityEvent(userId uint, ip string, status int, action string, message string) {
	user, _ := model.GetUser(map[string]any{"id": userId})
	username := ""
	if user != nil {
		username = user.Username
	}
	entry := model.AuditLog{
		Category: "security",
		UserID:   userId,
		Username: username,
		IP:       ip,
		Method:   http.MethodPost,
		Path:     "/v1/api/user/security",
		Status:   status,
		Action:   action,
		Message:  message,
	}
	go model.CreateAuditLog(entry)
}

func fallbackOperator(operator string) string {
	operator = strings.TrimSpace(operator)
	if operator == "" {
		return "unknown"
	}
	return operator
}

func formatSecurityEventTime(t model.JSONTime) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func (s *Service) JoinBlockList(userId uint, jwt string) {
	if jwt == "" {
		return
	}
	if err := s.cacheJWTBlacklist(userId, jwt); err != nil {
		logging.Warnf("write jwt blacklist redis failed: %v", err)
	}
	_ = model.CreateBlockList(userId, jwt)
	s.RemoveOnlineSessionByAccessToken(userId, jwt)
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
